package plugins

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Certs1cd93ef550bd93ecebe4131eaf715a1a193ce1b4 = "gen: # generate a server.key and server.pem file\n\topenssl genrsa -out server.key 2048\n\topenssl req -new -x509 -key server.key -out server.pem -days 3650\n\nhelp: ## help\n\t@awk 'BEGIN {FS = \":.*?## \"} /^[a-zA-Z_-]+:.*?## / {printf \"\\033[36m%-30s\\033[0m %s\\n\", $$1, $$2}' $(MAKEFILE_LIST) | sort"
var _Certsaa5aa80764cfb35580f57683c886d2b9d4202e56 = "*.key\n*.pem"

// Certs returns go-assets FileSystem
var Certs = assets.NewFileSystem(map[string][]string{"/": []string{"Makefile.tmpl", ".gitignore.tmpl"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     nil,
	}, "/Makefile.tmpl": &assets.File{
		Path:     "/Makefile.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Certs1cd93ef550bd93ecebe4131eaf715a1a193ce1b4),
	}, "/.gitignore.tmpl": &assets.File{
		Path:     "/.gitignore.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Certsaa5aa80764cfb35580f57683c886d2b9d4202e56),
	}}, "")
