// Package main sorts a small game leaderboard using Bubble Sort.
//
// Bubble Sort's simplicity makes it a practical choice when the dataset is
// small and readability matters more than raw performance. For a five-player
// leaderboard that updates after each match, the overhead of a more complex
// algorithm is unnecessary.
package main

import (
	"fmt"
	"github.com/danielriddell21/sortilege"
)

func main() {
	type Player struct {
		Name  string
		Score int
	}

	players := []Player{
		{"Alice", 4820},
		{"Bob", 3150},
		{"Carol", 5600},
		{"Dave", 2980},
		{"Eve", 4100},
	}

	scores := make([]int, len(players))
	for i, p := range players {
		scores[i] = p.Score
	}

	fmt.Println("=== Weekly Leaderboard ===")
	fmt.Println()

	sorted := sortilege.BubbleSorter{}.Sort(scores)

	// Rebuild the leaderboard in descending order.
	fmt.Printf("%-5s %-10s %s\n", "Rank", "Player", "Score")
	fmt.Println("----  ----------  -----")
	rank := 1
	for i := len(sorted) - 1; i >= 0; i-- {
		for _, p := range players {
			if p.Score == sorted[i] {
				fmt.Printf("%-5d %-10s %d\n", rank, p.Name, p.Score)
				rank++
				break
			}
		}
	}
}
