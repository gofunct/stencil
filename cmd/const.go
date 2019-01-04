package cmd

import (
	"fmt"
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
	configAsk       = stringQ("config file")
	exampleAsk      = stringQ("example usage")
	aliasesAsk      = stringSliceQ("aliases")
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
)

func stringQ(s string) string {
	return fmt.Sprint(wit + s + askTail)
}

func stringSliceQ(s string) string {
	return fmt.Sprint(wat + s + askTail)
}

func initConfig() {
	ask := notify.Ask
	switch {
	case cfgFile == "":
		cfgFile = ask(configAsk)
		fallthrough
	case options.Example == "":
		options.Example = ask(exampleAsk)
		fallthrough
	case len(options.Aliases) == 0:
		options.Aliases = strings.Split(ask(aliasesAsk), ",")
		fallthrough
	case options.Image == "":
		options.Image = ask(imageAsk)
		fallthrough
	case options.Version == "":
		options.Version = ask(versionAsk)
		fallthrough
	case options.Runvar == "":
		cfgFile = ask(runvarAsk)
		fallthrough
	case options.CloudRegion == "":
		options.CloudRegion = ask(cloudRegionAsk)
		fallthrough
	case options.RunvarConfig == "":
		options.RunvarConfig = ask(runvarConfigAsk)
		fallthrough
	case options.DbPassword == "":
		options.DbPassword = ask(dbPasswordAsk)
		fallthrough
	case options.DbName == "":
		options.DbName = ask(dbNameAsk)
		fallthrough
	case options.DbUser == "":
		options.DbUser = ask(dbUserAsk)
		fallthrough
	case options.DbHost == "":
		options.DbHost = ask(dbHostAsk)
		fallthrough
	case options.Bucket == "":
		options.Bucket = ask(bucketAsk)
		fallthrough
	case options.Lis == "":
		options.Lis = ask(lisAsk)
		fallthrough
	case options.Dockerhub == "":
		options.Dockerhub = ask(dockerhubAsk)
		fallthrough
	case options.Github == "":
		options.Github = ask(githubAsk)
		fallthrough
	case len(options.Services) == 0:
		options.Services = strings.Split(ask(bucketAsk), ",")
		fallthrough
	case options.Author == "":
		options.Author = ask(authorAsk)
		fallthrough
	case len(options.Tools) == 0:
		options.Tools = strings.Split(ask(toolsAsk), ",")
		fallthrough
	case options.Scope == "":
		options.Scope = ask(scopeAsk)
		fallthrough
	case options.Import == "":
		options.Import = ask(imageAsk)
		fallthrough
	case options.LogLevel == "":
		options.LogLevel = ask(logLevelAsk)
		fallthrough
	case options.Logo == "":
		options.Logo = ask(logoAsk)
		fallthrough
	default:
		notify.Success("all required config values are set!")
	}
}

func askWriteConfig() error {
	answer := notify.Ask("would you like to write an updated config file?")
	switch {
	case strings.Contains(answer, "y"):
		return config.WriteConfigAs("config.yaml")
	}
	return nil
}

func askme() {
	ask := notify.Ask
		options.Example = ask(exampleAsk)
		options.Aliases = strings.Split(ask(aliasesAsk), ",")
		options.Image = ask(imageAsk)
		options.Version = ask(versionAsk)
		options.CloudRegion = ask(cloudRegionAsk)
		options.RunvarConfig = ask(runvarConfigAsk)
		options.DbPassword = ask(dbPasswordAsk)
		options.DbName = ask(dbNameAsk)
		options.DbUser = ask(dbUserAsk)
		options.DbHost = ask(dbHostAsk)
		options.Bucket = ask(bucketAsk)
		options.Lis = ask(lisAsk)
		options.Dockerhub = ask(dockerhubAsk)
		options.Github = ask(githubAsk)
		options.Services = strings.Split(ask(bucketAsk), ",")
		options.Author = ask(authorAsk)
		options.Tools = strings.Split(ask(toolsAsk), ",")
		options.Scope = ask(scopeAsk)
		options.Import = ask(imageAsk)
		options.LogLevel = ask(logLevelAsk)
		options.Logo = ask(logoAsk)
		notify.Success("all required config values are set!")
}