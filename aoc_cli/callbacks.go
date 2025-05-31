package aoccli

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"
	promptlist "github.com/manifoldco/promptui/list"
	"github.com/urfave/cli/v3"
)

const exitLabel = "exit"

func findYearFolders() (map[string]bool, error) {
	years := make(map[string]bool)
	for y := 2015; y <= 2024; y++ {
		years[strconv.Itoa(y)] = true
	}

	var found map[string]bool = make(map[string]bool)
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && years[d.Name()] {
			found[d.Name()] = true
		}
		return nil
	})

	return found, err
}

func menu(ctx context.Context, cmd *cli.Command) error {
	for {
		years, err := findYearFolders()
		var years_vec []string
		for year := range years {
			years_vec = append(years_vec, year)
		}
		years_vec = append(years_vec, exitLabel)
		if err != nil {
			return err
		}

		prompt := promptui.Select{
			Label:    "Years",
			Items:    years_vec,
			Searcher: searcher(years_vec),
		}

		_, year, err := prompt.Run()
		if err != nil {
			return err
		}

		if year == exitLabel {
			return nil
		}

		day, err := daysMenu()
		if err != nil {
			return err
		}

		if day == exitLabel {
			return nil
		}

		if day == "back" {
			continue
		}

		fmt.Printf("You've selected year: %s, day: %s\n", year, day)
		return nil
	}
}

func daysMenu() (string, error) {
	days := []string{
		"01",
		"02",
		"back",
		"exit",
	}

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
