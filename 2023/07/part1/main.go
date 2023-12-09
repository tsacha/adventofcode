package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/tsacha/adventofcode/utils"
)

type hand []string

func card_value(card string) int {
	if n, err := strconv.Atoi(card); err == nil {
		return n
	} else {
		switch card {
		case "T":
			return 10
		case "J":
			return 11
		case "Q":
			return 12
		case "K":
			return 13
		case "A":
			return 14
		}
	}

	return 0
}

func count_cards(hand hand) (count_cards map[int]int) {
	count_cards = make(map[int]int)
	for _, c := range hand {
		count_cards[card_value(c)]++
	}
	return count_cards
}

func is_five_of_a_kind(hand hand) bool {
	if len(count_cards(hand)) == 1 {
		return true
	} else {
		return false
	}
}

func is_four_of_a_kind(hand hand) bool {
	for _, c := range count_cards(hand) {
		if c == 4 {
			return true
		}
	}

	return false
}
func is_full_house(hand hand) bool {
	for _, c := range count_cards(hand) {
		if c == 1 || c == 4 || c == 5 {
			return false
		}
	}
	return true
}
func is_three_of_a_kind(hand hand) bool {
	t := 0
	for _, c := range count_cards(hand) {
		if c == 2 || c == 4 || c == 5 {
			return false
		} else if c == 3 {
			t++
		}
	}
	if t > 0 {
		return true
	} else {
		return false
	}
}
func is_two_pair(hand hand) bool {
	pairs := 0
	for _, c := range count_cards(hand) {
		if c == 3 || c == 4 || c == 5 {
			return false
		} else if c == 2 {
			pairs++
		}
	}
	if pairs == 2 {
		return true
	}

	return false
}

func is_one_pair(hand hand) bool {
	pairs := 0
	for _, c := range count_cards(hand) {
		if c == 3 || c == 4 || c == 5 {
			return false
		} else if c == 2 {
			pairs++
		}
	}
	if pairs == 1 {
		return true
	}

	return false
}

func is_high_card(hand hand) bool {
	if len(count_cards(hand)) == 5 {
		return true
	} else {
		return false
	}
}

func hand_value(hand hand) int {
	if is_five_of_a_kind(hand) {
		return 7
	} else if is_four_of_a_kind(hand) {
		return 6
	} else if is_full_house(hand) {
		return 5
	} else if is_three_of_a_kind(hand) {
		return 4
	} else if is_two_pair(hand) {
		return 3
	} else if is_one_pair(hand) {
		return 2
	} else if is_high_card(hand) {
		return 1
	} else {
		return 0
	}
}

type hand_map struct {
	hand  hand
	value int
	bet   int
}

type game []hand_map

func (game game) sort_game(A int, B int) bool {
	if game[A].value > game[B].value {
		return true
	} else if game[A].value < game[B].value {
		return false
	} else if game[A].value == game[A].value {
		for n := range game[A].hand {
			if card_value(game[A].hand[n]) > card_value(game[B].hand[n]) {
				return true
			} else if card_value(game[A].hand[n]) < card_value(game[B].hand[n]) {
				return false
			}
		}
	}

	return false
}

func part1() (solution int) {
	input := string(utils.PuzzleInput(2023, 7))
	_ = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

	game := game{}
	for _, line := range strings.Split(string(input), "\n") {
		if line == "" {
			continue
		}
		l := strings.Fields(line)
		hand := strings.Split(l[0], "")
		bet, _ := strconv.Atoi(l[1])

		game = append(game, hand_map{
			hand:  hand,
			value: hand_value(hand),
			bet:   bet,
		})
	}

	sort.SliceStable(game, game.sort_game)

	fmt.Printf("%v\n", game)

	for n := range game {
		fmt.Printf("#%d %v %d %d\n", len(game)-n, game[n].hand, game[n].value, game[n].bet)
		solution += (len(game) - n) * game[n].bet
	}

	return solution
}

func main() {
	fmt.Printf("Solution 1: %d\n", part1())
}
