package main

import (
	"fmt"

	"github.com/b4b4r07/go-filter"
	"github.com/b4b4r07/go-filter/fzf"
	"github.com/b4b4r07/go-filter/peco"
)

var text = "hoge\nhoge\nhoge\n"

func main() {
	var (
		res []string
		err error
	)

	res, err = filter.Run(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res)

	res, err = peco.Run(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res)

	res, err = fzf.Run(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res)
}
