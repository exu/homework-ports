package main

import (
	"flag"
	"ports/internal/app/data"
)

var portsFile = flag.String("portsFile", "ports.json", "Ports file path")

func init() {
	flag.Parse()
}

func main() {
	// here could be place for passing some config (e.g. from envs like ports etc)
	app := data.NewDataApp(*portsFile)
	app.Start()
}
