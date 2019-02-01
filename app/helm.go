package app

import (
	"fmt"
	"github.com/spf13/pflag"
)

var ChartKeys = []string{"icon", "apiVersion", "version", "appVersion", "name", "description", "keywords", "home", "sources", "maintainers"}

type ConfigFile int
type ConfigMode int

const (
	ChartConfig  ConfigFile = 1
	ValuesConfig ConfigFile = 2
	TfConfig     ConfigFile = 3
	KubeConfig   ConfigFile = 4
)
const (
	Query ConfigMode = 1
	Flag  ConfigMode = 2
	File  ConfigMode = 4
	Env   ConfigMode = 3
)

type Parameter struct {
	*pflag.Flag
}

type ParamMap map[ConfigMode]Parameter
type QueryMap map[ConfigFile]ParamMap

func (q QueryMap) test() {
	for k, v := range q {
		fmt.Printf("%s%T%v", k, k, k)
		fmt.Printf("%s%T%v", v, v, v)
		for j, b := range v {
			fmt.Printf("%s%T%v", j, j, j)
			fmt.Printf("%s%T%v", b, b, b)
		}
	}
}

type Chartemp struct {
	Icon        Parameter     `json:"icon"`
	ApiVersion  Parameter     `json:"apiVersion"`
	Version     Parameter     `json:"version"`
	AppVersion  Parameter     `json:"appVersion"`
	Name        Parameter     `json:"name"`
	Description Parameter     `json:"description"`
	KeyWords    []string      `json:"keywords"`
	Home        string        `json:"home"`
	Sources     []string      `json:"sources"`
	Maintainers []*Maintainer `json:"maintainers"`
	*pflag.FlagSet
}
