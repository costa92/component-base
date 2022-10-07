package term

import (
	"errors"
	"io"

	"github.com/moby/term"
)

var harmlessError = errors.New("given writer is no terminal")

func TerminalSize(w io.Writer) (width, height int, err error) {
	outFd, isTerminal := term.GetFdInfo(w)
	if !isTerminal {
		return width, height, harmlessError
	}
	winsize, err := term.GetWinsize(outFd)
	if err != nil {
		return width, height, err
	}
	width = int(winsize.Width)
	height = int(winsize.Height)
	return width, height, nil
}
