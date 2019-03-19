package util

import (
	"errors"
	"os"
	"strconv"

	"golang.org/x/sys/unix"
)

func TerminalWidth() (width int, err error) {
	width, err = terminalWidthEnv()
	if err != nil {
		width, err = terminalWidthFd()
	}
	return
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
