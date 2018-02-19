package http

import (
	"fmt"
	"flag"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/subcommands"
	"{{cookiecutter.app_name}}/config"
	"{{cookiecutter.app_name}}/logging"
)

// Command :
type Command struct{}

// Name :
func (*Command) Name() string {
	return "server"
}

// Synopsis :
func (*Command) Synopsis() string {
	return "Run server"
}

// Usage :
func (*Command) Usage() string {
	return `server
  Run server.
`
}

// SetFlags :
func (*Command) SetFlags(f *flag.FlagSet) {

}

// Execute :
func (p *Command) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	logger := logging.GetLogger()
	conf := config.Configuration.HTTP
	address := fmt.Sprintf("%v:%v", conf.Host, conf.Port)
	if !config.Configuration.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()
	route(engine)

	logger.Infof("{{cookiecutter.app_name}} listening on %v", address)
	engine.Run(address)
	return subcommands.ExitSuccess
}