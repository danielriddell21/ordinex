// Package main sorts employee IDs using Radix Sort.
//
// Radix Sort achieves O(d*(n+k)) time — effectively linear for fixed-width
// integers — by sorting digit by digit rather than comparing whole values.
// Employee IDs, order numbers, and other fixed-format identifiers are ideal
// candidates because d (digit count) is constant and small.
package main

import (
	"fmt"
	"github.com/danielriddell21/sortilege"
)

func main() {
	type Employee struct {
		ID         int
		Name       string
		Department string
	}

	employees := []Employee{
		{100472, "Priya Sharma", "Engineering"},
		{100051, "James Okafor", "Finance"},
		{100389, "Lin Wei", "Engineering"},
		{100203, "Sara Johansson", "Product"},
		{100618, "Marcus Bell", "Engineering"},
		{100134, "Fatima Al-Hassan", "HR"},
		{100755, "Daniel Park", "Product"},
		{100027, "Niamh Murphy", "Finance"},
		{100563, "Raj Patel", "Engineering"},
		{100290, "Chloe Dubois", "HR"},
	}

	ids := make([]int, len(employees))
	for i, e := range employees {
		ids[i] = e.ID
	}

	fmt.Println("=== Employee Directory — Sorted by ID ===")
	fmt.Println()

	sorted := sortilege.RadixSorter{}.Sort(ids)

	fmt.Printf("%-10s  %-20s  %s\n", "ID", "Name", "Department")
	fmt.Println("----------  --------------------  -----------")

	for _, id := range sorted {
		for _, e := range employees {
			if e.ID == id {
				fmt.Printf("%-10d  %-20s  %s\n", e.ID, e.Name, e.Department)
				break
			}
		}
	}
}
