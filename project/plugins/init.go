package plugins

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Initf08871a4ac25f3ef36434d141f852b90f2f2bc45 = "coverage:\n  status:\n    project:\n      default:\n        threshold: 1%\n    patch: off\n"
var _Initf0083e572ff48d7111d178c5ccf0539aa0fcbf40 = "Permission is hereby granted, free of charge, to any person obtaining a copy\nof this software and associated documentation files (the \"Software\"), to deal\nin the Software without restriction, including without limitation the rights\nto use, copy, modify, merge, publish, distribute, sublicense, and/or sell\ncopies of the Software, and to permit persons to whom the Software is\nfurnished to do so, subject to the following conditions:\nThe above copyright notice and this permission notice shall be included in\nall copies or substantial portions of the Software.\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\nIMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\nFITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\nAUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\nLIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\nOUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\nTHE SOFTWARE.`,\n\t\tText: `The MIT License (MIT)\n{{ .copyright }}\nPermission is hereby granted, free of charge, to any person obtaining a copy\nof this software and associated documentation files (the \"Software\"), to deal\nin the Software without restriction, including without limitation the rights\nto use, copy, modify, merge, publish, distribute, sublicense, and/or sell\ncopies of the Software, and to permit persons to whom the Software is\nfurnished to do so, subject to the following conditions:\nThe above copyright notice and this permission notice shall be included in\nall copies or substantial portions of the Software.\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\nIMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\nFITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\nAUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\nLIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\nOUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\nTHE SOFTWARE."
var _Init2f8b03d8b9e9a1bf72da6ea9424da9a231fef8c2 = ".travis.yml\n.reviewdog.yml\n_tests\nbin\nvendor\nREADME.md\nLICENSE\n"
var _Init4b73f9b96c39121a8cbcb9892e9de616bc945621 = "package = \"{{.packageName}}\"\n\n[gogen]\nserver_dir = \"./app/server\"\n\n[protoc]\nprotos_dir = \"./api/protos\"\nout_dir = \"./api\"\nimport_dirs = [\n  \"./api/protos\",\n  \"./vendor/github.com/grpc-ecosystem/grpc-gateway\",\n  \"./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis\",\n]\n\n  [[protoc.plugins]]\n  name = \"go\"\n  args = { plugins = \"grpc\", paths = \"source_relative\" }\n\n  [[protoc.plugins]]\n  name = \"grpc-gateway\"\n  args = { logtostderr = true, paths = \"source_relative\" }\n\n  [[protoc.plugins]]\n  name = \"swagger\"\n  args = { logtostderr = true }\n"
var _Init560bfd3a40b63143fcc3be11ec4a6978feb0fd14 = "{{comment .copyright}}\n{{if .license}}{{comment .license}}{{end}}\npackage cmd\n\nimport (\n\t\"fmt\"\n\t\"os\"\n\t\"github.com/spf13/cobra\"\n\t\"github.com/spf13/viper\"\n\t\"github.com/gofunct/gogen/gocloud\"\n\t\"github.com/gofunct/common/logging\"\n)\n\nvar (\n    cfgFile string\n    config = viper.New()\n    )\n\n// rootCmd represents the base command when called without any subcommands\nvar rootCmd = &cobra.Command{\n\tUse:   \"temp\",\n\tShort: \"A brief description of your application\",\n\tLong: `A longer description that spans multiple lines and likely contains\nexamples and usage of using your application. For example:\n\nStencil is a CLI library for Go that empowers applications.\nThis application is a tool to generate the needed files\nto quickly create a Stencil application.`,\n\t// Uncomment the following line if your bare application\n\t// has an action associated with it:\n\t//\tRun: func(cmd *cobra.Command, args []string) { },\n}\n\n// Execute adds all child commands to the root command and sets flags appropriately.\n// This is called by main.main(). It only needs to happen once to the rootCmd.\nfunc Execute() {\n\tif err := rootCmd.Execute(); err != nil {\n\t\tfmt.Println(err)\n\t\tos.Exit(1)\n\t}\n}\n\nfunc init() {\n\tcobra.OnInitialize(initConfig)\n\tlogging.AddFlags(rootCmd)\n\trootCmd.PersistentFlags().StringVar(&cfgFile, \"config\", \"\", \"config file (default is $HOME/.temp.yaml)\")\n\trootCmd.Flags().BoolP(\"toggle\", \"t\", false, \"Help message for toggle\")\n\trootCmd.AddCommand(gocloud.NewGocloudCommand)\n}\n\n// initConfig reads in config file and ENV variables if set.\nfunc initConfig() {\n\tconfig.AutomaticEnv()\n\tif cfgFile != \"\" {\n\t\t// Use config file from the flag.\n\t\tconfig.SetConfigFile(cfgFile)\n\t} else {\n\t\tconfig.AddConfigPath(\".\")\n        config.SetConfigName(\"config\")\n\n\t}\n\n\t// If a config file is found, read it in.\n\tif err := config.ReadInConfig(); err == nil {\n\t\tfmt.Println(\"Using config file:\", viper.ConfigFileUsed())\n\t}\n}\n"
var _Init23b808cac963edf44a497827f2a6eff5ddac970f = ""
var _Init4a622176ca163dbfa742f4811d5365ee115f3311 = "// +build tools\n\npackage main\n\n// tool dependencies\nimport (\n\t_ \"github.com/golang/mock/mockgen\"\n\t_ \"github.com/golang/protobuf/protoc-gen-go\"\n\t_ \"github.com/google/wire/cmd/wire\"\n\t_ \"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway\"\n\t_ \"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger\"\n\t_ \"github.com/haya14busa/reviewdog/cmd/reviewdog\"\n\t_ \"github.com/izumin5210/gex/cmd/gex\"\n\t_ \"github.com/jessevdk/go-assets-builder\"\n\t_ \"github.com/kisielk/errcheck\"\n\t_ \"github.com/mitchellh/gox\"\n\t_ \"github.com/srvc/wraperr/cmd/wraperr\"\n\t_ \"golang.org/x/lint/golint\"\n\t_ \"honnef.co/go/tools/cmd/megacheck\"\n\t_ \"mvdan.cc/unparam\"\n)\n"
var _Inita2e3eb52bd2fece8f799a66e5e157007908dfdf2 = "The MIT License (MIT)\n\nCopyright © 2019 {{.developer}} <{{.email}}>\n\nPermission is hereby granted, free of charge, to any person obtaining a copy\nof this software and associated documentation files (the \"Software\"), to deal\nin the Software without restriction, including without limitation the rights\nto use, copy, modify, merge, publish, distribute, sublicense, and/or sell\ncopies of the Software, and to permit persons to whom the Software is\nfurnished to do so, subject to the following conditions:\n\nThe above copyright notice and this permission notice shall be included in\nall copies or substantial portions of the Software.\n\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR\nIMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,\nFITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE\nAUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER\nLIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,\nOUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN\nTHE SOFTWARE.\n"
var _Inita3426f4af9db68f6bf4705d905eb9527e5207334 = "language: go\n\ngo: '1.11'\n\nenv:\n  global:\n  - DEP_RELEASE_TAG=v0.5.0\n  - FILE_TO_DEPLOY=\"dist/*\"\n\n  # GITHUB_TOKEN\n  - secure: \"MhmvXAAzOA5HY7koCfagX1wJi2mBVQsVF5cCMaNx73l+7uDgNzGYfTn4OGKmckduiGB/mp5bTJ1DeMbPq+TDX1n/RE6kndu/Q/1vw4pbxm9BsmO9b3DizIFoWlnG+EABdAZa9igbCAfv+Jj57a0WjKGaiLazylj1mb7AYj6Vao+1zvm2ufoZvpKJcnKPqcWTsx/enJD3wx0LbqTpN5a/EdynJF9kj9Z97cGk9lS/hQHqmYVUYLYG5ZIvPjkuc6ho6pYaerupZ8aQgwraupRrNAzh70C3QgxnrCK+6RRmBMchhBsHOZq1MGhbN48ttlSMKow2NyVp8mK8+wLUnQgxEvYjVNJBXf5iKMmCTBiTO8IqgAKkkMgLaB3H0UpkeOoUQNTACPxR42+FJcwObmxYRSekTGFPwAAwnZV/1BuPrpxpT7JHa9ELlShz2OVEDz9aK/WC28/oEmtYKN8s9koKr1sx4OT5c0F/XG+er2idgCWwvfK5A0Om7Fudur+bbp1a38QWb00cAu8dPTIONe01vGXQ04d+NyohS2bcvK3iehVpa+WZ4CHkjRRuv6vQGvFMNCtwwQjXopBM99+yAykLm7yqOewbzbxFI7nCHNBc1zHvI13j7yniEoI/vdWk43e2H3Az0OOtdVASNmmp5Avwo/UWzjVACvlyNK1CST4pqYQ=\"\n  # REVIEWDOG_GITHUB_API_TOKEN\n  - secure: \"HIpuAXhIivyVkMKnWucjuFWJcDnGsvBPm4lQmpCnDOWFFWgblhBzojqN9q0DK91Sc7MEeZPDD3yhZAUOYK2mcRthLZYhbblCjZsE742i8dVB9Y8+PiMb/CHRdERCQUNvQKo+fiXJ4QWE42zx9ehTkKRQHGZkHx8cQVgtSnTyMD2lxxBJWRHUQ8OS+v0yKZCmERisClccbcm37vBQQe3/n6RhuhzxIBlA5G5MEt2ig3noocMcRjApl4Qz3eV/IqVrNs8iQeSm87N3eVxuqMS07SxpOBDhyq6tlU0Ab3VD6peY8aiQxqKLCNU5w0yL5ap9jLiHAV4TDYblS7wUAJLabp+Qdj4/5C2di+jfyn1ZITcKJu8H4kAr8hZqQXpAIQ9K6e/SUztyTfVlsPl9BBO56mx4FB2ZN2voAiJSE4ZUzXyp+zIPk2eiWfclPKiPyvPgFDF0RPV0n/EQXybGoJaLgEnZ4Tx7n2WTWCnZROZkw8EuldIY60D0qJiYYTDfhk2W3XBZUJ4isqrTYKdAP/SGcBLPRmWA2/Aaq7XaP6oHa9+jPIkmhyIALtarWESRbwzWtstXXBjPaUSStZx/J/lvJW2gpmbt8e74GKEEOv9FiX2NOglwN5vPwl7ZErPMdlEeMjOx+HOIts4BPfYwtGFD7Ws0WI4oiSf4PXmuvvzjf4s=\"\n\ncache:\n  directories:\n  - $GOPATH/pkg/dep\n  - $GOPATH/pkg/mod\n  - $HOME/.cache/go-build\n\njobs:\n  include:\n  - name: lint\n    install: make setup\n    script: make lint\n    if: type = 'pull_request'\n\n  - &test\n    install: make setup\n    script: make test\n    if: type != 'pull_request'\n\n  - &test-e2e\n    name: \"E2E test (go 1.11.2)\"\n    language: bash\n    sudo: required\n    services:\n    - docker\n    env:\n    - GO_VERSION=1.11.2\n    script: make test-e2e\n    if: type != 'pull_request'\n\n  - <<: *test\n    name: \"coverage (go 1.11)\"\n    script: make cover\n    after_success: bash <(curl -s https://codecov.io/bash)\n\n  - <<: *test\n    go: master\n\n  - <<: *test\n    go: '1.10'\n\n  - <<: *test-e2e\n    name: \"E2E test (go 1.10.5)\"\n    env:\n    - GO_VERSION=1.10.5\n\n  - stage: deploy\n    install: make setup\n    script: make packages -j4\n    deploy:\n    - provider: releases\n      skip_cleanup: true\n      api_key: $GITHUB_TOKEN\n      file_glob: true\n      file: $FILE_TO_DEPLOY\n      on:\n        tags: true\n    if: type != 'pull_request'\n"
var _Init033f9d45db331f8a4c2cfc3abbf3009df9a1538a = "{{ comment .copyright }}\n{{if .license}}{{ comment .license }}{{end}}\npackage main\nimport \"{{ .importpath }}\"\nfunc main() {\n\tcmd.Execute()\n}"
var _Initf8f5781f016ec71243395dbbb7bb65f373df2373 = "runner:\n  golint:\n    cmd: golint $(go list ./... | grep -v /vendor/)\n    format: golint\n  govet:\n    cmd: go vet $(go list ./... | grep -v /vendor/)\n    format: govet\n  errcheck:\n    cmd: errcheck -asserts -ignoretests -blank ./...\n    errorformat:\n      - \"%f:%l:%c:%m\"\n  wraperr:\n    cmd: wraperr ./...\n    errorformat:\n      - \"%f:%l:%c:%m\"\n  megacheck:\n    cmd: megacheck ./...\n    errorformat:\n      - \"%f:%l:%c:%m\"\n  unparam:\n    cmd: unparam ./...\n    errorformat:\n      - \"%f:%l:%c: %m\"\n"
var _Init109bef6bea90b5fdea7ed03f3c2b9a8f3a039ed4 = ""
var _Init38e76c5db8962fa825cf2bd8b23a2dc985c4513e = "*.so\n/vendor\n/bin\n/tmp\n/.idea\n"
var _Init840caece5084c7325cd0babad92224d9ac8a13ce = "# {{.packageName}}\n\n* Developer: {{.developer}}\n* Email: {{.email}}\n* Language: Golang\n* Download: `go get github.com/replace-me`\n* Summary: {.summary}}\n## Table of Contents\n\n- [{{.packageName}}](#{{.packageName}})\n  * [Table of Contents](#table-of-contents)."

// Init returns go-assets FileSystem
var Init = assets.NewFileSystem(map[string][]string{"/cmd": []string{"root.go.tmpl", ".keep.tmpl"}, "/": []string{"License.tmpl", "Gopkg.toml.tmpl", "LICENSE", ".dockerignore.tmpl", ".gitignore.tmpl", ".codecov.yml.tmpl", ".travis.yml.tmpl", "main.go.tmpl", ".reviewdog.yml.tmpl", "README.md.tmpl", "tools.go.tmpl", "config.toml.tmpl"}}, map[string]*assets.File{
	"/LICENSE": &assets.File{
		Path:     "/LICENSE",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Inita2e3eb52bd2fece8f799a66e5e157007908dfdf2),
	}, "/.travis.yml.tmpl": &assets.File{
		Path:     "/.travis.yml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Inita3426f4af9db68f6bf4705d905eb9527e5207334),
	}, "/main.go.tmpl": &assets.File{
		Path:     "/main.go.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546638060, 1546638060975037869),
		Data:     []byte(_Init033f9d45db331f8a4c2cfc3abbf3009df9a1538a),
	}, "/.reviewdog.yml.tmpl": &assets.File{
		Path:     "/.reviewdog.yml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Initf8f5781f016ec71243395dbbb7bb65f373df2373),
	}, "/cmd/.keep.tmpl": &assets.File{
		Path:     "/cmd/.keep.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Init109bef6bea90b5fdea7ed03f3c2b9a8f3a039ed4),
	}, "/.gitignore.tmpl": &assets.File{
		Path:     "/.gitignore.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Init38e76c5db8962fa825cf2bd8b23a2dc985c4513e),
	}, "/README.md.tmpl": &assets.File{
		Path:     "/README.md.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Init840caece5084c7325cd0babad92224d9ac8a13ce),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546639869, 1546639869949801680),
		Data:     nil,
	}, "/License.tmpl": &assets.File{
		Path:     "/License.tmpl",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1546639869, 1546639869949134462),
		Data:     []byte(_Initf0083e572ff48d7111d178c5ccf0539aa0fcbf40),
	}, "/.dockerignore.tmpl": &assets.File{
		Path:     "/.dockerignore.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Init2f8b03d8b9e9a1bf72da6ea9424da9a231fef8c2),
	}, "/.codecov.yml.tmpl": &assets.File{
		Path:     "/.codecov.yml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Initf08871a4ac25f3ef36434d141f852b90f2f2bc45),
	}, "/cmd": &assets.File{
		Path:     "/cmd",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1546638060, 1546638060973449175),
		Data:     nil,
	}, "/cmd/root.go.tmpl": &assets.File{
		Path:     "/cmd/root.go.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546638060, 1546638060972864123),
		Data:     []byte(_Init560bfd3a40b63143fcc3be11ec4a6978feb0fd14),
	}, "/Gopkg.toml.tmpl": &assets.File{
		Path:     "/Gopkg.toml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Init23b808cac963edf44a497827f2a6eff5ddac970f),
	}, "/tools.go.tmpl": &assets.File{
		Path:     "/tools.go.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Init4a622176ca163dbfa742f4811d5365ee115f3311),
	}, "/config.toml.tmpl": &assets.File{
		Path:     "/config.toml.tmpl",
		FileMode: 0x1ed,
		Mtime:    time.Unix(1546568716, 1546568716000000000),
		Data:     []byte(_Init4b73f9b96c39121a8cbcb9892e9de616bc945621),
	}}, "")
