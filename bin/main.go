package main

import (
	"fmt"
	"github.com/gofunct/stencil"
)

func tasks(p *stencil.Project) {
	p.Func("test", nil, func(c *stencil.Context) {
		c.Run("go test")
	})

	p.Func("test", stencil.Series{"build"}, func(c *stencil.Context) {
		c.Run("go test")
	})

	p.Func("dist", stencil.Series{"test", "lint"}, nil)

	p.Func("install", nil, func(c *stencil.Context) {
		c.Run("go get github.com/golang/lint/golint")
		c.Run("go get github.com/robertkrimen/godocdown")
	})

	p.Func("lint", nil, func(c *stencil.Context) {
		c.Run("golint .")
		c.Run("gofmt -w -s .")
		c.Run("go vet .")
	})

	p.Func("build", nil, func(c *stencil.Context) {
		c.Run("go install", stencil.M{"$in": "stencil"})
	})

	p.Func("interactive", nil, func(c *stencil.Context) {
		c.Bash(`
			echo name?
			read name
			echo hello $name
		`)
	})

	p.Func("whoami", nil, func(c *stencil.Context) {
		c.Run("whoami")
	})

	pass := 0
	p.Func("err2", nil, func(*stencil.Context) {
		if pass == 2 {
			stencil.Halt("oh oh")
		}
	})

	p.Func("err", stencil.Series{"err2"}, func(*stencil.Context) {
		pass++
		if pass == 1 {
			return
		}
		stencil.Halt("foo err")
	}).Src("test/*.txt")

	p.Func("hello", nil, func(c *stencil.Context) {
		name := c.Args.AsString("default value", "name", "n")
		fmt.Println("Hello", name)
	}).Src("*.hello").Wait(3000)

	p.Func("server", nil, func(c *stencil.Context) {
		c.Start("main.go", stencil.M{"$in": "example"})
	}).Src("example/**/*.go")

	p.Func("change-package", nil, func(c *stencil.Context) {
		// works on mac
		c.Run(`find . -name "*.go" -print | xargs sed -i "" 's|github.com/gofunct/stencil|github.com/gofunct/stencil|g'`)
		// maybe linux?
	})
}

func main() {
	stencil.Stencil(tasks)
}
