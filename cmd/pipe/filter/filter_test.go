package filter

import (
	//"log"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/gofunct/stencil/cmd/pipe"
	"github.com/mgutz/str"
)

func newAssetText(s, writePath string) *pipe.Asset {
	asst := &pipe.Asset{WritePath: writePath}
	asst.WriteString(s)
	return asst
}

func egAsset(asst *pipe.Asset, filter func(asset *pipe.Asset) error) {
	filter(asst)
	fmt.Println(asst.String())
}

func TestAddHeader(t *testing.T) {
	asst := &pipe.Asset{}
	asst.WriteString("foo")
	filter := AddHeader("bar")
	filter(asst)
	if asst.String() != "barfoo" {
		t.Error("should have prepended bar")
	}

	// try add again
	filter(asst)
	if asst.String() != "barfoo" {
		t.Error("should be idempotent")
	}
}

func TestLoad(t *testing.T) {
	pipe, _ := pipe.Pipe()
	batcher := Load("test/**/*.txt")
	batcher(pipe)

	if len(pipe.Assets) != 2 {
		t.Error("should have loaded two test files")
	}

	result := ""
	for _, asset := range pipe.Assets {
		result += asset.String() + " "
	}
	if !(strings.Contains(result, "1") && strings.Contains(result, "2.txt")) {
		t.Errorf("should have loaded content %s", result)
	}
}

func TestReplaceLeft(t *testing.T) {
	asset := &pipe.Asset{}
	asset.WritePath = "views/index.go"
	filter := ReplacePath("views/", "test/")
	filter(asset)
	if asset.WritePath != "test/index.go" {
		t.Error("should have replaced subpath")
	}
}

func TestWrite(t *testing.T) {
	os.RemoveAll("tmp")
	assets := []*pipe.Asset{
		{WritePath: "tmp/foo.txt", Buffer: *bytes.NewBufferString("foo")},
		{WritePath: "tmp/bar.txt", Buffer: *bytes.NewBufferString("bar")},
	}
	filter := Write()
	filter(assets)
	dat, _ := ioutil.ReadFile("tmp/foo.txt")
	if string(dat) != "foo" {
		t.Error("should have written foo.txt")
	}
	os.RemoveAll("tmp")
}

func TestCat(t *testing.T) {
	pi, _ := pipe.Pipe(
		Load("test/**/*.txt"),
		Cat(";", "dist/cat.txt"),
	)

	if len(pi.Assets) != 1 {
		t.Errorf("should only have 1 asset %+v\n", pi.Assets)
	}

	s := str.Clean(pi.Assets[0].String())
	if !strings.Contains(s, ";2.txt") {
		t.Errorf("should join: %+v\n", s)
	}
	os.RemoveAll("dist")
}

func ExampleReplacePattern() {
	egAsset(newAssetText("abcdef", ""), ReplacePattern(`abc`, "x"))
	// Output:
	// xdef
}
