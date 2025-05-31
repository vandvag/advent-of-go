package aoc

func ValidYear(year int) bool {
	if year < 2015 || year > 2024 {
		return false
	}

	return true
}

func ValidDay(day int) bool {
	if day < 1 || day > 25 {
		return false
	}

	return true
}
