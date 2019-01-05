![](https://github.com/gofunct/common/blob/master/logo/dark_logo_transparent_background.png?raw=true)

# Stencil

* Author: Coleman Word 
* Email: coleman.word@gofunct.com
* Language: Golang
* Download: `go get github.com/gofunct/stencil/...`
* Description: Golang funcmaps, virtual file systems, and code generation utilities
* Status: WIP

## Statement of Need
 I spend too much time writing boilerplate instead of designing apis and writing business logic
 
## Usage
```commandline

  _____ _                  _ _ 
 / ____| |                (_) |
| (___ | |_ ___ _ __   ___ _| |
 \___ \| __/ _ \ '_ \ / __| | |
 ____) | ||  __/ | | | (__| | |
|_____/ \__\___|_| |_|\___|_|_|
                               
By: Coleman Word
Email: coleman.word@gofunct.com

Usage:
  stencil [command]

Available Commands:
  config      A brief description of your command
  help        Help about any command
  init        Initialize a Stencil Application
  issues      a list of current issues with stencil
  roadmap     a list of goals for stencil
  version     print the version of stencil

Flags:
  -a, --ask                    fill out a survey to set configs
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
      --github string          what is the github user name of this project (default "gofunct")
  -h, --help                   help for stencil
      --image string           what is the docker base image of this project (default "alpine:3.8")
      --lis string             what is the server listen port of this project (default ":8080")
      --log-level string       what is the log level(debug or verbose) of this project (default "verbose")
      --logo string            what is the logo of this project (default "https://github.com/gofunct/common/blob/master/logo/dark_logo_transparent_background.png?raw=true")
      --plugins string         what is the config file of this project (default "init")
      --runvar string          what is the runtime variable of this project
      --runvar-config string   what is the runtime variable config name of this project
      --runvar-wait duration   what is the runtime variable timeout of this project (default 5s)
      --scope string           what is the scope of this project (default "not available")
      --services strings       what are the grpc service names of this project
      --tools strings          what are the third party dev tools of this project (default [gex,reviewdog,errcheck,gogoslick,grpc-swagger,grpc-gateway,protocgen-go,golint,megacheck,wire,unparam,gox,gotests])
      --version string         what is the version of this project (default "v0.1.1")
  -w, --write                  fill out a survey to set configs

Use "stencil [command] --help" for more information about a command.

``` 

## Project Variables (viper, cobra, env)

<details><summary>show</summary>
<p>
 
- [x] author
- [x] email
- [x] description
- [x] dependencies
- [x] serviceNames
- [x] gitUser
- [x] dockerUser
- [x] listenPort
- [x] bucket            
- [x] dbHost            
- [x] dbName            
- [x] dbUser            
- [x] dbPassword        
- [x] runVar            
- [x] runVarWaitTime    
- [x] listen            
- [x] cloudSQLRegion    
- [x] runtimeConfigName
- [x] baseImage 

</p>
</details>

## Issues

<details><summary>show</summary>
<p>

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

</p>
</details>

## Func Map

<details><summary>show</summary>
<p>
 
 - Manipulation
 
 ```text
String Functions: trim, wrap, randAlpha, plural, etc.
String List Functions: splitList, sortAlpha, etc.
Math Functions: add, max, mul, etc.
Integer Slice Functions: until, untilStep
Date Functions: now, date, etc.
Defaults Functions: default, empty, coalesce, toJson, toPrettyJson
Encoding Functions: b64enc, b64dec, etc.
Lists and List Functions: list, first, uniq, etc.
Dictionaries and Dict Functions: dict, hasKey, pluck, etc.
Type Conversion Functions: atoi, int64, toString, etc.
File Path Functions: base, dir, ext, clean, isAbs
Flow Control Functions: fail
Advanced Functions
UUID Functions: uuidv4
OS Functions: env, expandenv
Version Comparison Functions: semver, semverCompare
Reflection: typeOf, kindIs, typeIsLike, etc.
Cryptographic and Security Functions: derivePassword, sha256sum, genPrivateKey
 
 ```
 
 
-Strings
```text
{{trim "   hello    "}}:                                                            hello
{{trimAll "$" "$5.00"}}:                                                            5.00
{{trimSuffix "-" "hello-"}}:                                                        hello
{{upper "hello"}}:                                                                  HELLO
{{lower "HELLO"}}:                                                                  hello
{{title "hello world"}}:                                                            Hello World
{{untitle "Hello World"}}:                                                          hello world
{{repeat 3 "hello"}}:                                                               hellohellohello
{{substr 0 5 "hello world"}}:                                                       hello
{{nospace "hello w o r l d"}}:                                                      helloworld
{{trunc 5 "hello world"}}:                                                          hello
{{abbrev 5 "hello world"}}:                                                         he...
{{abbrevboth 5 10 "1234 5678 9123"}}:                                               ...5678...
{{initials "First Try"}}:                                                           FT
{{randNumeric 3}}:                                                                  528
{{- /*{{wrap 80 $someText}}*/}}:
{{wrapWith 5 "\t" "Hello World"}}:                                                  Hello	World
{{contains "cat" "catch"}}:                                                         true
{{hasPrefix "cat" "catch"}}:                                                        true
{{cat "hello" "beautiful" "world"}}:                                                hello beautiful world
{{- /*{{indent 4 $lots_of_text}}*/}}:
{{- /*{{indent 4 $lots_of_text}}*/}}:
{{"I Am Henry VIII" | replace " " "-"}}:                                            I-Am-Henry-VIII
{{len .Service.Method | plural "one anchovy" "many anchovies"}}:                    many anchovies
{{snakecase "FirstName"}}:                                                          first_name
{{camelcase "http_server"}}:                                                        HttpServer
{{shuffle "hello"}}:                                                                holle
{{regexMatch "[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}" "test@acme.com"}}:   true
{{- /*{{regexFindAll "[2,4,6,8]" "123456789"}}*/}}:
{{regexFind "[a-zA-Z][1-9]" "abcd1234"}}:                                           d1
{{regexReplaceAll "a(x*)b" "-ab-axxb-" "${1}W"}}:                                   -W-xxW-
{{regexReplaceAllLiteral "a(x*)b" "-ab-axxb-" "${1}"}}:                             -${1}-${1}-
{{regexSplit "z+" "pizza" -1}}:                                                     [pi a]

# Get one specific method on array method using index
{{ index .Service.Method 1 }}:                                                      name:"Iii" input_type:".dummy.Dummy2" output_type:".dummy.Dummy1" options:<> 

# Sprig: advanced
{{if contains "cat" "catch"}}yes{{else}}no{{end}}:   yes
{{1 | plural "one anchovy" "many anchovies"}}:       one anchovy
{{2 | plural "one anchovy" "many anchovies"}}:       many anchovies
{{3 | plural "one anchovy" "many anchovies"}}:       many anchovies

```

- Protoc

```text
# Common variables
{{.File.Name}}:                                                                           helpers.proto
{{.File.Name | upper}}:                                                                   HELPERS.PROTO
{{.File.Package | base | replace "." "-"}}                                                dummy
{{$packageDir := .File.Name | dir}}{{$packageDir}}                                        .
{{$packageName := .File.Name | base | replace ".proto" ""}}{{$packageName}}               helpers
{{$packageImport := .File.Package | replace "." "_"}}{{$packageImport}}                   dummy
{{$namespacedPackage := .File.Package}}{{$namespacedPackage}}                             dummy
{{$currentFile := .File.Name | getProtoFile}}{{$currentFile}}                             <nil>
{{- /*{{- $currentPackageName := $currentFile.GoPkg.Name}}{{$currentPackageName}}*/}}
```

</p>
</details>

## File Tree

<details><summary>show</summary>
<p>

```commandline
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── bin
│   └── doc.go
├── cmd
│   ├── config.go
│   ├── init.go
│   └── root.go
├── doc.go
├── docs
│   └── README.md
├── funcmap
│   ├── camel.go
│   ├── crypto.go
│   ├── crypto_test.go
│   ├── date.go
│   ├── date_test.go
│   ├── defaults.go
│   ├── defaults_test.go
│   ├── dict.go
│   ├── dict_test.go
│   ├── functions.go
│   ├── functions_test.go
│   ├── list.go
│   ├── list_test.go
│   ├── numbers.go
│   ├── numbers_test.go
│   ├── reflect.go
│   ├── reflect_test.go
│   ├── regex.go
│   ├── regex_test.go
│   ├── snake.go
│   ├── strings.go
│   ├── strings_test.go
│   └── viper.go
├── go.mod
├── go.sum
├── gocloud
│   ├── app.go
│   ├── aws
│   │   ├── aws.go
│   │   ├── blob.go
│   │   ├── runtimevar.go
│   │   └── user.go
│   ├── bucket.go
│   ├── flags.go
│   ├── google
│   │   ├── app.go
│   │   ├── blob.go
│   │   ├── db.go
│   │   ├── gcloud.go
│   │   ├── kube.go
│   │   ├── run.go
│   │   ├── runtime_config.go
│   │   └── user.go
│   ├── healthcheck.go
│   ├── inject_aws.go
│   ├── inject_gcp.go
│   ├── inject_local.go
│   └── wire_gen.go
├── hack
│   └── doc.go
├── main.go
├── project
│   ├── ask.go
│   ├── check.go
│   ├── convert.go
│   ├── create.go
│   ├── delete.go
│   ├── find.go
│   ├── gen.go
│   ├── generator.go
│   ├── http
│   │   └── httpvfs.go
│   ├── init.go
│   ├── join.go
│   ├── license.go
│   ├── match.go
│   ├── parse.go
│   ├── plugins
│   │   ├── certs
│   │   │   └── Makefile.tmpl
│   │   ├── gen.go
│   │   ├── grpc
│   │   │   ├── cmd
│   │   │   │   └── run.go.tmpl
│   │   │   ├── protos
│   │   │   │   └── service.proto.tmpl
│   │   │   ├── register.go.tmpl
│   │   │   ├── run.go.tmpl
│   │   │   ├── server.go.tmpl
│   │   │   └── server_test.go.tmpl
│   │   ├── init
│   │   │   ├── Gopkg.toml.tmpl
│   │   │   ├── LICENSE
│   │   │   ├── License.tmpl
│   │   │   ├── README.md.tmpl
│   │   │   ├── cmd
│   │   │   │   └── root.go.tmpl
│   │   │   ├── config.toml.tmpl
│   │   │   ├── main.go.tmpl
│   │   │   └── tools.go.tmpl
│   │   └── init.go
│   ├── project.go
│   ├── project_test.go
│   ├── testing
│   │   ├── generator_mock.go
│   │   └── script_mock.go
│   ├── trim.go
│   ├── types.go
│   ├── utils.go
│   └── write.go
├── temp
│   └── doc.go
├── tools.go
```
</p>
</details>

