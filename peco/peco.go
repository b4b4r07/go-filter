package peco

import (
	"github.com/b4b4r07/go-filter"
)

func Run(text string) (selectedLines []string, err error) {
	filter.Command = "peco"
	return filter.Run(text)
}
