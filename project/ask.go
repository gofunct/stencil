package project

import (
	"fmt"
	"github.com/gofunct/common/ui"
	"github.com/spf13/viper"
	"strings"
)

const (
	askTail         = " of this project"
	stringHead      = "[STRING] "
	stringSliceHead = "[STRINGSLICE] "
	mapHead         = "[MAP] "
	intHead         = "[INT] "
	otherHead       = "[OTHER] "
	floatHead       = "[FLOAT] "
	structHead      = "[STRUCT] "
	wit             = "what is the "
	wat             = "what are the "
)

var (
	projNameAsk     = stringQ("project name")
	configAsk       = stringQ("config file")
	toolsAsk        = stringSliceQ("third party dev tools")
	scopeAsk        = stringQ("scope")
	versionAsk      = stringQ("version")
	logLevelAsk     = stringQ("log level(debug or verbose)")
	servicesAsk     = stringSliceQ("grpc service names")
	dbUserAsk       = stringQ("database user name")
	dbNameAsk       = stringQ("datbase name")
	dbPasswordAsk   = stringQ("database password")
	dbHostAsk       = stringQ("datbase host")
	cloudRegionAsk  = stringQ("gcloud region")
	runvarAsk       = stringQ("runtime variable")
	runvarWaitAsk   = stringQ("runtime variable timeout")
	runvarConfigAsk = stringQ("runtime variable config name")
	dockerhubAsk    = stringQ("dockerhub user name")
	githubAsk       = stringQ("github user name")
	lisAsk          = stringQ("server listen port")
	imageAsk        = stringQ("docker base image")
	authorAsk       = stringQ("developer full name")
	emailAsk        = stringQ("developer email address")
	logoAsk         = stringQ("logo")
	bucketAsk       = stringQ("gcloud storage bucket")
	pluginsAsk      = stringQ("plugins")
)

func stringQ(s string) string {
	return fmt.Sprint(wit + s + askTail)
}

func stringSliceQ(s string) string {
	return fmt.Sprint(wat + s + askTail)
}

func (p *Project) CheckForEmptyVars() {
	if p.Notify == nil {
		p.Notify = ui.NewUI()
	}
	ask := p.Notify.Ask

	switch {
	case p.ConfigFile == "":
		p.ConfigFile = ask(configAsk)
		fallthrough
	case p.Image == "":
		p.Image = ask(imageAsk)
		fallthrough
	case p.Version == "":
		p.Version = ask(versionAsk)
		fallthrough
	case p.Runvar == "":
		p.Runvar = ask(runvarAsk)
		fallthrough
	case p.CloudRegion == "":
		p.CloudRegion = ask(cloudRegionAsk)
		fallthrough
	case p.RunvarConfig == "":
		p.RunvarConfig = ask(runvarConfigAsk)
		fallthrough
	case p.DbPassword == "":
		p.DbPassword = ask(dbPasswordAsk)
		fallthrough
	case p.DbName == "":
		p.DbName = ask(dbNameAsk)
		fallthrough
	case p.DbUser == "":
		p.DbUser = ask(dbUserAsk)
		fallthrough
	case p.DbHost == "":
		p.DbHost = ask(dbHostAsk)
		fallthrough
	case p.Bucket == "":
		p.Bucket = ask(bucketAsk)
		fallthrough
	case p.Lis == "":
		p.Lis = ask(lisAsk)
		fallthrough
	case p.Dockerhub == "":
		p.Dockerhub = ask(dockerhubAsk)
		fallthrough
	case p.Github == "":
		p.Github = ask(githubAsk)
		fallthrough
	case len(p.Services) == 0:
		p.Services = strings.Split(ask(bucketAsk), ",")
		fallthrough
	case p.Author == "":
		p.Author = ask(authorAsk)
		fallthrough
	case len(p.Tools) == 0:
		p.Tools = strings.Split(ask(toolsAsk), ",")
		fallthrough
	case p.Scope == "":
		p.Scope = ask(scopeAsk)
		fallthrough
	case p.Import == "":
		p.Import = ask(imageAsk)
		fallthrough
	case p.LogLevel == "":
		p.LogLevel = ask(logLevelAsk)
		fallthrough
	case p.Logo == "":
		p.Logo = ask(logoAsk)
		fallthrough
	default:
		notify.Success("all required config values are set!")
	}
}

func (p *Project) AskWriteConfig(v *viper.Viper, file string) error {
	answer := notify.Ask("would you like to write an updated config file?")
	switch {
	case strings.Contains(answer, "y"):
		return v.WriteConfigAs(file)
	}
	return nil
}

func (p *Project) Ask() {
	if p.Notify == nil {
		p.Notify = ui.NewUI()
	}
	ask := p.Notify.Ask

	p.Image = ask(imageAsk)
	p.Version = ask(versionAsk)
	p.CloudRegion = ask(cloudRegionAsk)
	p.RunvarConfig = ask(runvarConfigAsk)
	p.DbPassword = ask(dbPasswordAsk)
	p.DbName = ask(dbNameAsk)
	p.DbUser = ask(dbUserAsk)
	p.DbHost = ask(dbHostAsk)
	p.Bucket = ask(bucketAsk)
	p.Lis = ask(lisAsk)
	p.Dockerhub = ask(dockerhubAsk)
	p.Github = ask(githubAsk)
	p.Services = strings.Split(ask(bucketAsk), ",")
	p.Author = ask(authorAsk)
	p.Tools = strings.Split(ask(toolsAsk), ",")
	p.Scope = ask(scopeAsk)
	p.Import = ask(imageAsk)
	p.LogLevel = ask(logLevelAsk)
	p.Logo = ask(logoAsk)
	p.Notify.Success("all required config values are set!")
}
