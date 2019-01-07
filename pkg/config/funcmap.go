package config

import (
	"github.com/Masterminds/sprig"
	"github.com/gofunct/stencil/pkg/fs"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"html/template"
	"path"
	"time"
)

type Config struct {
	FMap        *template.FuncMap
	TemplString *TemplateString
	V           *viper.Viper
	P           *fs.Project
	Meta        map[string]interface{}
}

type defaultConfig struct {
	ConfigFile   string        `stencil:"config-file"`
	AbsolutePath string        `stencil:"absolute-path"`
	CommandPath  string        `stencil:"command-path"`
	SourcePath   string        `stencil:"source-path"`
	ProjName     string        `stencil:"proj-name"`
	Revision     string        `stencil:"revision"`
	Branch       string        `stencil:"branch"`
	HEAD         bool          `stencil:"head"`
	CopyRight    string        `stencil:"copyright"`
	Package      string        `stencil:"package"`
	Logo         string        `stencil:"logo"`
	Import       string        `stencil:"import"`
	Author       string        `stencil:"author"`
	Email        string        `stencil:"email"`
	Scope        string        `stencil:"scope"`
	Tools        []string      `stencil:"tools"`
	Version      string        `stencil:"version"`
	LogLevel     string        `stencil:"log-level"`
	Services     []string      `stencil:"services"`
	Github       string        `stencil:"github"`
	Dockerhub    string        `stencil:"dockerhub"`
	Lis          string        `stencil:"lis"`
	Bucket       string        `stencil:"bucket"`
	DbHost       string        `stencil:"db-host"`
	DbPassword   string        `stencil:"db-pw"`
	DbUser       string        `stencil:"db-user"`
	DbName       string        `stencil:"db-name"`
	Runvar       string        `stencil:"runvar"`
	RunVarWait   time.Duration `stencil:"runvar-wait"`
	CloudRegion  string        `stencil:"cloud-region"`
	RunvarConfig string        `stencil:"runvar-config"`
	Image        string        `stencil:"image"`
}

var stencilTag = func(config *mapstructure.DecoderConfig) {
	config.TagName = "stencil"
	config.ErrorUnused = false
	zap.L().Warn("unused keys in struct detected", zap.Strings("unused-keys", config.Metadata.Unused))
}

func (c *Config) MergeWithProject(p *fs.Project) {
	if c.V == nil {
		c.V = viper.New()
	}
	c.P = p
	var def *defaultConfig
	c.V.Unmarshal(def, stencilTag)

	appNameBase := path.Base(c.P.Name())
	absPathBase := path.Base(c.P.AbsPath())
	srcPathBase := path.Base(c.P.SrcPath())
	cmdPathBase := path.Base(c.P.CmdPath())

	appName := c.P.Name()
	absPath := c.P.AbsPath()
	srcPath := c.P.SrcPath()
	cmdPath := c.P.CmdPath()

	c.V.Set("appName", appName)
	c.V.Set("appNameBase", appNameBase)

	c.V.Set("absPath", absPath)
	c.V.Set("absPathBase", absPathBase)

	c.V.Set("srcPath", srcPath)
	c.V.Set("srcPathBase", srcPathBase)

	c.V.Set("cmdPath", cmdPath)
	c.V.Set("cmdPathBase", cmdPathBase)

	c.Meta = c.AddToSettings(sprig.GenericFuncMap())
}

func (c *Config) AddToSettings(m map[string]interface{}) map[string]interface{} {
	for k, v := range c.V.AllSettings() {
		m[k] = v
	}
	return m
}
