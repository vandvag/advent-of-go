package solutions_2024

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/vandvag/advent-of-go/aoc"
	"github.com/vandvag/advent-of-go/mathematics"
	"github.com/vandvag/advent-of-go/registry"
)

const day01Str = "01"

type Day01 struct{}

func (d Day01) Part1(input string) string {
	var left []int
	var right []int

	aoc.ForEachLine(input, func(line string) error {
		tokens := strings.Split(line, "   ")

		if len(tokens) != 2 {
			return fmt.Errorf("Line: ")
		}

		left_int, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
		if err != nil {
			return fmt.Errorf("Couldn't convert first token (%s) to int", tokens[0])
		}
		left = append(left, left_int)

		right_int, err := strconv.Atoi(strings.TrimSpace(tokens[1]))
		if err != nil {
			return fmt.Errorf("Couldn't convert second token (%s) to int", tokens[1])
		}
		right = append(right, right_int)

		return nil
	})

	sort.Ints(left)
	sort.Ints(right)

	var res int
	for i, num := range left {
		diff := num - right[i]
		res += mathematics.AbsInt(diff)
	}

	return fmt.Sprintf("%d", res)
}

func (d Day01) Part2(input string) string {
	var left []int
	right := make(map[int]int)

	aoc.ForEachLine(input, func(line string) error {
		tokens := strings.Split(line, "   ")

		if len(tokens) != 2 {
			return fmt.Errorf("Line: ")
		}

		left_int, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
		if err != nil {
			return fmt.Errorf("Couldn't convert first token (%s) to int", tokens[0])
		}
		left = append(left, left_int)

		right_int, err := strconv.Atoi(strings.TrimSpace(tokens[1]))
		if err != nil {
			return fmt.Errorf("Couldn't convert second token (%s) to int", tokens[1])
		}
		right[right_int]++

		return nil
	})

	var res int
	for _, num := range left {
		res += num * right[num]
	}

	return fmt.Sprintf("%d", res)
}

func (d Day01) Day() string {
	return day01Str
}

func (d Day01) Year() string {
	return year
}

func day01() {
	registry.Register(2024, 01, Day01{})
}
