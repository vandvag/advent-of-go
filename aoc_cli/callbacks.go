package aoccli

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	promptlist "github.com/manifoldco/promptui/list"
	"github.com/urfave/cli/v3"
	"github.com/vandvag/advent-of-go/aoc"
	"github.com/vandvag/advent-of-go/registry"
	"github.com/vandvag/advent-of-go/solver"
)

const (
	exitLabel = "exit"
	backLabel = "back"
)

func runCb(ctx context.Context, cmd *cli.Command) error {
	day := cmd.Int(dayFlag)
	year := cmd.Int(yearFlag)

	if year == 0 && day == 0 {
		menu(ctx, cmd)
	} else {
		if day == 0 && year != 0 {
			return fmt.Errorf("You haven't specified a day")
		}

		if year == 0 && day != 0 {
			return fmt.Errorf("You haven't specified a year")
		}

		ok := aoc.ValidYear(year)
		if !ok {
			return fmt.Errorf("Year %d is not valid! Please enter a year between 2015 and 2024\n", year)
		}

		ok = aoc.ValidDay(day)
		if !ok {
			return fmt.Errorf("Day %d is not valid! Please enter a day between 1 and 25\n", year)
		}

		solution, ok := registry.Get(year, day)
		if !ok {
			return fmt.Errorf("Somehow solution for %d/%d wasn't registered\n", year, day)
		}

		err := solver.Solve(solution, cmd.Bool(elapsedFlag))
		if err != nil {
			return err
		}
	}

	return nil
}

func menu(_ context.Context, cmd *cli.Command) error {
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

		err = solver.Solve(solution, cmd.Bool(elapsedFlag))
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
