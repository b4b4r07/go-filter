package filter

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/b4b4r07/go-colon"
)

var Command = "fzf --multi:peco:fzy"

func Run(text string) ([]string, error) {
	var (
		lines         []string
		selectedLines []string
		buf           bytes.Buffer
		err           error
	)
	if text == "" {
		return lines, errors.New("no input")
	}
	err = runFilter(Command, strings.NewReader(text), &buf)
	if err != nil {
		return lines, err
	}
	if buf.Len() == 0 {
		return lines, errors.New("no lines selected")
	}
	lines = strings.Split(buf.String(), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		selectedLines = append(selectedLines, line)
	}
	return selectedLines, nil
}

func runFilter(command string, r io.Reader, w io.Writer) error {
	if command == "" {
		return errors.New("invalid argument")
	}
	command = os.Expand(command, os.Getenv)
	result, err := colon.Parse(command)
	if err != nil {
		return err
	}
	command = strings.Join(result.Executable().One().Attr.Args, " ")
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Stderr = os.Stderr
	cmd.Stdout = w
	cmd.Stdin = r
	return cmd.Run()
}
