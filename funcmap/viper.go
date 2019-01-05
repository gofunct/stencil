package funcmap

import (
	"github.com/spf13/viper"
	"time"
)

func CopyrightFromConfig() string {
	author := viper.GetString("author")
	year := time.Now().Format("2006")
	return "Copyright © " + year + " " + author
}
