package ui

import (
	"github.com/gofunct/stencil/pkg/iio"
	"github.com/kyokomi/emoji"
	"gopkg.in/dixonwille/wlog.v2"
	"io"
)

type printer struct {
	UI      *wlog.ConcurrentUI
	IO      *iio.IO
	Closers []func()
}

func NewPrinter() *printer {
	i := iio.DefaultIO()
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

	return &printer{
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

// Close closes printer utilities.
func (u *printer) Close() {
	for _, f := range u.Closers {
		f()
	}
}

func (u *printer) AddCloseFunc(f func()) {
	u.Closers = append(u.Closers, f)
}
