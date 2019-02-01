package app

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gofunct/gofs"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/tcnksm/go-input"
	"go.opencensus.io/trace"
	"gocloud.dev/blob"
	"gocloud.dev/health"
	"gocloud.dev/health/sqlhealth"
	"gocloud.dev/runtimevar"
	"gocloud.dev/server"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const apiVersion = "v1"

var applicationSet = wire.NewSet(
	newApplication,
	appHealthChecks,
	trace.AlwaysSample,
)

type AppFunc func(*App) error

type tfItem struct {
	Sensitive bool
	Type      string
	Value     string
}

type Maintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Url   string `json:"url"`
}

type State struct {
	Project          tfItem `json:"project"`
	ClusterName      tfItem `json:"cluster_name"`
	ClusterZone      tfItem `json:"cluster_zone"`
	DatabaseInstance tfItem `json:"database_instance"`
	DatabaseRegion   tfItem `json:"database_region"`
	RUnVarConfig     tfItem `json:"run_var_config"`
	RunVarName       tfItem `json:"run_var_name"`
}
type Path struct {
	TfDir       string `json:"tfDir"`
	ProtoDir    string `json:"protoDir"`
	TemplateDir string `json:"templateDir"`
	StaticDir   string `json:"staticDir"`
	CertPath    string `json:"certPath"`
	KeyPath     string `json:"keyPath"`
}

type Chart struct {
	Icon        string        `json:"icon"`
	ApiVersion  string        `json:"apiVersion"`
	Version     string        `json:"version"`
	AppVersion  string        `json:"appVersion"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	KeyWords    []string      `json:"keywords"`
	Home        string        `json:"home"`
	Sources     []string      `json:"sources"`
	Maintainers []*Maintainer `json:"maintainers"`
}
type Values struct {
	Modules      bool     `json:"modules"`
	Bucket       string   `json:"bucket"`
	DbName       string   `json:"dbHost"`
	DbHost       string   `json:"dbHost"`
	DbUser       string   `json:"dbUser"`
	DbPassword   string   `json:"dbPassword"`
	RunVarWait   string   `json:"runVarWait"`
	GitIgnore    []string `json:"gitignore"`
	DockerIgnore []string `json:"dockerignore"`
	RunVar       string   `json:"run_var"`
	State        *State   `json:"state"`
	Path         *Path    `json:"path"`
}

type App struct {
	Values   *Values `json:"values"`
	Chart    *Chart  `json:"chart"`
	AppFuncs []AppFunc
	fs       afero.Fs
	chart    *viper.Viper
	values   *viper.Viper
	q        *input.UI
	buf      *bytes.Buffer
	srv      *server.Server
	db       *sql.DB
	bucket   *blob.Bucket

	// The following fields are protected by mu:
	mu   sync.RWMutex
	motd string
}

func (a *App) Execute(reader io.Reader, writer io.Writer) error {
	a.buf = bytes.NewBuffer(nil)
	if reader == nil {
		reader = os.Stdin
	}
	if writer == nil {
		writer = os.Stdout
	}
	if _, err := a.buf.ReadFrom(reader); err != nil {
		return err
	}
	{
		a.chart = viper.New()
		a.chart.SetConfigType("yaml")
		a.chart.SetConfigName("Chart")
		a.chart.AddConfigPath("helm")
		a.chart.AutomaticEnv()
		a.chart.AllowEmptyEnv(true)
		{
			a.chart.SetDefault("icon", "https://github.com/gofunct/logo/blob/master/white_logo_dark_background.jpg?raw=true")
			a.chart.SetDefault("apiVersion", apiVersion)
			a.chart.SetDefault("version", "v0.0.1")
			a.chart.SetDefault("appVersion", "0.1")
			a.chart.SetDefault("name", filepath.Base(os.Getenv("HOME")))
			a.chart.SetDefault("description", "this is your applications default description")
			a.chart.SetDefault("keywords", []string{"stencil", filepath.Base(os.Getenv("HOME"))})
			a.chart.SetDefault("home", "https://"+filepath.Join("github", os.Getenv("PWD"))+".com")
			a.chart.SetDefault("sources", filepath.Join("github", os.Getenv("PWD")))
			a.chart.SetDefault("maintainers", &Maintainer{
				Name:  "Coleman Word",
				Email: "coleman.word@gofunct.com",
				Url:   "https://github.com/gofunct",
			})
		}
		_ := a.values.SafeWriteConfig()
		if err := a.chart.ReadInConfig(); err != nil {
			return errors.Wrap(err, "failed to read Chart config")
		}
	}
	{
		a.values = viper.New()
		a.values.SetConfigType("yaml")
		a.values.SetConfigName("values")
		a.values.AddConfigPath("helm")
		a.values.AutomaticEnv()
		a.values.AllowEmptyEnv(true)
		{
			if gofs.FileExists("deploy/*.tf") {
				bytess, err := a.ShellOutput("terraform", "output", "-state", "deploy", "-json")
				if err != nil {
					return errors.Wrap(err, `failed to execute terraform with args: "terraform", "output", "-state", "deploy", "-json"`)
				}
				if err := a.values.ReadConfig(bytes.NewBuffer(bytess)); err != nil {
					return errors.Wrap(err, "failed to read in terraform config")
				}
			}
		}
		{
			a.values.SetDefault("modules", true)
			a.values.SetDefault("bucket", filepath.Base(os.Getenv("HOME"))+"-bucket")
			a.values.SetDefault("dbHost", filepath.Base(os.Getenv("HOME"))+"-db")
			a.values.SetDefault("dbHost", "")
			a.values.SetDefault("dbUser", os.Getenv("USER"))
			a.values.SetDefault("dbPassword", "admin")
			a.values.SetDefault("runVarWait", 15*time.Second)
			a.values.SetDefault("gitignore", []string{"vendor", ".idea", "bin", "temp", "certs"})
			a.values.SetDefault("dockerignore", []string{"*.md", ".idea"})
			a.values.SetDefault("tfDir", "./deploy")
			a.values.SetDefault("protoDir", "./proto")
			a.values.SetDefault("templateDir", "./templates")
			a.values.SetDefault("staticDir", "./static")
			a.values.SetDefault("certPath", "./certs/app.pem")
			a.values.SetDefault("certPath", "./certs/app.key")
		}
		_ = a.values.SafeWriteConfig()
		if err := a.values.ReadInConfig(); err != nil {
			return errors.Wrap(err, "failed to read config")
		}

	}
	{
		if err := a.chart.Unmarshal(&a.Chart); err != nil {
			return errors.Wrap(err, "failed to unmarshal Chart config")
		}
		if err := a.values.Unmarshal(&a.Chart); err != nil {
			return errors.Wrap(err, "failed to unmarshal config")
		}

		for _, f := range a.AppFuncs {
			if err := f(a); err != nil {
				return err
			}
		}

		if err := a.values.WriteConfig(); err != nil {
			return errors.Wrap(err, "failed to update values config")
		}
		if err := a.chart.WriteConfig(); err != nil {
			return errors.Wrap(err, "failed to update chart config")
		}
	}

	if _, err := a.buf.WriteTo(writer); err != nil {
		return err
	}
	return nil
}

func (a *App) WriteConfig() error {
	if err := a.chart.WriteConfig(); err != nil {
		return err
	}
	if err := a.values.WriteConfig(); err != nil {
		return err
	}
	return nil
}

func (a *App) Shell(args ...string) (stdout string, err error) {
	stdoutb, err := a.ShellOutput(args...)
	return strings.TrimSpace(string(stdoutb)), err
}

func (a *App) ShellOutput(args ...string) (stdout []byte, err error) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Env = append(cmd.Env, os.Environ()...)
	stdoutb, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("running %v: %v", cmd.Args, err)
	}
	return stdoutb, nil
}

func (a *App) MarshalJson(v interface{}) []byte {
	output, _ := json.MarshalIndent(v, "", "  ")
	return output
}

func (a *App) UnmarshalJson(v interface{}) error {
	s := a.MarshalJson(a.values.AllSettings())
	s = append(s, a.MarshalJson(a.chart.AllSettings())...)
	if err := json.Unmarshal(s, v); err != nil {
		return err
	}
	return nil
}

func (a *App) Read(p []byte) (n int, err error) {
	return a.buf.Read(p)
}

func (a *App) Write(p []byte) (n int, err error) {
	return a.buf.Write(p)
}

func (a *App) QueryChart(name, defaul, question string, loop, required bool) {
	if a.q == nil {
		a.q = input.DefaultUI()
	}
	ans, err := a.q.Ask(question, &input.Options{
		Default:      defaul,
		Loop:         loop,
		Required:     required,
		ValidateFunc: validateChart(name, required),
	})
	if err != nil {
		panic(err)
	}
}

func (a *App) QueryValue(name, defaul, question string, required bool) {
	if a.q == nil {
		a.q = input.DefaultUI()
	}

	ans, err := a.q.Ask(question, &input.Options{
		Default:      defaul,
		Loop:         required,
		Required:     required,
		ValidateFunc: validateValues(name, required),
	})
	if err != nil {
		panic(err)
	}
}

func Initialize(app *App) *App {
	app.buf = bytes.NewBuffer(nil)
	app.buf.ReadFrom(os.Stdin)
	app.buf.ReadFrom(os.Stdout)
	app.chart = viper.New()
	app.chart.SetConfigType("yaml")
	app.chart.SetConfigName("Chart")
	app.chart.AddConfigPath("helm")
	app.chart.AutomaticEnv()
	app.chart.AllowEmptyEnv(true)
	app.values = viper.New()
	app.values.SetConfigType("yaml")
	app.values.SetConfigName("values")
	app.values.AddConfigPath("helm")
	app.values.AutomaticEnv()
	app.values.AllowEmptyEnv(true)
	app.q = input.DefaultUI()
	return app
}

func validateChart(key string, required bool) func(s string) error {
	return func(s string) error {
		if s == "" && required {
			return errors.New("query error: empty input detected")
		}
		if len(s) > 50 {
			return errors.New("query error: over 50 characters")
		}

		return nil
	}
}
func validateValues(key string, required bool) func(s string) error {
	return func(s string) error {
		if s == "" {
			return errors.New("query error: empty input detected")
		}
		if len(s) > 50 {
			return errors.New("query error: over 50 characters")
		}
		return nil
	}
}

// appHealthChecks returns a health check for the database. This will signal
// to Kubernetes or other orchestrators that the server should not receive
// traffic until the server is able to connect to its database.
func appHealthChecks(db *sql.DB) ([]health.Checker, func()) {
	dbCheck := sqlhealth.New(db)
	list := []health.Checker{dbCheck}
	return list, func() {
		dbCheck.Stop()
	}
}

// newApplication creates a new application struct based on the backends and the message
// of the day variable.
func newApplication(srv *server.Server, db *sql.DB, bucket *blob.Bucket, motdVar *runtimevar.Variable) *App {
	app := &App{
		srv:    srv,
		db:     db,
		bucket: bucket,
	}
	go app.watchMOTDVar(motdVar)
	return app
}

// watchMOTDVar listens for changes in v and updates the app's message of the
// day. It is run in a separate goroutine.
func (a *App) watchMOTDVar(v *runtimevar.Variable) {
	ctx := context.Background()
	for {
		snap, err := v.Watch(ctx)
		if err != nil {
			log.Printf("watch MOTD variable: %v", err)
			continue
		}
		log.Println("updated MOTD to", snap.Value)
		a.mu.Lock()
		a.motd = snap.Value.(string)
		a.mu.Unlock()
	}
}
