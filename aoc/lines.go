package aoc

import (
	"bufio"
	"strings"
)

func MapLine[T any](input string, transform func(string) T) ([]T, error) {
	var results []T

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, transform(line))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func ForEachLine(input string, for_each func(string)) error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		for_each(line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
