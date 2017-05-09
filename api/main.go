package main

import (
	"flag"

	"github.com/gdperkins/tiny-apis/config"
	"github.com/gdperkins/tiny-apis/web"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Parse()

	config.Init(*environment)

	s := web.NewServer()
	s.Run()
}
