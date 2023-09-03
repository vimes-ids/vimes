package main

import (
	"flag"
	cmd "vimes/cmd"
	"vimes/vimes"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

func main() {
	flag.Parse()

	app := cmd.NewApplication()

	app.AddCommand(cmd.NewHelpCommand(app))
	app.AddCommand(cmd.NewInitCommand())
	app.AddCommand(cmd.NewPortsCommand(vimes.NewPortFinder()))

	app.Run(flag.Args())

	config.WithOptions(config.ParseEnv)
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles("vimes.yml")

	if err != nil {
		panic(err)
	}
}
