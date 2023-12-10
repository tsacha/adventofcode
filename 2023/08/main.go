package main

import (
	"fmt"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

type path map[string]string
type way map[string]path

func part1() (solution int) {
	input := string(utils.PuzzleInput(2023, 8))
	_ = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

	_ = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

	lines := strings.Split(input, "\n")
	directions := strings.Split(lines[0], "")
	fmt.Printf("%v\n", directions)

	way := make(way)
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}

		l := strings.Split(line, " = ")

		destination := strings.Split(l[1], ", ")
		way[l[0]] = path{
			"L": destination[0][1:],
			"R": destination[1][:len(destination[1])-1],
		}
	}

	fmt.Printf("%v\n", way)

	n := 0
	next := "AAA"
	direction_ptr := 0
	for next != "ZZZ" {
		n++
		direction := directions[direction_ptr]
		next = way[next][direction]
		fmt.Printf("%v\n", next)
		direction_ptr = (direction_ptr + 1) % len(directions)
	}

	return n
}

func part2() (solution int) {
	input := string(utils.PuzzleInput(2023, 8))
	_ = `LR

11A = (11B, ZZZ)
11B = (ZZZ, 11Z)
11Z = (11B, ZZZ)
22A = (22B, ZZZ)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
ZZZ = (ZZZ, ZZZ)`

	lines := strings.Split(input, "\n")
	directions := strings.Split(lines[0], "")

	paths := []string{}
	way := make(way)
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}

		l := strings.Split(line, " = ")

		destination := strings.Split(l[1], ", ")
		way[l[0]] = path{
			"L": destination[0][1:],
			"R": destination[1][:len(destination[1])-1],
		}
		if string(l[0][2]) == "A" {
			paths = append(paths, l[0])
		}
	}

	n := []int{}

	fmt.Printf("%v\n", paths)
	for i, path := range paths {
		n = append(n, 0)
		fmt.Printf("%v\n", path)
		direction_ptr := 0
		next := path
		for string(next[2]) != "Z" {
			n[i]++
			direction := directions[direction_ptr]
			next = way[next][direction]
			direction_ptr = (direction_ptr + 1) % len(directions)
		}
	}

	fmt.Printf("%v\n", LCM(n[0], n[1], n[2:]...))

	return 0
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	fmt.Printf("Solution 1: %d\n", part2())
}
