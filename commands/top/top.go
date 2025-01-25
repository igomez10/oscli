package top

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/urfave/cli/v3"
)

func GetCmd() *cli.Command {
	return &cli.Command{
		Name:  "top",
		Usage: "list the top 10 processes",
		Action: func(ctx context.Context, c *cli.Command) error {
			// Read the /proc directory
			files, err := ioutil.ReadDir("/proc")
			if err != nil {
				log.Fatalf("Error reading /proc directory: %v", err)
			}

			// Iterate over each file in /proc
			for _, file := range files {
				// Check if the file name is a PID (i.e., numeric)
				if pid, err := strconv.Atoi(file.Name()); err == nil {
					// Read the process name from /proc/<PID>/comm
					commPath := fmt.Sprintf("/proc/%d/comm", pid)
					comm, err := ioutil.ReadFile(commPath)
					if err == nil {
						fmt.Printf("\tPID: %d, Name: %s\n", pid, comm)
					}

					environ := fmt.Sprintf("/proc/%d/environ", pid)
					env, err := ioutil.ReadFile(environ)
					if err == nil {
						fmt.Printf("\tPID: %d, Environ: %s\n", pid, env)
					}

					// Read the process status from /proc/<PID>/status
					statusPath := fmt.Sprintf("/proc/%d/status", pid)
					status, err := ioutil.ReadFile(statusPath)
					if err == nil {
						fmt.Printf("\tPID: %d, Status: %s\n", pid, status)
					}
				}
			}
			return nil
		},
	}
}
