package cmd

import (
	"fmt"
	"github.com/gofunct/stencil/pkg/runtime"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

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
			//	er(err)
		}

		var app *runtime.App
		if len(args) == 0 {
			app = runtime.NewAppFromPath(wd)
		} else if len(args) == 1 {
			arg := args[0]
			if arg[0] == '.' {
				arg = filepath.Join(wd, arg)
			}
			if filepath.IsAbs(arg) {
				app = runtime.NewAppFromPath(arg)
			} else {
				app = runtime.NewApp(arg)
			}
		} else {
			//	er("please provide only one argument")
		}

		initializeProject(app)

		fmt.Fprintln(cmd.OutOrStdout(), `Your Stencil application is ready at
`+app.MyAbsPath()+`

Give it a try by going there and running `+"`go run main.go`."+`
Add commands to it by running `+"`cobra add [cmdname]`.")
	},
}


func initializeProject(a *runtime.App) {
	app = runtime.NewApp(appName)
	if !a.Runtime.Exists(a.MyAbsPath()) { // If path doesn't yet exist, create it
		err := os.MkdirAll(a.MyAbsPath(), os.ModePerm)
		if err != nil {
			er(err)
		}
	} else if !app.Runtime.IsEmpty(app.MyAbsPath()) { // If path exists and is not empty don't use it
		er("Stencil will not create a new project in a non empty directory: " + app.MyAbsPath())
	}

	// We have a directory and it's empty. Time to initialize it.
	app.CreateMainFile()
	app.CreateRootCmdFile()
}
