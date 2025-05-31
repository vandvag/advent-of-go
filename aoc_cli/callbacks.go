package aoccli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	promptlist "github.com/manifoldco/promptui/list"
	"github.com/urfave/cli/v3"
	"github.com/vandvag/advent-of-go/registry"
	"github.com/vandvag/advent-of-go/solver"
)

const (
	exitLabel = "exit"
	backLabel = "back"
)

func menu(_ context.Context, _ *cli.Command) error {
	for {
		years := registry.RegisteredYears()
		years = append(years, exitLabel)
		prompt := promptui.Select{
			Label:    "Years",
			Items:    years,
			Searcher: searcher(years),
		}

		_, year, err := prompt.Run()
		if err != nil {
			return err
		}

		if year == exitLabel {
			return nil
		}

		day, err := daysMenu(year)
		if err != nil {
			return err
		}

		if day == exitLabel {
			return nil
		}

		if day == "back" {
			continue
		}

		yearInt, err := strconv.Atoi(year)
		if err != nil {
			return err
		}
		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return err
		}
		solution, ok := registry.Get(yearInt, dayInt)
		if !ok {
			return fmt.Errorf("Somehow solution for %d/%d wasn't registered\n", yearInt, dayInt)
		}

		err = solver.Solve(solution)
		if err != nil {
			return err
		}

		return nil
	}
}

func daysMenu(year string) (string, error) {
	days := registry.RegisteredDays(year)
	days = append(days, backLabel)
	days = append(days, exitLabel)

	prompt := promptui.Select{
		Label:    "Days",
		Items:    days,
		Searcher: searcher(days),
	}

	_, day, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return day, nil
}

func searcher(inputList []string) promptlist.Searcher {
	return func(input string, index int) bool {
		item := inputList[index]

		item = strings.ReplaceAll(strings.ToLower(item), " ", "")
		input = strings.ReplaceAll(strings.ToLower(input), " ", "")

		return strings.Contains(input, item)
	}
}
