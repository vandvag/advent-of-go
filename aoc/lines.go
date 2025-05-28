package aoc

import (
	"bufio"
	"strings"
)

func MapLine[T any](input string, transform func(string) (T, error)) ([]T, error) {
	var results []T

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		res, err := transform(line)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func ForEachLine(input string, for_each func(string) error) error {
	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()
		if err := for_each(line); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
