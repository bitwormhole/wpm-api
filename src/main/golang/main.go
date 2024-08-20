package main

import (
	"os"

	"github.com/bitwormhole/wpm-api/modules/wpmapi"
	"github.com/starter-go/starter"
)

func main() {
	m := wpmapi.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
