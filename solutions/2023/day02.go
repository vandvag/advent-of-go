package solutions_2023

import "github.com/vandvag/advent-of-go/registry"

type Day02 struct{}

func (d Day02) Part1(input string) string {
	return "This day2, part1"
}

func (d Day02) Part2(input string) string {
	return "This day2, part2"
}

func (d Day02) Day() string {
	return "02"
}

func (d Day02) Year() string {
	return year_2023
}

func day02() {
	registry.Register(2023, 2, Day02{})
}
