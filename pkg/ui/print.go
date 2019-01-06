package ui

import (
	"github.com/kyokomi/emoji"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/dixonwille/wlog.v2"
	"io"
)

type printer struct {
	UI       *wlog.ConcurrentUI
	IO       *IO
	LogLevel func() zap.Config
	Closers  []func()
	Status   status
}

func NewPrinter() printer {
	i := DefaultIO()
	ui := wlog.New(i.In, i.Out, i.Err)

	pui := &wlog.PrefixUI{
		LogPrefix:     emoji.Sprint(":speech_balloon:"),
		OutputPrefix:  emoji.Sprint(":boom:"),
		SuccessPrefix: emoji.Sprint(":white_check_mark:"),
		InfoPrefix:    emoji.Sprint(":wave:"),
		ErrorPrefix:   emoji.Sprint(":x:"),
		WarnPrefix:    emoji.Sprint(":grimacing:"),
		RunningPrefix: emoji.Sprint(":fire:"),
		AskPrefix:     emoji.Sprint(":question:"),
		UI:            ui,
	}

	cui := wlog.AddConcurrent(pui)

	wlog.AddColor(wlog.Green, wlog.Red, wlog.BrightBlue, wlog.Blue, wlog.Yellow, wlog.BrightMagenta, wlog.Yellow, wlog.BrightGreen, wlog.BrightRed, cui)

	return printer{
		UI: cui,
		IO: i,
	}
}

func (r *printer) ChangeReaderWriter(reader io.Reader, writer, errorWriter io.Writer) {
	ui := wlog.New(reader, writer, errorWriter)

	pui := &wlog.PrefixUI{
		LogPrefix:     emoji.Sprint(":speech_balloon:"),
		OutputPrefix:  emoji.Sprint(":boom:"),
		SuccessPrefix: emoji.Sprint(":white_check_mark:"),
		InfoPrefix:    emoji.Sprint(":wave:"),
		ErrorPrefix:   emoji.Sprint(":x:"),
		WarnPrefix:    emoji.Sprint(":grimacing:"),
		RunningPrefix: emoji.Sprint(":fire:"),
		AskPrefix:     emoji.Sprint(":question:"),
		UI:            ui,
	}

	cui := wlog.AddConcurrent(pui)

	wlog.AddColor(wlog.Green, wlog.Red, wlog.BrightBlue, wlog.Blue, wlog.Yellow, wlog.BrightMagenta, wlog.Yellow, wlog.BrightGreen, wlog.BrightRed, cui)
	r.UI.UI = cui
}

// checks error and panics
func (r *printer) CheckErr(msg string, e error) {
	if e != nil {
		r.UI.Warn(msg)
		r.UI.Error(errors.WithStack(e).Error())
	}
}

// Close closes printer utilities.
func (u *printer) Close() {
	for _, f := range u.Closers {
		f()
	}
}

func (u *printer) Print(msg string) {
	switch u.Status {
	case StatusCreate:
		u.UI.Success("[Created] " + msg)
	case StatusForce:
		u.UI.Warn("[FORCED] " + msg)
	case StatusDelete:
		u.UI.Warn("[DELETED] " + msg)
	case StatusExist:
		u.UI.Info("[EXIST] ")
	case StatusConflicted:
		u.UI.Error("[CONFLICTED] " + msg)
	case StatusIdentical:
		u.UI.Warn("[IDENTICAL] " + msg)
	default:
		u.UI.Warn("[SKIPPED] " + msg)
	}
}

func (u *printer) ShouldCreate() bool {
	_, ok := creatableStatusSet[u.Status]
	return ok
}

func (u *printer) AddCloseFunc(f func()) {
	u.Closers = append(u.Closers, f)
}
