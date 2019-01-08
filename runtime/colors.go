package runtime

import (
	"github.com/fatih/color"
)

func BrightRed(args ...interface{}) string {
	return color.HiRedString("%s", args)
}

func BrightGreen(args ...interface{}) string {
	return color.HiGreenString("%s", args)
}

func BrightWhite(args ...interface{}) string {
	return color.HiWhiteString("%s", args)
}

func BrightBlack(args ...interface{}) string {
	return color.HiBlackString("%s", args)
}

func BrightCyan(args ...interface{}) string {
	return color.HiCyanString("%s", args)
}

func BrightMagenta(args ...interface{}) string {
	return color.HiMagentaString("%s", args)
}

func BrightBlue(args ...interface{}) string {
	return color.HiBlueString("%s", args)
}

///

func Red(args ...interface{}) string {
	return color.RedString("%s", args)
}

func Green(args ...interface{}) string {
	return color.GreenString("%s", args)
}

func White(args ...interface{}) string {
	return color.WhiteString("%s", args)
}

func Black(args ...interface{}) string {
	return color.BlackString("%s", args)
}

func Magenta(args ...interface{}) string {
	return color.MagentaString("%s", args)
}

func Blue(args ...interface{}) string {
	return color.BlueString("%s", args)
}

func Cyan(args ...interface{}) string {
	return color.CyanString("%s", args)
}
