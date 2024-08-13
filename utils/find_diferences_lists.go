package utils

import (
	"fmt"
	"github.com/gocql/gocql"
)

func FindIdDiferencesBetweenLists(listA, listB []gocql.UUID) {
	// Convert the lists to sets for easier comparison
	set1 := make(map[gocql.UUID]bool)
	for _, id := range listA {
		set1[id] = true
	}

	set2 := make(map[gocql.UUID]bool)
	for _, id := range listB {
		set2[id] = true
	}

	// Find the difference by checking which IDs are in set1 but not in set2
	var difference []gocql.UUID
	for id := range set1 {
		if !set2[id] {
			difference = append(difference, id)
		}
	}

	// Print the difference
	fmt.Printf("IDs in list1 but not in list2: %v\n", difference)
}
