package root

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gofunct/iio"
	"github.com/gofunct/stencil/runtime/ui"
	"github.com/mattn/go-colorable"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/tcnksm/go-input"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"runtime"
)

func NewContext(fs afero.Fs) Context {
	i := iio.DefaultIO()
	v := viper.New()
	v.SetFs(fs)
	c :=  &con{
		v:      viper.New(),
		q:      &input.UI{
			Writer: i.Out,
			Reader: i.In,
		},
		Meta:   v.AllSettings(),
	}

	c.v.AddConfigPath(".")
	c.v.AutomaticEnv()
	c.v.SetConfigName("stencil")
	c.v.SetConfigFile("stencil.yaml")
	c.v.BindEnv("GO111MODULE")
	c.v.BindEnv("PWD")
	c.v.BindEnv("HOME")
	c.v.BindEnv("GOROOT")
	c.v.BindEnv("GOPATH")
	c.v.BindEnv("USER")
	c.v.BindEnv("SHELL")
	src := filepath.Join(c.v.GetString("GOPATH"), "src")
	c.v.Set("gosrc", src)
	bin := filepath.Join(c.v.GetString("GOPATH"), "bin")
	c.v.Set("gobin", bin)
	imp, err := filepath.Rel(c.v.GetString("gosrc"), c.v.GetString("pwd"))
	if err != nil {
		zap.L().Fatal("failed to locate working directory")
	}
	c.v.Set("import", imp)
	c.v.Set("download", "go get "+imp)
	c.v.WriteConfig()
	c.v.WatchConfig()
	c.v.OnConfigChange(func(e fsnotify.Event) {
		zap.L().Debug(ui.Blue("Config file changed:"), zap.String("event", e.Name), zap.String("stringer", e.String()))
		c.v.WriteConfig()
	})
	return c
}

// Stdio returns a standard IO object.
func DefaultIO() ioService {
	io := &ioContainer{
		InR:  os.Stdin,
		OutW: os.Stdout,
		ErrW: os.Stderr,
	}
	if runtime.GOOS == "windows" {
		io.OutW = colorable.NewColorableStdout()
		io.ErrW = colorable.NewColorableStderr()
	}
	return io
}
