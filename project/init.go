package project

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go/build"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func init() {
	BuildContext = build.Default
	GetOSUser = user.Current
}

func (p *Project) ProjectInitCommand(v *viper.Viper) *cobra.Command {
	var initCmd = &cobra.Command{
		Use:     "init [name]",
		Aliases: []string{"initialize", "initialise", "create"},
		Short:   "Initialize a Stencil Application",
		Long: `Initialize (cobra init) will create a new application, with a license
and the appropriate structure for a Stencil-based CLI application.
  * If a name is provided, it will be created in the current directory;
  * If no name is provided, the current directory will be assumed;
  * If a relative path is provided, it will be created inside $GOPATH
    (e.g. github.com/spf13/hugo);
  * If an absolute path is provided, it will be created;
  * If the directory already exists but is empty, it will be used.
Init will not use an existing directory with contents.`,

		Run: func(cmd *cobra.Command, args []string) {
			wd, err := os.Getwd()
			if err != nil {
				exit(err)
			}

			var p *Project
			if len(args) == 0 {
				p = CreateProjectFromPath(wd, v)
			} else if len(args) == 1 {
				arg := args[0]
				if arg[0] == '.' {
					arg = filepath.Join(wd, arg)
				}
				if filepath.IsAbs(arg) {
					p = CreateProjectFromPath(arg, v)
				} else {
					p = CreateProject(v)
				}
			} else {
				p.Exit("please provide only one argument")
			}

			p.InitializeProject(v)

			fmt.Fprintln(cmd.OutOrStdout(), `Your Stencil application is ready at
`+p.AbsPath()+`
Give it a try by going there and running `+"`go run main.go`."+`
Add commands to it by running `+"`cobra add [cmdname]`.")
		},
	}
	return initCmd
}

func (p *Project) InitializeProject(v *viper.Viper) {
	if !p.CheckPathExists(p.AbsPath()) { // If path doesn't yet exist, create it
		err := os.MkdirAll(p.AbsPath(), os.ModePerm)
		if err != nil {
			p.Exit(err)
		}
	} else if !p.CheckPathEmpty(p.AbsPath()) { // If path exists and is not empty don't use it
		p.Exit("Stencil will not create a new p in a non empty directory: " + p.AbsPath())
	}

	// We have a directory and it's empty. Time to initialize it.
	p.CreateLicenseFile(p.AbsPath(), v)
	p.CreateMainFile(v)
	p.CreateRootCmdFile(v)
}

func (p *Project) CreateMainFile(v *viper.Viper) {
	mainTemplate := `{{ comment .copyright }}
{{if .license}}{{ comment .license }}{{end}}
package main
import "{{ .importpath }}"
func main() {
	cmd.Execute()
}
`
	data := make(map[string]interface{})
	data["copyright"] = p.CopyrightLine(v)
	data["license"] = p.License().Header
	data["importpath"] = path.Join(p.Name(), filepath.Base(p.CmdPath(v)))

	mainScript, err := p.ExecuteTemplate(mainTemplate, data)
	if err != nil {
		p.Exit(err)
	}

	err = p.WriteStringToFile(filepath.Join(p.AbsPath(), "main.go"), mainScript)
	if err != nil {
		p.Exit(err)
	}
}

func (p *Project) CreateRootCmdFile(v *viper.Viper) {
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

	data := make(map[string]interface{})
	data["copyright"] = p.CopyrightLine(v)
	data["viper"] = v.GetBool("useViper")
	data["license"] = p.License().Header
	data["appName"] = path.Base(p.Name())

	rootCmdScript, err := p.ExecuteTemplate(template, data)
	if err != nil {
		p.Exit(err)
	}

	err = p.WriteStringToFile(filepath.Join(p.CmdPath(v), "root.go"), rootCmdScript)
	if err != nil {
		p.Exit(err)
	}

}
