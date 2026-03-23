// Package main uses Selection Sort to find the cheapest items in a basket.
//
// Selection Sort repeatedly selects the minimum element and places it at the
// front. This makes it intuitive for scenarios where you want to extract the
// k smallest (or largest) values — you can stop after k passes rather than
// sorting the entire collection.
package main

import (
	"fmt"
	"github.com/danielriddell21/sortilege"
)

func main() {
	type Item struct {
		Name  string
		Price int // pence
	}

	basket := []Item{
		{"Pasta (500g)", 149},
		{"Olive Oil (500ml)", 449},
		{"Tinned Tomatoes", 89},
		{"Garlic Bulb", 65},
		{"Parmesan (100g)", 279},
		{"Fresh Basil", 120},
		{"Sea Salt", 175},
		{"Black Pepper", 199},
		{"Onion (3 pack)", 79},
		{"Chilli Flakes", 155},
	}

	prices := make([]int, len(basket))
	for i, item := range basket {
		prices[i] = item.Price
	}

	topN := 4

	fmt.Println("=== Bargain Finder — Cheapest Items ===")
	fmt.Println()

	sorted := sortilege.SelectionSorter{}.Sort(prices)

	fmt.Printf("%-22s  %s\n", "Item", "Price")
	fmt.Println("----------------------  -----")
	for _, item := range basket {
		fmt.Printf("%-22s  £%d.%02d\n", item.Name, item.Price/100, item.Price%100)
	}

	fmt.Printf("\nTop %d cheapest items:\n", topN)
	fmt.Println("----------------------  -----")
	for i := 0; i < topN; i++ {
		price := sorted[i]
		for _, item := range basket {
			if item.Price == price {
				fmt.Printf("%-22s  £%d.%02d\n", item.Name, item.Price/100, item.Price%100)
				break
			}
		}
	}

	saving := 0
	total := 0
	for _, p := range prices {
		total += p
	}
	for i := topN; i < len(sorted); i++ {
		saving += sorted[i]
	}
	fmt.Printf("\nTotal basket: £%d.%02d\n", total/100, total%100)
	fmt.Printf("Remove the %d priciest items and save: £%d.%02d\n", len(basket)-topN, saving/100, saving%100)
}
