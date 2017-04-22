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

func Run(text string) (selectedLines []string, err error) {
	if text == "" {
		return selectedLines, errors.New("no input")
	}
	var buf bytes.Buffer
	err = runFilter(Command, strings.NewReader(text), &buf)
	if err != nil {
		return selectedLines, err
	}
	if buf.Len() == 0 {
		return selectedLines, errors.New("no lines selected")
	}
	lines := strings.Split(buf.String(), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		selectedLines = append(selectedLines, line)
	}
	return selectedLines, nil
}

func runFilter(command string, r io.Reader, w io.Writer) error {
	command = os.Expand(command, os.Getenv)
	result, err := colon.Parse(command)
	if err != nil {
		return err
	}
	command = strings.Join(result.Executable().One().Attr.Args, " ")
	if command == "" {
		return errors.New("invalid command")
	}
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
