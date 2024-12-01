package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//go:embed template.go
var template string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var day, year int
	var cookie string

	today := time.Now()
	flag.IntVar(&day, "day", today.Day(), "day number to fetch, 1-25")
	flag.IntVar(&year, "year", today.Year(), "day of year to fetch")
	flag.StringVar(&cookie, "cookie", os.Getenv("AOC_SESSION"), "session cookie from website")
	flag.Parse()

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	prefixedDay := fmt.Sprintf("day_%02d", day)

	fmt.Printf("Sending request to %s...\n", url)
	response, err := sendGetRequest(url, cookie)
	if err != nil {
		log.Fatal(err)
	}
	stringResponse := string(response)
	stringResponse = strings.TrimRight(stringResponse, "\n")
	fmt.Println("Received response!")

	err = writeToFile(filepath.Join(prefixedDay, "input.txt"), stringResponse)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	err = writeToFile(filepath.Join(prefixedDay, "main.go"), template)
	if err != nil {
		log.Fatal("Error writing to file:", err)
	}

	fmt.Println()
	fmt.Printf("Successfully created all template files for day %d of Advent of Code %d\n", day, year)
	fmt.Printf("You can view task description here: https://adventofcode.com/%d/day/%d\n", year, day)

}

func sendGetRequest(url, cookie string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: cookie,
	})

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = response.Body.Close()
	}()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func writeToFile(filePath, content string) error {
	if _, err := os.Stat(filePath); err == nil {
		fmt.Println("file already exists:", filePath)
		return nil
	}

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("creating directory: %w", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("writing to file: %w", err)
	}
	return nil
}
