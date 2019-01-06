package app

import (
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"os"
	"os/exec"
	"os/user"
	"runtime"
)

func username() string {
	user, _ := user.Current()
	return user.Username
}

// returns user home directory
func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

// Generates App Path
func createAppPath(VCSHost string, user string, appname string) string {
	gopath := os.Getenv("GOPATH")
	if len(gopath) < 1 {
		gopath = userHomeDir() + "/go"
	}
	apppath := gopath + "/src/" + VCSHost + "/" + user + "/" + appname
	os.MkdirAll(apppath, 0755)
	return apppath
}

// runs go format on all generated go files
func runGoFormat(VCSHost string, user string, app string) error {
	gopath := VCSHost + "/" + user + "/" + app
	_, err := exec.Command("go", "fmt", gopath).Output()
	return errors.Wrapf(err, "%s", "failed to execute go format")
}

func mkDirAll(fs afero.Fs, path string) {
	fs.MkdirAll(path, 0755)
	os.MkdirAll(path, 0755)
}
