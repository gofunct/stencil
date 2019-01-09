package runtime

import (
	"fmt"
	"github.com/gofunct/iio"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/tcnksm/go-input"
	"go.uber.org/zap"
)

type UI struct {
	Config *Config
	IO     *iio.IO
	Q      *input.UI
	V      *viper.Viper
	F      []*pflag.Flag
	Z      *zap.Logger
}

func NewUI() *UI {
	var c Config
	dio := iio.DefaultIO()
	u := new(UI)
	u.Config = &c
	u.SetupLogger()
	u.IO = dio
	u.Q = &input.UI{
		Writer: u.IO.Out,
		Reader: u.IO.In,
	}
	u.InitConfig()
	u.FlagsFromSettings()
	return u
}

func (u *UI) SetupLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
	u.Z = logger
}

func (u *UI) FlagsFromSettings() {
	for k, v := range u.V.AllSettings() {
		f := &pflag.Flag{
			Name:     k,
			DefValue: v.(string),
		}
		u.V.BindPFlag(k, f)
		u.F = append(u.F, f)
	}
}
func (u *UI) BindCobra(cmd *cobra.Command) {
	cmd.Flags().SetOutput(u.IO.Out)
	for _, f := range u.F {
		cmd.Flags().String(f.Name, f.DefValue, "")
	}
	merge := func() map[string]string {
		a := make(map[string]string)
		for k, v := range u.V.AllSettings() {
			a[k] = v.(string)
		}
		return a
	}
	cmd.Annotations = merge()
	cmd.Version = u.V.GetString("version")
	u.AddLoggingFlags(cmd)
}

func (u *UI) Debug() {
	u.Z.Debug("\n ui.Config", zap.Any("config", u.Config))
	u.Z.Debug("\n viper.AllSettings()", zap.Any("settings", u.V.AllSettings()))
	fmt.Println("\n viper.Debug()")
	u.V.Debug()
}
