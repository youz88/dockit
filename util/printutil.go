package util

import (
	"fmt"
	"strconv"
)

// Color defines a single SGR Code
type Color int

// Foreground text colors
const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func PrintInfo(s string) {
	printWithColor(Green, s)
}

func PrintInfoFormat(format string, a ...any) {
	PrintInfo(fmt.Sprintf(format, a...))
}

func PrintError(s string) {
	printWithColor(Red, s)
}

func printWithColor(color Color, s string) {
	fmt.Printf("\033[0;%s;8m%s\033[0m\n", strconv.Itoa(int(color)), s)
}
