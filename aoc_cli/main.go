package aoccli

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

const (
	aocCLILabel = "aoc-cli"
)

func Run() {
	ctx := context.Background()

	cmd := &cli.Command{
		Name: aocCLILabel,
		Commands: []*cli.Command{
			{
				Name:    cmdRunLabel,
				Aliases: nil,
				Usage:   "Runs advent-of-go application",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return menu(ctx, cmd)
				},
			},
		},
		Description: "Solutions of puzzles for Advent of code (https://adventofcode.com)",
	}

	if err := cmd.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}

}
