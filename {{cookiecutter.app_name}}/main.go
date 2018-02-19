package main

import (
    "os"
    "flag"
    "context"

	"github.com/google/subcommands"
    "{{cookiecutter.app_name}}/config"
)

func main() {
  subcommands.Register(subcommands.HelpCommand(), "")
  subcommands.Register(subcommands.FlagsCommand(), "")
  subcommands.Register(subcommands.CommandsCommand(), "")
  subcommands.Register(&config.VersionCommand{}, "")

  flag.Parse()
  ctx := context.Background()
  os.Exit(int(subcommands.Execute(ctx)))
}
