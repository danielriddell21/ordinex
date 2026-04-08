// Package main demonstrates Heap Sort applied to A&E triage prioritisation.
//
// Heap Sort's O(n log n) worst-case guarantee and in-place operation make it
// suitable for safety-critical systems where predictable performance matters.
// Here it orders patients by triage score so clinical staff see the most
// urgent cases first.
package main

import (
	"fmt"
	"github.com/danielriddell21/ordinex"
)

func main() {
	type Patient struct {
		Name  string
		Score int // higher = more urgent
	}

	patients := []Patient{
		{"Morgan, J.", 42},
		{"Patel, A.", 91},
		{"Chen, R.", 17},
		{"O'Brien, S.", 78},
		{"Kowalski, T.", 55},
		{"Nakamura, Y.", 88},
		{"Fischer, L.", 33},
	}

	scores := make([]int, len(patients))
	for i, p := range patients {
		scores[i] = p.Score
	}

	fmt.Println("=== A&E Triage Queue ===")
	fmt.Println()

	sorted := ordinex.HeapSorter{}.Sort(scores)

	urgency := func(score int) string {
		switch {
		case score >= 80:
			return "IMMEDIATE"
		case score >= 60:
			return "URGENT"
		case score >= 40:
			return "SEMI-URGENT"
		default:
			return "NON-URGENT"
		}
	}

	fmt.Printf("%-5s  %-16s  %-6s  %s\n", "Queue", "Patient", "Score", "Category")
	fmt.Println("-----  ----------------  ------  -----------")

	position := 1
	for i := len(sorted) - 1; i >= 0; i-- {
		for _, p := range patients {
			if p.Score == sorted[i] {
				fmt.Printf("%-5d  %-16s  %-6d  %s\n", position, p.Name, p.Score, urgency(p.Score))
				position++
				break
			}
		}
	}
}
