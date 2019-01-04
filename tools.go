// +build tools

package tools

// tool dependencies
import (
	_ "github.com/golang/mock/mockgen"
	_ "github.com/google/wire/cmd/wire"
	_ "github.com/haya14busa/reviewdog/cmd/reviewdog"
	_ "github.com/jessevdk/go-assets-builder"
	_ "github.com/kisielk/errcheck"
	_ "github.com/mitchellh/gox"
	_ "github.com/srvc/wraperr/cmd/wraperr"
	_ "golang.org/x/lint/golint"
	_ "honnef.co/go/tools/cmd/megacheck"
	_ "mvdan.cc/unparam"
)
