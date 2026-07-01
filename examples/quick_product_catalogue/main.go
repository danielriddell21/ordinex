// Command quick_product_catalogue sorts a product catalogue by price using Quick
// Sort.
//
// Quick Sort's average O(n log n) performance and cache-friendly in-place
// partitioning make it a practical default for general-purpose sorting.
// Here it orders a product catalogue so customers can browse from cheapest
// to most expensive.
package main

import (
	"fmt"
	"github.com/danielriddell21/ordinex"
)

func main() {
	type Product struct {
		Name  string
		Price int // pence
	}

	catalogue := []Product{
		{"Mechanical Keyboard", 8999},
		{"USB-C Hub", 2499},
		{"Monitor Stand", 3450},
		{"Webcam 1080p", 4999},
		{"Desk Mat", 1899},
		{"Cable Management Kit", 799},
		{"LED Desk Lamp", 2999},
		{"Laptop Stand", 3199},
		{"Noise-Cancelling Headphones", 14999},
		{"Mouse Pad XL", 1299},
	}

	prices := make([]int, len(catalogue))
	for i, p := range catalogue {
		prices[i] = p.Price
	}

	fmt.Println("=== Home Office Catalogue — Sorted by Price ===")
	fmt.Println()

	sorted := ordinex.QuickSorter[int]{}.Sort(prices)

	fmt.Printf("%-32s  %s\n", "Product", "Price")
	fmt.Println(fmt.Sprintf("%-32s  %s", "--------------------------------", "--------"))

	for _, price := range sorted {
		for _, p := range catalogue {
			if p.Price == price {
				fmt.Printf("%-32s  £%d.%02d\n", p.Name, p.Price/100, p.Price%100)
				break
			}
		}
	}

	total := 0
	for _, p := range sorted {
		total += p
	}
	fmt.Printf("\nTotal (all items): £%d.%02d\n", total/100, total%100)
}
