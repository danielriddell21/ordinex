// Package main sorts network round-trip times using Cocktail Shaker Sort.
//
// Cocktail Shaker Sort is a bidirectional variant of Bubble Sort. It handles
// sequences where small values are clustered at the high end ("turtles") more
// efficiently than standard Bubble Sort, making it a reasonable choice for
// latency data that is partially sorted after a warm-up period.
package main

import (
	"fmt"
	"github.com/danielriddell21/ordinex"
)

func main() {
	// Round-trip times (ms) recorded from 15 consecutive pings to a CDN node.
	// The cache warmed up mid-sequence, so later readings are generally lower
	// but a few outliers remain scattered throughout.
	rtts := []int{
		142, 138, 145, 131, 127, 122, 118, 115, 160, 112,
		109, 119, 106, 103, 98,
	}

	fmt.Println("=== CDN Latency Report — Node eu-west-2 ===")
	fmt.Println()
	fmt.Printf("Raw RTTs (ms):    %v\n", rtts)

	sorted := ordinex.CocktailShakerSorter{}.Sort(rtts)

	fmt.Printf("Sorted RTTs (ms): %v\n\n", sorted)

	n := len(sorted)
	min, max := sorted[0], sorted[n-1]
	p50 := sorted[n/2]
	p95 := sorted[int(float64(n)*0.95)]

	fmt.Printf("min:  %d ms\n", min)
	fmt.Printf("p50:  %d ms\n", p50)
	fmt.Printf("p95:  %d ms\n", p95)
	fmt.Printf("max:  %d ms\n", max)
	fmt.Printf("jitter (max-min): %d ms\n", max-min)
}
