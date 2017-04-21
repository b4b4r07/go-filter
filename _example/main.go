package main

import (
	"fmt"

	"github.com/b4b4r07/go-filter"
)

var text = "hoge\nhoge\nhoge\n"

func main() {
	res, err := filter.Run(text)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", res)
}
