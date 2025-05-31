package solver

import (
	"fmt"
	"strconv"
	"time"

	"github.com/vandvag/advent-of-go/aoc"
)

type Solver interface {
	Part1(input string) string
	Part2(input string) string
	Day() string
	Year() string
}

func Solve(solver Solver, showElapsed bool) error {
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

	if showElapsed {
		decorateWithElapsedTime(input, 1, solver.Part1)
		decorateWithElapsedTime(input, 2, solver.Part2)
	} else {
		fmt.Printf("Part 1: %s\n", solver.Part1(input))
		fmt.Printf("Part 2: %s\n", solver.Part2(input))
	}

	return nil
}

func decorateWithElapsedTime(input string, part int, fn func(string) string) {
	start := time.Now()
	res := fn(input)
	elapsed := time.Since(start)

	fmt.Printf("Part %d: %s\n", part, res)
	fmt.Printf("\telapsed: %s\n", elapsed)
}
