package runtime

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"path/filepath"
	"time"
)

type Config struct {
	ENV     Env     `mapstructure:",squash"`
	PROJECT Project `mapstructure:",squash"`
	CLOUD   Cloud   `mapstructure:",squash"`
}

type Env struct {
	PWD    string `mapstructure:"pwd"`
	Shell  string `mapstructure:"shell"`
	HOME   string `mapstructure:"home"`
	PATH   string `mapstructure:"path"`
	GOPATH string `mapstructure:"gopath"`
	GOROOT string `mapstructure:"goroot"`
	USER   string `mapstructure:"user"`
	GOBIN  string `mapstructure:"gobin"`
	GOSRC  string `mapstructure:"gosrc"`
}

type Project struct {
	AppName   string   `mapstructure:"pwd"`
	Logo      string   `mapstructure:"logo"`
	Email     string   `mapstructure:"email"`
	Summary   string   `mapstructure:"summary"`
	Tools     []string `mapstructure:"tools"`
	Version   string   `mapstructure:"version"`
	LogLevel  string   `mapstructure:"logLevel"`
	Services  []string `mapstructure:"services"`
	Github    string   `mapstructure:"github"`
	Dockerhub string   `mapstructure:"dockerhub"`
	Domain    string   `mapstructure:"domain"`
}

type Cloud struct {
	Lis          string        `mapstructure:"lis"`
	Bucket       string        `mapstructure:"bucket"`
	DbHost       string        `mapstructure:"dbHost"`
	DbPassword   string        `mapstructure:"dbPw"`
	DbUser       string        `mapstructure:"dbUser"`
	DbName       string        `mapstructure:"dbName"`
	Runvar       string        `mapstructure:"runvar"`
	RunVarWait   time.Duration `mapstructure:"runvarWait"`
	Cloud        string        `mapstructure:"cloud"`
	Region       string        `mapstructure:"region"`
	RunvarConfig string        `mapstructure:"runvarConfig"`
	Image        string        `mapstructure:"image"`
}

func (u *UI) InitConfig() {
	u.V = viper.New()
	u.V.AddConfigPath(".")
	u.V.AutomaticEnv()
	u.V.SetConfigName("stencil")
	u.V.SetConfigFile("stencil.yaml")
	u.V.BindEnv("GO111MODULE")
	u.V.BindEnv("PWD")
	u.V.BindEnv("HOME")
	u.V.BindEnv("GOROOT")
	u.V.BindEnv("GOPATH")
	u.V.BindEnv("USER")
	u.V.BindEnv("SHELL")
	src := filepath.Join(u.V.GetString("GOPATH"), "src")
	u.V.Set("gosrc", src)
	bin := filepath.Join(u.V.GetString("GOPATH"), "bin")
	u.V.Set("gobin", bin)
	imp, err := filepath.Rel(u.V.GetString("gosrc"), u.V.GetString("pwd"))
	if err != nil {
		u.Z.Fatal("faied to find import path", zap.Error(err))
	}
	u.V.Set("import", imp)
	u.V.Set("download", "go get "+imp)

	if err := u.V.Unmarshal(u.Config); err != nil {
		u.Z.Fatal(Red("failed to unmarshal config"), zap.Error(err))
	}
	u.WriteConfig()
	u.V.WatchConfig()
	u.V.OnConfigChange(func(e fsnotify.Event) {
		u.Z.Debug(Blue("Config file changed:"), zap.String("event", e.Name), zap.String("stringer", e.String()))
		u.WriteConfig()
	})
}

func (u *UI) Unmarshal(cfg *Config) {
	if err := u.V.Unmarshal(cfg); err != nil {
		u.Z.Fatal(Red("failed to unmarshal config"), zap.Error(err))
	}
}

func (u *UI) WriteConfig() {
	u.V.ReadInConfig()
	if err := u.V.WriteConfig(); err != nil {
		u.Z.Fatal(Red("faied to write config"), zap.Error(err))
	}
}
