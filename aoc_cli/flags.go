package aoccli

import "github.com/urfave/cli/v3"

const (
	elapsedFlag      = "elapsed"
	elapsedFlagAlias = "e"
)

func flags() []cli.Flag {
	var flags []cli.Flag

	elapsed := cli.BoolFlag{
		Name:        elapsedFlag,
		Aliases:     []string{elapsedFlagAlias},
		Usage:       "Enables elapsed time metric",
		Required:    false,
		Hidden:      false,
		Value:       false,
		DefaultText: "",
		Destination: nil,
	}

	flags = append(flags, &elapsed)
	return flags
}
