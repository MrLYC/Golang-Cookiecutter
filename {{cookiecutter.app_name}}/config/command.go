package config

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

// VersionCommand :
type VersionCommand struct{}

// Name :
func (*VersionCommand) Name() string {
	return "version"
}

// Synopsis :
func (*VersionCommand) Synopsis() string {
	return "Print version infomations"
}

// Usage :
func (*VersionCommand) Usage() string {
	return `version
  Print version infomations.
`
}

// SetFlags :
func (*VersionCommand) SetFlags(f *flag.FlagSet) {
}

// Execute :
func (p *VersionCommand) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	fmt.Printf("Mode: %v, Version: %v, BuildHash: %v\n", Mode, Version, BuildHash)
	return subcommands.ExitSuccess
}
