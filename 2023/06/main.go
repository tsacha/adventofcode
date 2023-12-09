package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

func do_the_maths(a float64, b float64, c float64) (roots []float64, err error) {
	discriminant := b*b - (4 * a * c)

	if discriminant > 0 {
		x1 := (-b - math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b + math.Sqrt(discriminant)) / (2 * a)
		roots = []float64{x1, x2}
	} else {
		return nil, errors.New("nyi")
	}

	return roots, nil
}

func beat_the_record(time int, distance int) (solution int) {
	roots, err := do_the_maths(float64(-1), float64(time), float64(-distance))
	if err != nil {
		log.Fatalf("wtf")
	}

	max := math.Ceil(math.Max(roots[0], roots[1]))
	min := math.Floor(math.Min(roots[0], roots[1]))

	return int(max - min - 1)
}

func part1() (solution int) {
	solution = 1

	input := string(utils.PuzzleInput(2023, 6))
	input_str := strings.Split(input, "\n")
	text_times_str, text_distances_str := input_str[0], input_str[1]
	times_str, distances_str := strings.Fields(text_times_str)[1:], strings.Fields(text_distances_str)[1:]
	for n := 0; n < len(times_str); n++ {
		time, _ := strconv.Atoi(times_str[n])
		distance, _ := strconv.Atoi(distances_str[n])

		solution *= beat_the_record(time, distance)
	}
	return solution
}

func part2() (solution int) {
	input := string(utils.PuzzleInput(2023, 6))
	input_str := strings.Split(input, "\n")
	text_times_str, text_distances_str := input_str[0], input_str[1]
	times_str, distances_str := strings.Fields(text_times_str)[1:], strings.Fields(text_distances_str)[1:]
	time_str, distance_str := strings.Join(times_str, ""), strings.Join(distances_str, "")

	time, _ := strconv.Atoi(time_str)
	distance, _ := strconv.Atoi(distance_str)

	return beat_the_record(time, distance)
}

func main() {
	fmt.Printf("Solution 1: %d\n", part1())
	fmt.Printf("Solution 2: %d\n", part2())

}
