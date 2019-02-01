package application

import (
	"fmt"
	"github.com/gofunct/stencil/app"
	"github.com/pkg/errors"
)

var myApp = &app.App{
	AppFuncs: []app.AppFunc{
		func(i *app.App) error {
			fmt.Println("hello from func 1")
			return nil
		},

		func(i *app.App) error {
			fmt.Println("hello from func 2")
			return nil
		},

		func(i *app.App) error {
			fmt.Println("hello from func 3")
			return nil
		},
		func(i *app.App) error {
			fmt.Println("going to debug now...")
			i.Debug()
			return nil
		},
	},
}

func Execute() {
	if err := myApp.Execute(); err != nil {
		fmt.Printf("%s\n%v\n", err.Error(), errors.WithStack(err))
	}
}
