package root_test

import (
	"fmt"
	"context"
	"github.com/gofunct/stencil/runtime/root"
)

func ExampleChain() {
	e := root.Chain(
		annotate("first"),
		annotate("second"),
		annotate("third"),
	)(myRoot)

	if _, err := e(ctx, req); err != nil {
		panic(err)
	}
}

var (
	ctx = context.Background()
	req = struct{}{}
)

func annotate(s string) root.Wrapper {
	return func(next root.Root) root.Root {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			fmt.Println(s, "pre")
			defer fmt.Println(s, "post")
			return next(ctx, request)
		}
	}
}

func myRoot(context.Context, interface{}) (interface{}, error) {
	fmt.Println("my endpoint!")
	return struct{}{}, nil
}
