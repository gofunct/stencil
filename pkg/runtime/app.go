package runtime

import (
	"github.com/izumin5210/gex"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// App contains dependencies to execute a generator.
type App struct {
	Name       string
	AbsPath    string
	CmdPath    string
	SrcPath    string
	Runtime    *Runtime
	GexCfg     *gex.Config
	TemplateFS http.FileSystem
}

// NewProject returns App with specified project Name.
func NewApp(projectName string) *App {
	if projectName == "" {
		er("can't create project with blank Name")
	}

	a := new(App)
	a.Name = projectName

	// 1. Find already created protect.
	a.AbsPath = a.Runtime.FindPackage(projectName)

	// 2. If there are no created project with this path, and user is in GOPATH,
	// then use GOPATH/src/projectName.
	if a.AbsPath == "" {
		wd, err := os.Getwd()
		if err != nil {
			er(err)
		}
		for _, srcPath := range srcPaths {
			goPath := filepath.Dir(srcPath)
			if a.Runtime.FilepathHasPrefix(wd, goPath) {
				a.AbsPath = filepath.Join(srcPath, projectName)
				break
			}
		}
	}

	// 3. If user is not in GOPATH, then use (first GOPATH)/src/projectName.
	if a.AbsPath == "" {
		a.AbsPath = filepath.Join(srcPaths[0], projectName)
	}

	return a
}

// NewProjectFromPath returns App with specified absolute path to
// package.
func NewAppFromPath(AbsPath string) *App {
	if AbsPath == "" {
		er("can't create project: AbsPath can't be blank")
	}
	if !filepath.IsAbs(AbsPath) {
		er("can't create project: AbsPath is not absolute")
	}

	// If AbsPath is symlink, use its destination.
	fi, err := os.Lstat(AbsPath)
	if err != nil {
		er("can't read path info: " + err.Error())
	}
	if fi.Mode()&os.ModeSymlink != 0 {
		path, err := os.Readlink(AbsPath)
		if err != nil {
			er("can't read the destination of symlink: " + err.Error())
		}
		AbsPath = path
	}

	a := new(App)
	a.AbsPath = strings.TrimSuffix(AbsPath, a.Runtime.FindCmdDir(AbsPath))
	a.Name = filepath.ToSlash(a.Runtime.TrimSrcPath(a.AbsPath, a.SrcPath))
	return a
}

// AbsPath returns absolute path of project.
func (a *App) MyAbsPath() string {
	return a.AbsPath
}

// CmdPath returns absolute path to directory, where all commands are located.
func (a *App) MyCmdPath() string {
	if a.AbsPath == "" {
		return ""
	}
	if a.CmdPath == "" {
		a.CmdPath = filepath.Join(a.AbsPath, a.Runtime.FindCmdDir(a.AbsPath))
	}
	return a.CmdPath
}

// SrcPath returns absolute path to $GOPATH/src where project is located.
func (a *App) MySrcPath() string {
	if a.SrcPath != "" {
		return a.SrcPath
	}
	if a.AbsPath == "" {
		a.SrcPath = srcPaths[0]
		return a.SrcPath
	}

	for _, srcPath := range srcPaths {
		if a.Runtime.FilepathHasPrefix(a.AbsPath, srcPath) {
			a.SrcPath = srcPath
			break
		}
	}

	return a.SrcPath
}

func (a *App) CreateCmdFile(path, cmdName string) {
	template := `
package {{.cmdPackage}}

import (
	"fmt"

	"github.com/spf13/cobra"
)

// {{.cmdName}}Cmd represents the {{.cmdName}} command
var {{.cmdName}}Cmd = &cobra.Command{
	Use:   "{{.cmdName}}",
	Short: "A brief description of your command",
	Long: ` + "`" + `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.` + "`" + `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("{{.cmdName}} called")
	},
}

func init() {
	{{.parentName}}.AddCommand({{.cmdName}}Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// {{.cmdName}}Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// {{.cmdName}}Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
`

	data := a.Runtime.FMap

	cmdScript, err := a.Runtime.ExecuteTemplate(template, data)
	if err != nil {
		er(err)
	}
	err = a.Runtime.WriteStringToFile(path, cmdScript)
	if err != nil {
		er(err)
	}
}

func (a *App) ValidateCmdName(source string) string {
	i := 0
	l := len(source)
	// The output is initialized on demand, then first dash or underscore
	// occurs.
	var output string

	for i < l {
		if source[i] == '-' || source[i] == '_' {
			if output == "" {
				output = source[:i]
			}

			// If it's last rune and it's dash or underscore,
			// don't add it output and break the loop.
			if i == l-1 {
				break
			}

			// If next character is dash or underscore,
			// just skip the current character.
			if source[i+1] == '-' || source[i+1] == '_' {
				i++
				continue
			}

			// If the current character is dash or underscore,
			// upper next letter and add to output.
			output += string(unicode.ToUpper(rune(source[i+1])))
			// We know, what source[i] is dash or underscore and source[i+1] is
			// uppered character, so make i = i+2.
			i += 2
			continue
		}

		// If the current character isn't dash or underscore,
		// just add it.
		if output != "" {
			output += string(source[i])
		}
		i++
	}

	if output == "" {
		return source // source is initially valid name.
	}
	return output
}

func (app *App) CreateMainFile() {
	mainTemplate := `{{ comment .copyright }}
{{if .license}}{{ comment .license }}{{end}}

package main

import "{{ .importpath }}"

func main() {
	cmd.Execute()
}
`
	data := app.Runtime.FMap

	mainScript, err := app.Runtime.ExecuteTemplate(mainTemplate, data)
	if err != nil {
		er(err)
	}

	err = app.Runtime.WriteStringToFile(filepath.Join(app.MyAbsPath(), "main.go"), mainScript)
	if err != nil {
		er(err)
	}

}

func (app *App) CreateRootCmdFile() {
	template := `{{comment .copyright}}
{{if .license}}{{comment .license}}{{end}}

package cmd

import (
	"fmt"
	"os"
{{if .viper}}
	homedir "github.com/mitchellh/go-homedir"{{end}}
	"github.com/spf13/cobra"{{if .viper}}
	"github.com/spf13/viper"{{end}}
){{if .viper}}

var cfgFile string{{end}}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "{{.appName}}",
	Short: "A brief description of your application",
	Long: ` + "`" + `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Stencil is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Stencil application.` + "`" + `,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() { {{- if .viper}}
	cobra.OnInitialize(initConfig)
{{end}}
	// Here you will define your flags and configuration settings.
	// Stencil supports persistent flags, which, if defined here,
	// will be global for your application.{{ if .viper }}
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.{{ .appName }}.yaml)"){{ else }}
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.{{ .appName }}.yaml)"){{ end }}

	// Stencil also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}{{ if .viper }}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".{{ .appName }}" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".{{ .appName }}")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}{{ end }}
`

	data := app.Runtime.FMap

	rootScript, err := app.Runtime.ExecuteTemplate(template, data)
	if err != nil {
		er(err)
	}

	err = app.Runtime.WriteStringToFile(filepath.Join(app.MyAbsPath(), "main.go"), rootScript)
	if err != nil {
		er(err)
	}
}
