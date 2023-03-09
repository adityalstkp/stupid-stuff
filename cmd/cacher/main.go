package main

import (
	"flag"

	stupidstuff "github.com/adityalstkp/stupid-stuff"
)

var listenAddr string

func init() {
	flag.StringVar(&listenAddr, "addr", "127.0.0.1:3000", "TCP listen address")
	flag.Parse()

}

func main() {
	c := stupidstuff.NewCache()
	s := stupidstuff.NewServer(stupidstuff.ServerOpts{ListenAddr: listenAddr}, c)

	s.StartAndListen()
}
