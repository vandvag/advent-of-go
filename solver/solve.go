package solver

import (
	"fmt"
	"strconv"

	"github.com/vandvag/advent-of-go/aoc"
)

type Solver interface {
	Part1(input string) string
	Part2(input string) string
	Day() string
	Year() string
}

func Solve(solver Solver) error {
	year, err := strconv.Atoi(solver.Year())
	if err != nil {
		return err
	}

	day, err := strconv.Atoi(solver.Day())
	if err != nil {
		return err
	}

	input, err := aoc.GetInput(year, day)
	if err != nil {
		return err
	}

	fmt.Printf("Part 1: %s\n", solver.Part1(input))
	fmt.Printf("Part 2: %s\n", solver.Part2(input))

	return nil
}
