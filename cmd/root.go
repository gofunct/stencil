/*
Add runtime to cobras project type; use runtime.methods and delete cobra.Projects.methods;
replace cobra.cmds templates with assetsFS var;
pass runtime.data to cobra.cmds template.Execute functions
*/
package cmd

import (
	"fmt"
	"github.com/gofunct/stencil/pkg/runtime"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   "stencil",
		Short: "A generator for Stencil based Applications",
	}
	appName string
	app     *runtime.App
)

// Execute executes the root command.
func Execute() {
	rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.PersistentFlags().StringVarP(&appName, "app", "a", "", "project name")
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
