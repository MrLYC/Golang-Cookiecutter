package main

import (
	"os"
	"flag"
	"context"
	"fmt"

	"github.com/google/subcommands"
	"{{cookiecutter.app_path}}/config"
	"{{cookiecutter.app_path}}/http"
)

type initialHandler func() bool

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(&config.VersionCommand{}, "")
	subcommands.Register(&config.ConfInfoCommand{}, "")
	subcommands.Register(&http.Command{}, "")

	flag.StringVar(
	    &(config.Configuration.ConfigurationPath),
	    "c", config.Configuration.ConfigurationPath,
	    "Configuration file",
	)

	flag.Parse()

	initialHandlers := []initialHandler{
	    initRandomSeed,
	    initConfiguration,
	}

	for _, handler := range initialHandlers {
	    if !handler() {
	        os.Exit(255)
	    }
	}

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
