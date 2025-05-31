package aoccli

import "github.com/urfave/cli/v3"

const (
	elapsedFlag      = "elapsed"
	elapsedFlagAlias = "e"
	dayFlag          = "day"
	dayFlagAlias     = "d"
	yearFlag         = "year"
	yearFlagAlias    = "y"
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

	day := cli.IntFlag{
		Name:        dayFlag,
		Aliases:     []string{dayFlagAlias},
		Usage:       "Run specific day",
		Required:    false,
		Hidden:      false,
		Destination: nil,
	}

	year := cli.IntFlag{
		Name:        yearFlag,
		Aliases:     []string{yearFlagAlias},
		Usage:       "Run specific year",
		Required:    false,
		Hidden:      false,
		Destination: nil,
	}

	flags = append(flags, &elapsed, &day, &year)
	return flags
}
