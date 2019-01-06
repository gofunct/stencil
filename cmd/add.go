package cmd

import (
	"fmt"
	"github.com/gofunct/stencil/pkg/runtime"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var packageName, parentName string

func init() {
	addCmd.Flags().StringVarP(&packageName, "package", "t", "", "target package name (e.g. github.com/spf13/hugo)")
	addCmd.Flags().StringVarP(&parentName, "parent", "p", "rootCmd", "variable name of parent command for this command")
}

var addCmd = &cobra.Command{
	Use:     "add [command name]",
	Aliases: []string{"command"},
	Short:   "Add a command to a Cobra Application",
	Long: `Add (stencil add) will create a new command, with a license and
the appropriate structure for a Cobra-based CLI application,
and register it to its parent (default rootCmd).

If you want your command to be public, pass in the command name
with an initial uppercase letter.

Example: cobra add server -> resulting in a new cmd/server.go`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			//er("add needs a name for the command")
		}

		var app *runtime.App
		if packageName != "" {
			app = runtime.NewApp(packageName)
		} else {
			wd, err := os.Getwd()
			if err != nil {
				//er(err)
			}
			app = runtime.NewAppFromPath(wd)
		}

		cmdName := app.ValidateCmdName(args[0])
		cmdPath := filepath.Join(app.MyCmdPath(), cmdName+".go")
		//	createCmdFile(project.License(), cmdPath, cmdName)

		fmt.Fprintln(cmd.OutOrStdout(), cmdName, "created at", cmdPath)
	},
}
