![](https://github.com/gofunct/common/blob/master/logo/dark_logo_transparent_background.png?raw=true)

# Stencil

* Author: Coleman Word 
* Email: coleman.word@gofunct.com
* Language: Golang
* Download: `go get github.com/gofunct/stencil/...`
* Description: Golang funcmaps, virtual file systems, and code generation utilities
* Status: WIP

## Table of Contents
- [Stencil](#Stencil)
  * [Table of Contents](#Table-of-Contents)
  * [Statement of Need](#Statement of Need)
  * [Project Roadmap](#Project Roadmap)
  * [Project Roadmap](#Issues)
  * [File Tree](#File-Tree)

## Statement of Need
 I spend too much time writing boilerplate instead of designing apis and writing business logic
 
## Usage
```commandline
Usage:
  stencil [command]

Available Commands:
  config      A brief description of your command
  help        Help about any command

Flags:
      --aliases strings        what are the aliases of this project
      --ask                    fill out a survey to set configs
      --author string          what is the developer full name of this project (default "Coleman Word")
      --bucket string          what is the gcloud storage bucket of this project (default "gofunct-storage")
      --cloud-region string    what is the gcloud region of this project
      --config string          what is the config file of this project
      --db-host string         what is the datbase host of this project
      --db-name string         what is the datbase name of this project (default "gofunct-db")
      --db-pw string           what is the database password of this project (default "admin")
      --db-user string         what is the database user name of this project (default "gofunct")
      --dockerhub string       what is the dockerhub user name of this project
      --email string           what is the developer email address of this project (default "coleman.word@gofunct.com")
      --example string         what is the example usage of this project
      --github string          what is the github user name of this project (default "gofunct")
  -h, --help                   help for stencil
      --image string           what is the docker base image of this project (default "alpine:3.8")
      --lis string             what is the server listen port of this project (default ":8080")
      --log-level string       what is the log level(debug or verbose) of this project (default "verbose")
      --logo string            what is the logo of this project (default "https://github.com/gofunct/common/blob/master/logo/dark_logo_transparent_background.png?raw=true")
      --runvar string          what is the runtime variable of this project
      --runvar-config string   what is the runtime variable config name of this project
      --runvar-wait duration   what is the runtime variable timeout of this project (default 5s)
      --scope string           what is the scope of this project (default "not available")
      --services strings       what are the grpc service names of this project
      --tools strings          what are the third party dev tools of this project (default [gex,reviewdog,errcheck,gogoslick,grpc-swagger,grpc-gateway,protocgen-go,golint,megacheck,wire,unparam,gox,gotests])
      --version string         what is the version of this project (default "v0.1.1")
      --write                  fill out a survey to set configs

Use "stencil [command] --help" for more information about a command.

``` 
## Default Config File
```yaml
aliases: []
author: Coleman Word
bucket: gofunct-storage
cloud-region: ""
config: ""
db-host: ""
db-name: gofunct-db
db-pw: gofunct-db
db-user: gofunct-db
dockerhub: ""
email: coleman.word@gofunct.com
example: ""
github: gofunct
image: alpine:3.8
lis: :8080
log-level: verbose
logo: https://github.com/gofunct/common/blob/master/logo/dark_logo_transparent_background.png?raw=true
runvar: ""
runvar-config: ""
runvar-wait: 5s
scope: not available
services: []
tools:
- gex
- reviewdog
- errcheck
- gogoslick
- grpc-swagger
- grpc-gateway
- protocgen-go
- golint
- megacheck
- wire
- unparam
- gox
- gotests
version: v0.1.1

```
## Project Roadmap

- [ ] submit beta-ready realease v0.1.1
- [ ] pass variables to funcmap through config file, flags, or env(viper, cobra)
- [ ] embed template assets in virtual os with go-asset-builder and afero
- [ ] create custom func map that supports file system methods
- [ ] create plugins architecture to support extensibility
- [ ] create cli utility for common template executions

### Template Variables (viper, cobra, env)
- [ ] author
- [ ] email
- [ ] languages
- [ ] description
- [ ] dependencies
- [ ] serviceNames
- [ ] gitUser
- [ ] dockerUser
- [ ] listenPort
- [ ] bucket            
- [ ] dbHost            
- [ ] dbName            
- [ ] dbUser            
- [ ] dbPassword        
- [ ] runVar            
- [ ] runVarWaitTime    
- [ ] listen            
- [ ] cloudSQLRegion    
- [ ] runtimeConfigName
- [ ] baseImage 

### Func Map(methods)
- [ ] Gopath
- [ ] ImportDir
- [ ] PackageDir
- [ ] CmdDir
- [ ] RootDir
- [ ] VendorDir
- [ ] ProtoDir
- [ ] CamelCase
- [ ] Proto Methods
- [ ] Proto Services
- [ ] Download(go get ...)
- [ ] GoDocs
- [ ] Project File Tree

### File Generation
- [ ] cmd/root.go(cobra): root cobra command
- [ ] cmd/run.go(cobra): start a stackdriver server
- [ ] deploy/: main.tf, vars.tf, output.tf, deployment.yaml(kube manifest)
- [ ] certs/: makefile and .gitignore to generate certs
- [ ] proto/: protobuf file with gogo & gateway annotations
- [ ] bin/: third party binaries
- [ ] main.go: executes root command
- [ ] Dockerfile
- [ ] Makefile
- [ ] README.md
- [ ] CONTRIBUTING.md
- [ ] MIT LICENSE
- [ ] GoDocs2md
- [ ] reviewdog.yaml
- [ ] config.yaml
- [ ] .gitignore
- [ ] .dockerignore
- [ ] .gcloudignore
- [ ] .tools.go: 3rd party binary imports
- [ ] homebrew formulae

## Issues

- [ ] Not Done
- [ ] 
- [ ]
- [ ]
- [ ]
- [ ]
- [ ]
- [ ]
- [ ]
- [ ]
- [ ]
- [ ]
- [ ]

## File Tree

```commandline
└── README.md

```


