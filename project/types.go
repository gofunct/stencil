package project

import (
	"github.com/gofunct/common/ui"
	"github.com/gofunct/stencil/project/plugins"
	"github.com/izumin5210/gex"
	"github.com/jessevdk/go-assets"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go/build"
	"os/user"
	"strings"
	"time"
)

var (
	notify = ui.NewUI()
	ask    bool
	write  bool
)

// Project contains name, license and paths to projects.
type Project struct {
	Os           afero.Fs
	Assets       *assets.FileSystem
	Plugin       string `mapstructure:"plugin"`
	Notify       ui.UI
	GexCfg       *gex.Config
	ConfigFile   string        `mapstructure:"config-file"`
	AbsolutePath string        `mapstructure:"absolute-path"`
	CommandPath  string        `mapstructure:"command-path"`
	SourcePath   string        `mapstructure:"source-path"`
	LicenseType  License       `mapstructure:",squash"`
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

func (p *Project) BindToRoot(rootCmd *cobra.Command, v *viper.Viper) {
	pFlags := rootCmd.PersistentFlags()
	v.SetConfigFile("stencil.yaml")
	v.AddConfigPath(".")
	v.SetEnvPrefix("stencil")
	v.AutomaticEnv()

	{
		pFlags.StringVar(&p.ConfigFile, "config", "", configAsk)
		pFlags.BoolVarP(&ask, "ask", "a", false, "fill out a survey to set configs")
		pFlags.BoolVarP(&write, "write", "w", false, "fill out a survey to set configs")
		pFlags.StringVar(&p.Logo, "logo", "https://github.com/gofunct/common/blob/master/logo/dark_logo_transparent_background.png?raw=true", logoAsk)
		pFlags.StringVar(&p.Author, "author", "Coleman Word", authorAsk)
		pFlags.StringVar(&p.Email, "email", "coleman.word@gofunct.com", emailAsk)
		pFlags.StringVar(&p.Scope, "scope", "not available", scopeAsk)
		pFlags.StringVar(&p.Version, "version", "v0.1.1", versionAsk)
		pFlags.StringVar(&p.LogLevel, "log-level", "verbose", logLevelAsk)
		pFlags.StringSliceVar(&p.Tools, "tools", []string{"gex", "reviewdog", "errcheck", "gogoslick", "grpc-swagger", "grpc-gateway", "protocgen-go", "golint", "megacheck", "wire", "unparam", "gox", "gotests"}, toolsAsk)
		pFlags.StringSliceVar(&p.Services, "services", []string{""}, servicesAsk)
		pFlags.StringVar(&p.Github, "github", "gofunct", githubAsk)
		pFlags.StringVar(&p.Dockerhub, "dockerhub", "", dockerhubAsk)
		pFlags.StringVar(&p.Lis, "lis", ":8080", lisAsk)
		pFlags.StringVar(&p.Bucket, "bucket", "gofunct-storage", bucketAsk)
		pFlags.StringVar(&p.DbHost, "db-host", "", dbHostAsk)
		pFlags.StringVar(&p.DbPassword, "db-pw", "admin", dbPasswordAsk)
		pFlags.StringVar(&p.DbPassword, "db-user", "gofunct", dbUserAsk)
		pFlags.StringVar(&p.DbPassword, "db-name", "gofunct-db", dbNameAsk)
		pFlags.StringVar(&p.Runvar, "runvar", "", runvarAsk)
		pFlags.DurationVar(&p.RunVarWait, "runvar-wait", 5*time.Second, runvarWaitAsk)
		pFlags.StringVar(&p.CloudRegion, "cloud-region", "", cloudRegionAsk)
		pFlags.StringVar(&p.RunvarConfig, "runvar-config", "", runvarConfigAsk)
		pFlags.StringVar(&p.Image, "image", "alpine:3.8", imageAsk)
		pFlags.StringVar(&p.Plugin, "plugins", "init", configAsk)
	}

	v.BindPFlags(pFlags)
	v.Unmarshal(&p)
	v.Set("copyright", p.CopyrightLine(v))
	v.Set("license", p.GetLicense().Header)
	v.Set("cmdPackage", p.CopyrightLine(v))
	v.Set("cmdName", p.CopyrightLine(v))

	switch p.Plugin {
	case "init":
		p.Assets = plugins.Init
	}

	if p.Assets == nil {
		a := p.Notify.Ask("No assets have been registered, would you like to use the defualt inititializers?")
		if strings.Contains(a, "y") {
			p.Assets = plugins.Init
		} else {
			p.Exit("failed to register assets in filesytstem")
		}
	}
}

// Licenses contains all possible licenses a user can choose from.
var Licenses = make(map[string]License)

// License represents a software license agreement, containing the Name of
// the license, its possible matches (on the command line as given to cobra),
// the header to be used with each file on the file's creating, and the text
// of the license
type License struct {
	Name            string   // The type of license in use
	PossibleMatches []string // Similar names to guess
	Text            string   // License text data
	Header          string   // License header for source files
}

//All constants, vars and types

const (
	// PackageSeparator is a package separator string on protobuf.
	PackageSeparator = "."
)

// Make visible for testing
var (
	BuildContext build.Context
	GetOSUser    getOSUserFunc
)

type getOSUserFunc func() (*user.User, error)

// RootDir represents a project root directory.
type RootDir struct {
	Path
}

// Path represents a filepath in filesystem.
type Path string
