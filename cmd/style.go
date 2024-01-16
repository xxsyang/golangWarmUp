package cmd

import "fmt"

// Color ANSI escape codes for different text colors
const (
	ColorDefault = "\x1b[39m"
	ColorRed     = "\x1b[41m"
	ColorGreen   = "\x1b[42m"
	ColorCyan    = "\x1b[46m"
	ColorGray    = "\x1b[90m"
)

func red(s string) string {
	return fmt.Sprintf("%s%s%s", ColorRed, s, ColorDefault)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGreen, s, ColorDefault)
}

func cyan(s string) string {
	return fmt.Sprintf("%s%s%s", ColorCyan, s, ColorDefault)
}

func gray(s string) string {
	return fmt.Sprintf("%s%s%s", ColorGray, s, ColorDefault)
}
