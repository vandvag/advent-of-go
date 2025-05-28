package aoc_2024

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/vandvag/advent-of-go/aoc"
	"github.com/vandvag/advent-of-go/mathematics"
)

func part1(input string) int {
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

	return res
}

func part2(input string) int {
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

	return res
}

func Day01() {
	fmt.Println("Day ", "01")
	input, err := aoc.GetInput(2024, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: ", part1(input))
	fmt.Println("Part 2: ", part2(input))
}
