package util

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/sys/unix"
)

const (
	noColor     = "\033[0m"
	bold        = "\033[0;1m"
	black       = "\033[0;30m"
	red         = "\033[0;31m"
	green       = "\033[0;32m"
	yellow      = "\033[0;33m"
	blue        = "\033[0;34m"
	magenta     = "\033[0;35m"
	cyan        = "\033[0;36m"
	white       = "\033[0;37m"
	boldBlack   = "\033[1;30m"
	boldRed     = "\033[1;31m"
	boldGreen   = "\033[1;32m"
	boldYellow  = "\033[1;33m"
	boldBlue    = "\033[1;34m"
	boldMagenta = "\033[1;35m"
	boldCyan    = "\033[1;36m"
	boldWhite   = "\033[1;37m"
)

type Colors struct {
	Title   string
	Repo    string
	Version string
	NoColor string
}

var colors *Colors

func init() {
	colors = new(Colors)
	colors.Title = bold
	colors.Repo = boldMagenta
	colors.Version = boldGreen
	colors.NoColor = noColor
}

func TerminalWidth() (width int, err error) {
	width, err = terminalWidthEnv()
	if err != nil {
		width, err = terminalWidthFd()
	}
	return
}

func TerminalColors() Colors {
	return *colors
}

func terminalWidthEnv() (width int, err error) {
	widthString := os.Getenv("COLUMNS")
	if len(widthString) > 0 {
		width, err = strconv.Atoi(widthString)
	} else {
		width, err = -1, errors.New("Cannot get terminal width using COLUMNS environment variable")
	}
	return
}

func terminalWidthFd() (int, error) {
	ws, err := winSize()
	if err != nil {
		return 0, err
	}
	return int(ws.Col), nil
}

func winSize() (*unix.Winsize, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	return ws, nil
}
