package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func PuzzleInput(year int, day int) (input []byte) {
	session_id := os.Getenv("AOC_SESSION")
	if session_id == "" {
		fmt.Println("No session ID (AOC_SESSION)")
		os.Exit(1)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		fmt.Println("Error while creating request")
	}
	req.Header.Add("Cookie", "session="+session_id)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error while fetching calibration data")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body
}
