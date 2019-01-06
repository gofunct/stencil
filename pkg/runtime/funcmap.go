package runtime

import (
	"bytes"
	"github.com/Masterminds/sprig"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"text/template"
	"time"
)

type defaultConfig struct {
	ConfigFile   string        `mapstructure:"config-file"`
	AbsolutePath string        `mapstructure:"absolute-path"`
	CommandPath  string        `mapstructure:"command-path"`
	SourcePath   string        `mapstructure:"source-path"`
	ProjName     string        `mapstructure:"proj-name"`
	Revision     string        `mapstructure:"revision"`
	Branch       string        `mapstructure:"branch"`
	HEAD         bool          `mapstructure:"head"`
	CopyRight    string        `mapstructure:"copyright"`
	Package      string        `mapstructure:"package"`
	Logo         string        `mapstructure:"logo"`
	Import       string        `mapstructure:"import"`
	Author       string        `mapstructure:"author"`
	Email        string        `mapstructure:"email"`
	Scope        string        `mapstructure:"scope"`
	Tools        []string      `mapstructure:"tools"`
	Version      string        `mapstructure:"version"`
	LogLevel     string        `mapstructure:"log-level"`
	Services     []string      `mapstructure:"services"`
	Github       string        `mapstructure:"github"`
	Dockerhub    string        `mapstructure:"dockerhub"`
	Lis          string        `mapstructure:"lis"`
	Bucket       string        `mapstructure:"bucket"`
	DbHost       string        `mapstructure:"db-host"`
	DbPassword   string        `mapstructure:"db-pw"`
	DbUser       string        `mapstructure:"db-user"`
	DbName       string        `mapstructure:"db-name"`
	Runvar       string        `mapstructure:"runvar"`
	RunVarWait   time.Duration `mapstructure:"runvar-wait"`
	CloudRegion  string        `mapstructure:"cloud-region"`
	RunvarConfig string        `mapstructure:"runvar-config"`
	Image        string        `mapstructure:"image"`
}

var stencilTag = func(config *mapstructure.DecoderConfig) {
	config.TagName = "stencil"
	config.ErrorUnused = false
	zap.L().Warn("unused keys in struct detected", zap.Strings("unused-keys", config.Metadata.Unused))
}

func (r *Runtime) mapFromDefaultConfig(root string) error {
	if r.Config == nil {
		return errors.New("must initialize the runtime viper config before using funcmap")
	}
	var def *defaultConfig
	r.Config.Unmarshal(def, stencilTag)

	pkgName, err := r.GetPackageName(root)
	if err != nil {
		return errors.Wrap(err, "failed to decide a package name")
	}
	importPath, err := r.GetImportPath(root)
	if err != nil {
		return errors.WithStack(err)
	}
	sources, err := r.FindMainPackagesAndSources(root)
	if err != nil {
		return errors.WithStack(err)
	}
	r.Config.Set("sources", sources)
	r.Config.Set("pkgName", pkgName)
	r.Config.Set("importPath", importPath)

	r.FMap = r.configToFuncMap(sprig.GenericFuncMap())

	return nil
}

func (r *Runtime) configToFuncMap(new map[string]interface{}) map[string]interface{} {
	for k, v := range r.Config.AllSettings() {
		new[k] = v
	}
	return new
}

// TemplateString is a compilable string with text/template package
type TemplateString string

// Compile generates textual output applied a parsed template to the specified values
func (s TemplateString) Compile(v interface{}) (string, error) {
	tmpl, err := template.New("").Parse(string(s))
	if err != nil {
		return string(s), errors.Wrapf(err, "failed to parse a template: %q", string(s))
	}
	buf := &bytes.Buffer{}
	err = tmpl.Execute(buf, v)
	if err != nil {
		return string(s), errors.Wrapf(err, "failed to execute a template: %q", string(s))
	}
	return string(buf.Bytes()), nil
}
