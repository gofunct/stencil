package main

import (
	"github.com/gofunct/stencil"
	"github.com/gofunct/stencil/filter"
	"github.com/mgutz/str"
)

func tasks(p *stencil.Project) {
	p.Func("test", nil, func(c *stencil.Context) {
		c.Run("go test")
	})

	p.Func("test", stencil.S{"build"}, func(c *stencil.Context) {
		c.Run("go test")
	})

	p.Func("dist", stencil.S{"test", "lint"}, nil)

	p.Func("install", nil, func(c *stencil.Context) {
		c.Run("go get github.com/golang/lint/golint")
		// Run("go get github.com/mgutz/goa")
		c.Run("go get github.com/robertkrimen/godocdown/godocdown")
	})

	p.Func("build", nil, func(c *stencil.Context) {
		c.Run("go install", stencil.M{"$in": "./..."})
	})

	p.Func("lint", nil, func(c *stencil.Context) {
		c.Run("golint .")
		c.Run("gofmt -w -s ./...")
		c.Run("go vet ./...")
	})

	p.Func("push", stencil.S{"test", "lint", "build"}, func(c *stencil.Context) {
		c.Run("git add .")
		c.Run(`git commit -m "default git commit"`)
		c.Run("git push origin master")
	})

	p.Func("readme", stencil.S{"install"}, func(c *stencil.Context) {
		c.Run("godocdown -o README.md")
		c.Pipe(
			filter.Load("./README.md"),
			filter.Str(str.ReplaceF("--", "\n[stencil](https://github.com/stencil.com)\n", 1)),
			filter.Write(),
		)
	})
}

func main() {
	stencil.Stencil(tasks)
}
