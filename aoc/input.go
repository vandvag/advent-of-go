package aoc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

const baseURL = "https://adventofcode.com/%d/day/%d/input"

func GetInput(year int, day int) (string, error) {
	if year > 2024 || year < 2015 {
		return "", fmt.Errorf("Invalid year: %d\nAdvent of code started in 2015!\n", year)
	}

	if day <= 0 || day > 25 {
		return "", fmt.Errorf("Invalid day: %d\nAdvent of code runs from 1st till the 25th of December each year!\n", day)
	}

	file_path := filepath.Join("input", fmt.Sprintf("%d", year), fmt.Sprintf("%02d", day)+".in")

	data, err := os.ReadFile(file_path)

	if err == nil {
		return string(data), nil
	}

	content, err := getInputFromAOCSite(year, day)
	if err != nil {
		return "", err
	}

	err = writeInputToFile(file_path, content)
	if err != nil {
		return "", err
	}

	data, err = os.ReadFile(file_path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func writeInputToFile(file_path string, content []byte) error {
	dir := filepath.Dir(file_path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(file_path, content, 0644)
}

func getInputFromAOCSite(year int, day int) ([]byte, error) {
	if err := godotenv.Load(); err != nil {
		log.Printf("[WARN]: Failed to load .env file: %v", err)
	}

	url := fmt.Sprintf(baseURL, year, day)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %s\n", err)
	}

	session_cookie, exists := os.LookupEnv("AOC_SESSION")
	if !exists {
		return nil, fmt.Errorf("Variable AOC_SESSION was not found.\nPlease set up an environment variable, AOC_SESSION, with the session id cookie\n")
	}

	req.Header.Add("Cookie", "session="+session_cookie)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request to %s\n", url)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error reading respone from %s\n", url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response body")
	}

	return body, nil
}
