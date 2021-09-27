package main

import (
	"flag"

	"github.com/pscompsci/pmtools/internal/pmtools"
)

func main() {
	cfg := pmtools.Config{}

	flag.IntVar(&cfg.Port, "port", 5050, "Port for the server to listen on")
	flag.StringVar(&cfg.Env, "env", "Development", "Application environment (Development, Production")
	flag.Parse()

	app := pmtools.New(&cfg)
	app.Run()
}
