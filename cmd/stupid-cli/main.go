package main

import (
	"flag"

	"github.com/adityalstkp/stupid-stuff"
)

var scrollDur int64

func init() {
	flag.Int64Var(&scrollDur, "scroll-dur", 90, "scroll duration")
	flag.Parse()
}

func main() {
	t := stupidstuff.NewTexter("jokes on you!")
	t.Scroll(scrollDur)
}
