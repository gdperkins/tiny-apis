package main

import (
	"github.com/gdperkins/tiny-apis/web"
)

func main() {

	s := web.NewServer()

	s.Run()
}
