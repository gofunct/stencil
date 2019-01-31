package stencil

import (
	"bytes"
	"github.com/gofunct/common/pkg/logger/zap"
	"github.com/gofunct/gofs"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

var (
	homeDir, _ = homedir.Dir()
	Viper      = viper.New()
	sys        = &afero.Afero{
		Fs: afero.NewOsFs(),
	}
)

type tfItem struct {
	Sensitive bool
	Type      string
	Value     string
}

var Configuration = &Config{}

// Config contains service configuration
type Config struct {
	Author      string   `json:"author"`
	Github      string   `json:"github"`
	App         string   `json:"app"`
	Version     string   `json:"version"`
	Description string   `json:"description"`
	Bin         []string `json:"bin"`
	Templates   string   `json:"source.templates"`
	Proto       string   `json:"source.proto"`
	State       struct {
		Project          tfItem `json:"project"`
		ClusterName      tfItem `json:"cluster_name"`
		ClusterZone      tfItem `json:"cluster_zone"`
		Bucket           tfItem `json:"bucket"`
		DatabaseInstance tfItem `json:"database_instance"`
		DatabaseRegion   tfItem `json:"database_region"`
		RUnVarConfig     tfItem `json:"run_var_config"`
		RunVarName       tfItem `json:"run_var_name"`
	}
}

func init() {
	Viper.SetConfigName(".stencil")
	Viper.AllowEmptyEnv(true)
	Viper.SetConfigType("yaml")
	Viper.AddConfigPath(".")
	Viper.AutomaticEnv()
	zap.LogE("Writing config", Viper.WriteConfig())
	{
		Viper.SetDefault("author", "Coleman Word")
		Viper.SetDefault("github", path.Join(filepath.Join(os.Getenv("PWD"))))
		Viper.SetDefault("app", "sample")
		Viper.SetDefault("version", "v0.0.1")
		Viper.SetDefault("description", "change me to a helpful application description")
		Viper.SetDefault("bin", []string{
			"github.com/robertkrimen/godocdown/godocdown",
			"github.com/golang/lint/golint",
			"github.com/kisielk/errcheck",
			"mvdan.cc/unparam",
			"honnef.co/go/tools/cmd/megacheck",
			"github.com/golang/mock/mockgen",
			"github.com/google/wire/cmd/wire",
			"github.com/jessevdk/go-assets-builder",
			"github.com/mitchellh/gox",
			"github.com/srvc/wraperr/cmd/wraperr",
			"github.com/golang/protobuf/protoc-gen-go",
		})
		Viper.SetDefault("templates", os.Getenv("PWD")+"/templates")
		Viper.SetDefault("proto", os.Getenv("PWD")+"/proto")
		Viper.SetDefault("bucket", "default")
		Viper.SetDefault("projcet", "")
	}

	{
		Viper.SetDefault("os.runtime.goarch", runtime.GOARCH)
		Viper.SetDefault("os.runtime.compiler", runtime.Compiler)
		Viper.SetDefault("os.runtime.version", runtime.Version())
		Viper.SetDefault("os.runtime.goos", runtime.GOOS)
	}
	{
		Viper.SetDefault("env", os.Environ())
	}
	if gofs.FileExists("deploy/*.tf") {
		bytess, err := ShellOutput("terraform", "output", "-state", "deploy", "-json")
		zap.LogF("terraform json output", err)
		zap.LogF("terraform state to viper config", Viper.ReadConfig(bytes.NewBuffer(bytess)))
	}
	zap.LogF("unmarshaling config", Viper.Unmarshal(Configuration))
	zap.LogE("write config if no exist", Viper.WriteConfig())
	zap.LogE("Reading config", Viper.ReadInConfig())
	zap.Debug("Current config file-->", "config", Viper.ConfigFileUsed())
}

func GetConfig() *Config {
	zap.LogF("unmarshaling config", Viper.Unmarshal(Configuration))
	return Configuration
}
