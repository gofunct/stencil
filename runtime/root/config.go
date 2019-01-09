package root

import (
	"github.com/gofunct/stencil/runtime/ui"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/tcnksm/go-input"
	"go.uber.org/zap"
	"strings"
)

type configService struct {
	v		*viper.Viper
	q 		*input.UI
	Meta map[string]interface{}
}

func (c *configService) SetFs(a afero.Fs) {
	c.v.SetFs(a)

}

func (c *configService) Unmarshal(i interface{}) error {
	if err := c.v.Unmarshal(i); err != nil {
		return errors.Wrapf(err, "failed to unmarshal object")
	}

	return nil
}

func (c *configService) Bytes(i interface{}) ([]byte, error) {
	b, ok  := i.([]byte)
	if ok {
		return b, nil
	}
	return nil, errors.New("type does not contain any bytes")

}

func (c *configService) GetEnv() ([]string, error) {
	s := c.v.GetStringSlice("ENV")
	if len(s)== 0 {
		return nil, errors.New("configurator was queried for ENV, but no values were returned")
	}
	return s, nil
}

func (c *configService) Query() error {
	k, err := c.q.Ask(ui.Blue("what key would you like to set?"),  &input.Options{
		HideOrder: true,
		Loop:      true,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to ask for input")
	}
	v, err := c.q.Ask(ui.Blue(k+"<-key what is the value?"),  &input.Options{
		HideOrder: true,
		Loop:      true,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to ask for input")
	}
	b, err := c.q.Ask(ui.Blue("would you like to write this to config?"),  &input.Options{
		HideOrder: true,
		Loop:      true,
	})
	if err != nil {
		return errors.Wrapf(err, "Failed to ask for input")
	}
	c.v.Set(k, v)
	c.v.SetDefault(k, v)
	if strings.Contains(b, "y") ||strings.Contains(b, "Y")||strings.Contains(b, "Yes")||strings.Contains(b, "yes")  {
		if err := c.v.WriteConfig(); err != nil {
			return err
		}
	}
	zap.L().Debug(ui.Green("successfully updated config"))
	return nil
}


func (c *configService) Value(key string) interface{} {
	return c.v.Get(key)
}

func (c *configService) MergeMeta(v map[string]interface{}) {
	c.v.MergeConfigMap(v)
}

func (c *configService) GetMeta() map[string]interface{} {
	return c.v.AllSettings()
}

func (c *configService) Debug() {
	viper.Debug()
}



