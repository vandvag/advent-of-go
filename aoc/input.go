package aoc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const baseURL = "https://adventofcode.com/%d/day/%d/input"

func ReadInput(year int, day int) (string, error) {
	if year > 2024 || year < 2015 {
		return "", fmt.Errorf("Invalid year: %d", year)
	}

	if err := godotenv.Load(); err != nil {
		log.Printf("[WARN]: Failed to load .env file: %v", err)
	}

	url := fmt.Sprintf(baseURL, year, day)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error creating HTTP request: %s\n", err)
	}

	session_cookie, exists := os.LookupEnv("AOC_SESSION")
	if !exists {
		return "", fmt.Errorf("Variable AOC_SESSION was not found.\nPlease set up an environment variable, AOC_SESSION, with the session id cookie\n")
	}

	req.Header.Add("Cookie", "session="+session_cookie)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error sending request to %s\n", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error reading respone from %s\n", url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body")
	}

	return string(body), nil
}
