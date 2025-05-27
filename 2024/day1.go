package aoc_2024

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/vandvag/advent-of-go/utils"
)

func part1(input string) int {
	var left []int
	var right []int

	utils.ForEachLine(input, func(line string) {
		tokens := strings.Split(line, "   ")

		left_int, err := strconv.Atoi(strings.TrimSpace(tokens[0]))
		if err != nil {
			// Handle it
			log.Fatal("[FATAL]: left WTF\n")
		}
		left = append(left, left_int)

		right_int, err := strconv.Atoi(strings.TrimSpace(tokens[1]))
		if err != nil {
			// Handle it
			log.Fatal("[FATAL]: right WTF\n")
		}
		right = append(right, right_int)
	})

	sort.Ints(left)
	sort.Ints(right)

	var res int
	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff >= 0 {
			res += diff
		} else {
			res += -diff
		}
	}

	return res
}

func Day01() {
	fmt.Println("Day ", "01")
	input, err := utils.ReadInput(2024, 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1: ", part1(input))
}
