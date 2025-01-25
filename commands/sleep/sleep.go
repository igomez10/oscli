package sleep

import (
	"context"
	"time"

	"github.com/urfave/cli/v3"
)

func GetCmd() *cli.Command {
	return &cli.Command{
		Name:  "sleep",
		Usage: "sleep for a specified number of seconds",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "seconds",
				Aliases: []string{"s"},
				Usage:   "number of seconds to sleep",
				Value:   1,
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			seconds := c.Int("seconds")
			time.Sleep(time.Duration(seconds) * time.Second)
			return nil
		},
	}
}
