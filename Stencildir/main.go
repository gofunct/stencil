package main

import (
	"fmt"
	"github.com/gofunct/stencil/cmd/task"
)

func tasks(p *task.Project) {
	p.Task("test", nil, func(c *task.Context) {
		c.Run("go test")
	})

	p.Task("test", task.S{"build"}, func(c *task.Context) {
		c.Run("go test")
	})

	p.Task("dist", task.S{"test", "lint"}, nil)

	p.Task("install", nil, func(c *task.Context) {
		c.Run("go get github.com/golang/lint/golint")
		// Run("go get github.com/mgutz/goa")
		c.Run("go get github.com/robertkrimen/gotaskctaskwn")
	})

	p.Task("lint", nil, func(c *task.Context) {
		c.Run("golint .")
		c.Run("gofmt -w -s .")
		c.Run("go vet .")
	})

	p.Task("build", nil, func(c *task.Context) {
		c.Run("go install", task.M{"$in": "stencil"})
	})

	p.Task("interactive", nil, func(c *task.Context) {
		c.Bash(`
			echo name?
			read name
			echo hello $name
		`)
	})

	p.Task("whoami", nil, func(c *task.Context) {
		c.Run("whoami")
	})

	pass := 0
	p.Task("err2", nil, func(*task.Context) {
		if pass == 2 {
			task.Halt("oh oh")
		}
	})

	p.Task("err", task.S{"err2"}, func(*task.Context) {
		pass++
		if pass == 1 {
			return
		}
		task.Halt("foo err")
	}).Src("test/*.txt")

	p.Task("hello", nil, func(c *task.Context) {
		name := c.Args.AsString("default value", "name", "n")
		fmt.Println("Hello", name)
	}).Src("*.hello").Debounce(3000)

	p.Task("server", nil, func(c *task.Context) {
		c.Start("main.go", task.M{"$in": "example"})
	}).Src("example/**/*.go")

	p.Task("change-package", nil, func(c *task.Context) {
		// works on mac
		c.Run(`find . -name "*.go" -print | xargs sed -i "" 's|github.com/gofunct/stencil|github.com/gofunct/stencil|g'`)
		// maybe linux?
		//Run(`find . -name "*.go" -print | xargs sed -i 's|gopkg.in/stencil.v1|github.com/gofunct/stencil|g'`)
	})
}

func main() {
	task.Stencil(tasks)
}