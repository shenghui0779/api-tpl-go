package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var Commands = []*cli.Command{
	{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "test cmd",
		Action: func(c *cli.Context) error {
			fmt.Println("this is a test cmd")

			return nil
		},
	},
}
