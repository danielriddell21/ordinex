// Package main sorts application log entries by timestamp using Shell Sort.
//
// Shell Sort generalises Insertion Sort by first sorting elements far apart,
// progressively reducing the gap. This eliminates long-distance disorder
// quickly, leaving a nearly-sorted sequence for the final insertion pass.
// It performs well on medium-sized datasets without the memory overhead of
// Merge Sort or the worst-case risk of naive Quick Sort.
package main

import (
	"fmt"
	"sortilege"
)

func main() {
	type LogEntry struct {
		Timestamp int // Unix ms, truncated for readability
		Level     string
		Message   string
	}

	// Log entries arriving out of order due to buffering across three services.
	logs := []LogEntry{
		{1714, "INFO", "Request received"},
		{1698, "DEBUG", "Cache miss — querying database"},
		{1721, "INFO", "Auth token validated"},
		{1705, "WARN", "Response time approaching threshold"},
		{1689, "INFO", "Connection established"},
		{1731, "ERROR", "Upstream timeout after 5000ms"},
		{1712, "DEBUG", "Query executed in 14ms"},
		{1694, "INFO", "TLS handshake complete"},
		{1727, "WARN", "Retry attempt 1/3"},
		{1701, "DEBUG", "Session resumed from cache"},
	}

	timestamps := make([]int, len(logs))
	for i, l := range logs {
		timestamps[i] = l.Timestamp
	}

	fmt.Println("=== Application Log — Chronological View ===")
	fmt.Println()

	sorted := sortilege.ShellSorter{}.Sort(timestamps)

	fmt.Printf("%-12s  %-7s  %s\n", "Timestamp", "Level", "Message")
	fmt.Println("------------  -------  -------------------------------")

	for _, ts := range sorted {
		for _, l := range logs {
			if l.Timestamp == ts {
				fmt.Printf("T+%-10d  %-7s  %s\n", l.Timestamp, l.Level, l.Message)
				break
			}
		}
	}

	errors := 0
	warnings := 0
	for _, l := range logs {
		switch l.Level {
		case "ERROR":
			errors++
		case "WARN":
			warnings++
		}
	}

	fmt.Printf("\nSummary: %d error(s), %d warning(s) across %d entries\n",
		errors, warnings, len(logs))
}
