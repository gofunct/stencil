package runtime

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/tcnksm/go-input"
	"go.uber.org/zap"
)

func validateTF(q string) input.ValidateFunc {
	return func(ans string) error {
		zap.L().Debug("received response", zap.String("question", q), zap.String("answer", ans))
		if ans != "true" && ans != "false" {
			return fmt.Errorf("input must be true or false")
		}
		return nil
	}
}

func validateYN(q string) input.ValidateFunc {
	return func(ans string) error {
		zap.L().Debug("received response", zap.String("question", q), zap.String("answer", ans))
		if ans != "y" && ans != "n" {
			return fmt.Errorf("input must be y or n")
		}
		return nil
	}
}

func (u *UI) YesNoBool(q string) (bool, error) {
	ans, err := u.Q.Ask(fmt.Sprintf("%s [y/n]", Q+" "+q), &input.Options{
		HideOrder:    true,
		Loop:         true,
		ValidateFunc: validateYN(q),
	})
	if err != nil {
		return false, errors.WithStack(err)
	}
	return ans == "y", nil
}

func (u *UI) TrueFalseBool(q string) (bool, error) {
	ans, err := u.Q.Ask(fmt.Sprintf("%s [true/fase]", Q+" "+q), &input.Options{
		HideOrder:    true,
		Loop:         true,
		ValidateFunc: validateTF(q),
	})
	if err != nil {
		return false, errors.WithStack(err)
	}
	return ans == "true", nil
}

func (u *UI) String(q string) string {
	ans, err := u.Q.Ask(fmt.Sprintf("%s", Q+" "+q), &input.Options{
		HideOrder: true,
		Loop:      true,
	})
	if err != nil {
		u.Z.Fatal("Failed to ask for input", zap.Error(err))
	}
	return ans
}

func (u *UI) Email() {
	if u.Config.PROJECT.Email == "" {
		s := u.String(Q + " What is your email?")
		u.V.Set("email", s)
		u.Z.Debug("Result: ", zap.String("email", s))
	}
}

func (u *UI) Domain() {
	if u.Config.PROJECT.Domain == "" {
		s := u.String(Q + " What is your domain name address?")
		u.V.SetDefault("domain", s)
		u.Z.Debug("Result: ", zap.String("domain", s))
	}
}

func (u *UI) AskAppName() {
	if len(u.Config.PROJECT.Services) == 0 {
		s := u.String(Q + " What is the name of your new app")
		u.V.SetDefault("appName", s)
		u.Z.Debug("Result: ", zap.String("appName", s))
	}
}

func (u *UI) AskServices() {
	if len(u.Config.PROJECT.Services) == 0 {
		s := u.String(Q + " What is the name of your new service")
		u.V.SetDefault("serviceName", s)
		u.Z.Info("Result: ", zap.String("serviceName", s))
	}
}

func (u *UI) GithubUN() {
	if u.Config.PROJECT.Github == "" {
		s := u.String(Q + " What is your github user name?")
		u.V.SetDefault("githubUN", s)
		u.Z.Info("Result: ", zap.String("githubUN", s))
	}
}

func (u *UI) DockerHub() {
	if u.Config.PROJECT.Dockerhub == "" {
		s := u.String(Q + " What is your dockerhub user name?")
		u.V.SetDefault("dockerhubUN", s)
		u.Z.Info("Result: ", zap.String("dockerhubUN", s))
	}
}

func (u *UI) Summary() {
	if u.Config.PROJECT.Summary == "" {
		s := u.String(Q + " What is a summary of your new app?")
		u.V.SetDefault("summary", s)
		u.Z.Info("Result: ", zap.String("summary", s))
	}
}

func (u *UI) Version() {
	if u.Config.PROJECT.Version == "" {

		s := u.String(Q + " What is the semantic version of this app")
		u.V.SetDefault("version", s)
		u.Z.Info("Result: ", zap.String("version", s))
	}
}

func (u *UI) Cloud() string {
	s := u.String(Q + " What Cloud Provider do you use? [aws/gcloud]")
	u.V.SetDefault("cloud", s)
	u.Z.Info("Result: ", zap.String("cloud", s))
	return s
}

//TODO: add inputs for all config fields and a func to run all of them to initialize config if viper cant find the values
