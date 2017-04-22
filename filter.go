package filter

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/b4b4r07/go-colon"
	"github.com/pkg/errors"
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
	results, err := colon.Parse(command)
	if err != nil {
		return err
	}
	result, err := results.Executable().First()
	if err != nil {
		return errors.Wrap(err, "available command should be specified")
	}
	command = result.Item
	if command == "" {
		return errors.New("command not specified")
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
