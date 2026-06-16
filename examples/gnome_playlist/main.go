// Command gnome_playlist sorts a music playlist by track duration using Gnome
// Sort.
//
// Gnome Sort works by moving an element backwards until it is in the right
// position, then stepping forward again — conceptually similar to a garden
// gnome organising flower pots by moving one pot at a time. It performs well
// on small lists and requires no extra memory.
package main

import (
	"fmt"
	"github.com/danielriddell21/ordinex"
)

func main() {
	type Track struct {
		Title    string
		Duration int // seconds
	}

	playlist := []Track{
		{"Bohemian Rhapsody", 354},
		{"Blinding Lights", 200},
		{"Hotel California", 391},
		{"Shape of You", 234},
		{"Stairway to Heaven", 482},
		{"Smells Like Teen Spirit", 301},
		{"Rolling in the Deep", 228},
	}

	durations := make([]int, len(playlist))
	for i, t := range playlist {
		durations[i] = t.Duration
	}

	fmt.Println("=== Playlist — Sorted by Duration ===")
	fmt.Println()

	sorted := ordinex.GnomeSorter{}.Sort(durations)

	fmt.Printf("%-3s  %-30s  %s\n", "#", "Title", "Duration")
	fmt.Println("---  ------------------------------  --------")
	for i, d := range sorted {
		for _, t := range playlist {
			if t.Duration == d {
				mins := d / 60
				secs := d % 60
				fmt.Printf("%-3d  %-30s  %d:%02d\n", i+1, t.Title, mins, secs)
				break
			}
		}
	}

	total := 0
	for _, d := range sorted {
		total += d
	}
	fmt.Printf("\nTotal runtime: %d:%02d\n", total/60, total%60)
}
