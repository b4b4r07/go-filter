package fzf

import (
	"github.com/b4b4r07/go-filter"
)

func Run(text string) (selectedLines []string, err error) {
	filter.Command = "fzf"
	return filter.Run(text)
}
