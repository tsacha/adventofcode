package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

type report []int
type reports []report

func parse() reports {
	input := utils.PuzzleInput(2024, 2)
	reports := reports{}
	_ = input

	for _, line := range strings.Split(string(input), "\n") {
		report := report{}
		if line == "" {
			break
		}

		reportsStr := strings.Split(line, " ")
		for _, levelStr := range reportsStr {
			level, _ := strconv.Atoi(levelStr)
			report = append(report, level)
		}

		reports = append(reports, report)
	}

	return reports
}

func isSafe(report report) bool {
	safe := true
	increasing := true
	if report[1]-report[0] < 0 {
		increasing = false
	}
	for n := 1; n < len(report); n++ {
		if report[n]-report[n-1] > 3 || report[n]-report[n-1] < -3 {
			safe = false
		}
		if increasing && report[n]-report[n-1] <= 0 {
			safe = false
		}
		if !increasing && report[n]-report[n-1] >= 0 {
			safe = false
		}
	}

	return safe
}

func part1() string {
	reports := parse()
	validReports := 0

	for _, report := range reports {
		if isSafe(report) {
			validReports++
		}
	}
	return fmt.Sprintf("%d\n", validReports)
}

func part2() string {
	reports := parse()
	validReports := 0

	for _, report := range reports {
		if isSafe(report) {
			validReports++
			continue
		}

		for n := range report {
			fixedReport := make([]int, len(report))
			copy(fixedReport, report)
			fixedReport = slices.Delete(fixedReport, n, n+1)

			if isSafe(fixedReport) {
				validReports++
				break
			}
		}

	}
	return fmt.Sprintf("%d\n", validReports)
}

func main() {
	fmt.Printf(part1())
	fmt.Printf(part2())
}
