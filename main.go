package main

import (
	"os"

	"github.com/UnnoTed/horizontal"
	"github.com/endigma/basedfetch/based"
	"github.com/endigma/basedfetch/sys"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(horizontal.ConsoleWriter{Out: os.Stderr})
	s, _ := sys.Fetch()
	based.Compute(s)
}
