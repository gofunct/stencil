package logging

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func ErrWithResolution(msg, resolution string, er error) {
	zap.L().Fatal(msg, zap.String("resolution", resolution), zap.Error(errors.WithStack(er)))
}

func Er(msg string, er error) {
	zap.L().Fatal(msg, zap.Error(er))
}

func Exit(msg string) {
	zap.L().Fatal(msg)
}
