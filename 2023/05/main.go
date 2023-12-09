package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/tsacha/adventofcode/utils"
)

type almanach struct {
	seeds []int
	steps map[int][]almanach_map
}
type almanach_map struct {
	dr_start int
	sr_start int
	r_len    int
}

func transform_seed(step int, seed int, almap []almanach_map) int {
	for _, row := range almap {
		if seed >= row.sr_start && seed < row.sr_start+row.r_len {
			seed = seed - row.sr_start + row.dr_start
			break
		}
	}

	return seed
}

func part1() (solution int) {
	input := string(utils.PuzzleInput(2023, 5))

	var almanach almanach
	step := -1
	for _, l := range strings.Split(input, "\n") {
		if strings.Contains(l, "seeds: ") {
			line := strings.Split(l, ": ")
			seeds_str := strings.Fields(line[1])
			for _, seed_str := range seeds_str {
				seed, err := strconv.Atoi(seed_str)
				if err == nil {
					almanach.seeds = append(almanach.seeds, seed)
				}
			}
			almanach.steps = make(map[int][]almanach_map)
		} else if strings.Contains(l, "map:") {
			step++
		} else if len(l) > 0 {
			str_map := strings.Fields(l)

			dr_start, _ := strconv.Atoi(str_map[0])
			sr_start, _ := strconv.Atoi(str_map[1])
			r_len, _ := strconv.Atoi(str_map[2])

			almanach.steps[step] = append(almanach.steps[step], almanach_map{
				dr_start: dr_start,
				sr_start: sr_start,
				r_len:    r_len,
			})

		}
	}

	n_seeds := []int{}
	for _, seed := range almanach.seeds {
		n_seed := seed
		for step := 0; step < len(almanach.steps); step++ {
			n_seed = transform_seed(step, n_seed, almanach.steps[step])
		}
		n_seeds = append(n_seeds, n_seed)
	}
	return slices.Min(n_seeds)
}

type almanach2 struct {
	seeds map[int]int
	steps map[int][]almanach_map
}

func part2() (solution int) {
	input := string(utils.PuzzleInput(2023, 5))

	var almanach almanach2
	step := -1
	for _, l := range strings.Split(input, "\n") {
		if strings.Contains(l, "seeds: ") {
			line := strings.Split(l, ": ")
			seeds_str := strings.Fields(line[1])

			almanach.seeds = make(map[int]int)
			seed := 0
			for n, seed_str := range seeds_str {
				p, err := strconv.Atoi(seed_str)
				if err == nil {
					if n%2 == 0 {
						seed = p
					} else {
						almanach.seeds[seed] = p
					}
				}
			}
			almanach.steps = make(map[int][]almanach_map)

			fmt.Printf("Seeds: %v\n", almanach.seeds)
		} else if strings.Contains(l, "map:") {
			step++
		} else if len(l) > 0 {
			str_map := strings.Fields(l)

			dr_start, _ := strconv.Atoi(str_map[0])
			sr_start, _ := strconv.Atoi(str_map[1])
			r_len, _ := strconv.Atoi(str_map[2])

			almanach.steps[step] = append(almanach.steps[step], almanach_map{
				dr_start: dr_start,
				sr_start: sr_start,
				r_len:    r_len,
			})

		}
	}

	min_seed := math.MaxInt
	var wg sync.WaitGroup
	for seed_start, seed_range := range almanach.seeds {
		wg.Add(1)

		go func(seed_start int, seed_range int) {
			defer wg.Done()
			min_seed_range := transform_seed_range(seed_start, seed_range, almanach)
			if min_seed_range < min_seed {
				min_seed = min_seed_range
			}
		}(seed_start, seed_range)
	}
	wg.Wait()
	return min_seed
}

func transform_seed_range(seed_start int, seed_range int, almanach almanach2) (min_seed_range int) {
	min_seed_range = math.MaxInt
	for seed := seed_start; seed < (seed_start + seed_range); seed++ {
		n_seed := seed
		for step := 0; step < len(almanach.steps); step++ {
			n_seed = transform_seed(step, n_seed, almanach.steps[step])
		}
		if n_seed < min_seed_range {
			min_seed_range = n_seed
		}
	}
	return min_seed_range
}

func main() {
	//fmt.Printf("Solution: %d\n", part1())
	fmt.Printf("\nSolution: %d\n", part2())

}
