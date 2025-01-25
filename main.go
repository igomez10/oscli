package main

import (
	"context"
	"fmt"
	"os"

	"github.com/igomez10/oscli/commands/sleep"
	"github.com/igomez10/oscli/commands/top"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "oscli",
		Usage: "A cli to to common OS operations",
		Commands: []*cli.Command{
			top.GetCmd(),
			sleep.GetCmd(),
		},
		EnableShellCompletion: true,
		HideHelp:              false,
		HideVersion:           false,
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
