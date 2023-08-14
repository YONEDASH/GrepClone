package main

import "fmt"

type ANSI string

const (
	// Foreground colors
	Black   ANSI = "\033[30m"
	Red     ANSI = "\033[31m"
	Green   ANSI = "\033[32m"
	Yellow  ANSI = "\033[33m"
	Blue    ANSI = "\033[34m"
	Magenta ANSI = "\033[35m"
	Cyan    ANSI = "\033[36m"
	White   ANSI = "\033[37m"
	Reset   ANSI = "\033[0m"

	// Background colors
	BgBlack   ANSI = "\033[40m"
	BgRed     ANSI = "\033[41m"
	BgGreen   ANSI = "\033[42m"
	BgYellow  ANSI = "\033[43m"
	BgBlue    ANSI = "\033[44m"
	BgMagenta ANSI = "\033[45m"
	BgCyan    ANSI = "\033[46m"
	BgWhite   ANSI = "\033[47m"

	// Text formatting
	Bold          ANSI = "\033[1m"
	Italic        ANSI = "\033[3m"
	Underline     ANSI = "\033[4m"
	StrikeThrough ANSI = "\033[9m"
	Concealed     ANSI = "\033[8m"
	Blink         ANSI = "\033[5m"
)

func PrintReset() {
	fmt.Print(Reset)
}
