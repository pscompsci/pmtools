package main

import (
	"flag"

	"github.com/pscompsci/pmtools/internal/pmtools"
)

func main() {
	cfg := pmtools.Config{}

	flag.IntVar(&cfg.Port, "port", 5050, "Port for the server to listen on")
	flag.Parse()

	app := pmtools.New(&cfg)
	app.Run()
}
