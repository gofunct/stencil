package ui

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/tcnksm/go-input"
	"go.uber.org/zap"
	"strings"
)

type UI struct {
	io        IO
	inputUI   *input.UI
	inSection bool
	Closers   []func()
	*printer
}

func (u *UI) Section(msg string) {
	if u.inSection {
		fmt.Fprintln(u.io.Out)
		u.inSection = false
	}
	u.Print(msg)
}

func (u *UI) Subsection(msg string) {
	if u.inSection {
		fmt.Fprintln(u.io.Out)
		u.inSection = false
	}
	u.Print(msg)
}

func (u *UI) ItemSuccess(msg string) {
	u.inSection = true
	u.Print(msg)
}

func (u *UI) ItemSkipped(msg string) {
	u.inSection = true
	u.Print(msg)
}

func (u *UI) ItemFailure(msg string, errs ...error) {
	u.inSection = true
	u.Print(msg)

	fprintln := color.New(color.FgRed).FprintlnFunc()
	for _, err := range errs {
		for _, _ = range strings.Split(err.Error(), "\n") {
			fprintln(u.io.Out)
		}
	}
}

func (u *UI) Confirm(msg string) (bool, error) {
	ans, err := u.inputUI.Ask(fmt.Sprintf("%s [Y/n]", msg), &input.Options{
		HideOrder: true,
		Loop:      true,
		ValidateFunc: func(ans string) error {
			zap.L().Debug("receive user input", zap.String("query", msg), zap.String("input", ans))
			if ans != "Y" && ans != "n" {
				return fmt.Errorf("input must be Y or n")
			}
			return nil
		},
	})
	if err != nil {
		return false, errors.WithStack(err)
	}
	return ans == "Y", nil
}
