package main

import (
    "os"
    "flag"
    "context"
    "math/rand"
    "time"
    "fmt"

    "github.com/google/subcommands"
    "{{cookiecutter.app_name}}/logging"
    "{{cookiecutter.app_name}}/config"
    "{{cookiecutter.app_name}}/http"
)

type initialHandler func() bool

func initRandomSeed() bool {
    rand.Seed(time.Now().Unix())
    return true
}

func initConfiguration() bool {
    var err error
    err = config.Configuration.Read()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
    }

    err = config.Configuration.Validate()
    if err != nil {
        fmt.Fprintln(os.Stderr, err.Error())
        return false
    }
    return true
}

func initLogging() bool {
    logging.Init()
    return true
}

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
