package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/k0kubun/pp"
	"github.com/tsacha/adventofcode/utils"
)

const (
	OPERATIONAL = 0
	DAMAGED     = 1
	UNKNOWN     = 2
)

type row map[int]int
type damaged []int
type springs struct {
	row     row
	damaged damaged
	unknown int
}

func parse_line(line string, repeat int) springs {
	row := make(row)
	unknown := 0

	f := strings.Fields(line)
	for i := 0; i < repeat; i++ {
		for n, c := range strings.Split(f[0], "") {
			switch c {
			case "#":
				row[i*len(f[0])+i+n] = DAMAGED
			case ".":
				row[i*len(f[0])+i+n] = OPERATIONAL
			case "?":
				row[i*len(f[0])+i+n] = UNKNOWN
				unknown++
			}
		}
	}

	s := springs{
		row:     row,
		unknown: unknown,
	}

	// if damaged string is in line, parse it
	if len(f) > 1 {
		for i := 0; i < repeat; i++ {
			for _, c := range strings.Split(f[1], ",") {
				n, _ := strconv.Atoi(c)
				s.damaged = append(s.damaged, n)
			}
		}
	} else {
		// else count it
		s.damage_recount()
	}

	return s
}

func (s *springs) pp_letters() {
	for n := 0; n < len(s.row); n++ {
		switch s.row[n] {
		case OPERATIONAL:
			fmt.Print(". ")
		case DAMAGED:
			fmt.Print("# ")
		case UNKNOWN:
			fmt.Print("? ")
		}
	}
	fmt.Println()
}

func (s *springs) pp_damaged() {
	for n := 0; n < len(s.damaged); n++ {
		fmt.Printf("%d ", s.damaged[n])
	}
	fmt.Println()
}

func (s *springs) bruteforce_possible_solutions() int {
	loop := int(math.Pow(2.0, float64(s.unknown)))

	var list_springs []springs
	for n := 0; n < loop; n++ {
		try := springs{
			row: make(row),
		}
		b := 0
		for p := 0; p < len(s.row); p++ {
			if s.row[p] == UNKNOWN {
				try.row[p] = (n >> (p - b)) & 1
			} else {
				try.row[p] = s.row[p]
				b++
			}
		}

		try.damage_recount()
		if reflect.DeepEqual(s.damaged, try.damaged) {
			list_springs = append(list_springs, try)
		}
	}

	return len(list_springs)
}

func (s *springs) damage_recount() {
	n := damaged{0}
	for v := 0; v < len(s.row); v++ {
		if s.row[v] == DAMAGED {
			n[len(n)-1]++
		} else if s.row[v] == OPERATIONAL {
			if n[len(n)-1] > 0 {
				n = append(n, 0)
			}
		}
	}
	if n[len(n)-1] == 0 {
		n = n[0 : len(n)-1]
	}
	s.damaged = n
}

func part1() (solution int) {
	var wg sync.WaitGroup
	input := string(utils.PuzzleInput(2023, 12))
	for _, line := range strings.Split(input, "\n") {
		if line != "" {
			wg.Add(1)
			go func(line string) {
				defer wg.Done()
				s := parse_line(line, 1)
				solution += s.bruteforce_possible_solutions()
			}(line)
		}
	}
	wg.Wait()
	return solution
}

func part2() (solution int) {
	s := parse_line(".#.###.#.###?##", 1)
	pp.Print(s)
	return solution
}

func main() {
	fmt.Printf("Solution 2: %d\n", part2())
}
