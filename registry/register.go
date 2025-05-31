package registry

import (
	"fmt"
	"sort"
	"strconv"
)

type Solution interface {
	Part1(input string) string
	Part2(input string) string
}

var funcRegistry = make(map[int]map[int]Solution)

func Register(year, day int, fn Solution) {
	if funcRegistry[year] == nil {
		funcRegistry[year] = make(map[int]Solution)
	}
	funcRegistry[year][day] = fn
}

func Get(year, day int) (Solution, bool) {
	yearMap, ok := funcRegistry[year]
	if !ok {
		return nil, false
	}

	fn, ok := yearMap[day]

	return fn, ok
}

func RegisteredYears() []string {
	registeredYears := make([]string, 0, len(funcRegistry))

	for year := range funcRegistry {
		registeredYears = append(registeredYears, strconv.Itoa(year))
	}

	sort.Strings(registeredYears)
	return registeredYears
}

func RegisteredDays(year string) []string {
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		panic("You aren't supposed to send anything else here")
	}

	registeredDays := make([]string, 0, len(funcRegistry[yearInt]))

	for day := range funcRegistry[yearInt] {
		dayStr := fmt.Sprintf("%02d", day)
		registeredDays = append(registeredDays, dayStr)
	}

	return registeredDays
}
