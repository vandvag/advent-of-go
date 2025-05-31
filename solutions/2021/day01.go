package solutions_2021

import (
	"fmt"
	"strconv"

	"github.com/vandvag/advent-of-go/aoc"
	"github.com/vandvag/advent-of-go/registry"
)

type Day01 struct{}

func (d Day01) Part1(input string) string {
	numbers, err := aoc.MapLine(input, func(line string) (int, error) {
		return strconv.Atoi(line)
	})

	if err != nil {
		// TODO: Part1 needs to throw error
		panic(err)
	}

	count := 0
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] < numbers[i] {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

func (d Day01) Part2(input string) string {
	numbers, err := aoc.MapLine(input, func(line string) (int, error) {
		return strconv.Atoi(line)
	})

	if err != nil {
		// TODO: Part1 needs to throw error
		panic(err)
	}

	count := 0
	for i := range len(numbers)-3 {
		if numbers[i]+numbers[i+1]+numbers[i+2] < numbers[i+1]+numbers[i+2]+numbers[i+3] {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

func (d Day01) Day() string {
	return "01"
}

func (d Day01) Year() string {
	return year_2021
}

func day01() {
	registry.Register(2021, 1, Day01{})
}
