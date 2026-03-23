// Package main sorts a poker hand using Insertion Sort.
//
// Insertion Sort mirrors exactly how most people sort cards in their hand:
// pick up a card, slide it left past higher cards until it reaches the right
// position. For small, human-scale datasets this natural analogy translates
// directly into efficient code.
package main

import (
	"fmt"
	"github.com/danielriddell21/sortilege"
)

func main() {
	// Card values: 2–10 are face value; 11=J, 12=Q, 13=K, 14=A
	type Card struct {
		Value int
		Suit  string
	}

	hand := []Card{
		{10, "♠"},
		{14, "♥"}, // Ace
		{7, "♦"},
		{10, "♣"},
		{13, "♠"}, // King
	}

	values := make([]int, len(hand))
	for i, c := range hand {
		values[i] = c.Value
	}

	label := func(v int) string {
		switch v {
		case 11:
			return "J"
		case 12:
			return "Q"
		case 13:
			return "K"
		case 14:
			return "A"
		default:
			return fmt.Sprintf("%d", v)
		}
	}

	fmt.Println("=== Texas Hold'em — Sort Your Hand ===")
	fmt.Println()

	fmt.Print("Dealt:  ")
	for _, c := range hand {
		fmt.Printf("%s%s ", label(c.Value), c.Suit)
	}
	fmt.Println()

	sorted := sortilege.InsertionSorter{}.Sort(values)

	fmt.Print("Sorted: ")
	for _, v := range sorted {
		// Find a matching card from the hand for the suit.
		for _, c := range hand {
			if c.Value == v {
				fmt.Printf("%s%s ", label(c.Value), c.Suit)
				break
			}
		}
	}
	fmt.Println()
	fmt.Println()

	// Detect pairs.
	counts := map[int]int{}
	for _, v := range sorted {
		counts[v]++
	}
	pairs := 0
	for _, count := range counts {
		if count == 2 {
			pairs++
		}
	}

	switch {
	case pairs == 2:
		fmt.Println("Hand: Two pair")
	case pairs == 1:
		fmt.Println("Hand: One pair")
	default:
		fmt.Println("Hand: High card —", label(sorted[len(sorted)-1]))
	}
}
