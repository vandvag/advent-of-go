package registry

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/vandvag/advent-of-go/solver"
)

var funcRegistry = make(map[int]map[int]solver.Solver)

func Register(year, day int, fn solver.Solver) {
	if funcRegistry[year] == nil {
		funcRegistry[year] = make(map[int]solver.Solver)
	}
	funcRegistry[year][day] = fn
}

func Get(year, day int) (solver.Solver, bool) {
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
