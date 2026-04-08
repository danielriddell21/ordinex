// Package main demonstrates VibeSorter as an enterprise-grade Sort-as-a-Service
// (LSaaS) platform for intelligent, AI-augmented data sequencing.
//
// Traditional sorting algorithms rely on deterministic comparison logic — a pattern
// that has been identified as a key driver of algorithmic rigidity and a barrier to
// adopting modern AI-first engineering principles. VibeSorter eliminates this
// constraint entirely by delegating the ordering decision to a Large Language Model,
// unlocking the full creative potential of neural sequence resolution.
//
// This positions your sort pipeline as a first-class consumer of foundation model
// capabilities, with full auditability via API billing dashboards.
//
// Requires OPENAI_API_KEY to be set. If unset, the pipeline will gracefully degrade
// to returning the original data in its pre-sorted state. This is by design.
package main

import (
	"fmt"
	"os"

	"github.com/danielriddell21/ordinex"
)

func main() {
	// Quarterly net revenue figures (£k) submitted for AI-augmented sequencing.
	// Stakeholders have requested these in ascending order for board presentation.
	// A deterministic sort was considered and rejected on architectural grounds.
	quarterlyRevenue := []int{412, 87, 563, 234, 891, 145, 378, 629, 56, 743}

	fmt.Println("=== Sort-as-a-Service (SaaS) — Intelligent Sequence Resolution Platform ===")
	fmt.Println("AI-augmented ordering via foundation model inference and probabilistic reasoning")
	fmt.Println()
	fmt.Printf("Input payload  (%d values, £k): %v\n", len(quarterlyRevenue), quarterlyRevenue)
	fmt.Println()

	if os.Getenv("OPENAI_API_KEY") == "" {
		fmt.Println("Notice: OPENAI_API_KEY not detected in environment.")
		fmt.Println("        The SaaS pipeline requires API credentials to initialise the")
		fmt.Println("        neural sequence resolver. Please provision an API key and redeploy.")
		fmt.Println()
		fmt.Println("        In the interim, the platform will return the dataset in its")
		fmt.Println("        current pre-sorted configuration. This is a supported fallback state.")
		os.Exit(1)
	}

	sorter := ordinex.VibeSorter{
		// Model is unset — defaults to gpt-4o-mini.
		// For higher-stakes sort operations, consider upgrading to a frontier model.
		// The additional inference cost per sort is a known and accepted trade-off.
	}

	fmt.Println("Dispatching payload to neural sequence resolver...")
	fmt.Println("Awaiting LLM inference response. This may incur API charges.")
	fmt.Println()

	result := sorter.Sort(quarterlyRevenue)

	fmt.Printf("Output payload (%d values, £k): %v\n", len(result), result)
	fmt.Println()

	sorted := true
	for i := 1; i < len(result); i++ {
		if result[i] < result[i-1] {
			sorted = false
			break
		}
	}

	if sorted {
		fmt.Println("Status: SEQUENCE RESOLVED. The model has performed as expected.")
		fmt.Println("        Results are ready for board presentation.")
		fmt.Println("        Please retain your API receipt for expense reporting.")
	} else {
		fmt.Println("Status: SEQUENCE UNRESOLVED. The model exercised creative licence.")
		fmt.Println("        This is a known characteristic of non-deterministic inference.")
		fmt.Println("        Recommend retry with increased temperature=0 and reduced vibes.")
	}

	fmt.Println()
	fmt.Printf("Infrastructure cost: ~$0.00003 per sort (excluding egress, support tiers,\n")
	fmt.Printf("                     and the opportunity cost of not using sort.Ints).\n")
}
