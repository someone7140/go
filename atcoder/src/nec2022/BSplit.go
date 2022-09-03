package main

import (
	"fmt"
)

func BSplitMain() {
	var s string
	fmt.Scan(&s)

	sArray := make([]string, 10)
	for i, c := range s {
		sArray[i] = string([]rune{c})
	}

	if sArray[0] == "1" {
		fmt.Println("No")
	} else {
		allTaoreFlagArray := make([]bool, 7)

		if sArray[6] == "0" {
			allTaoreFlagArray[0] = true
		}
		if sArray[3] == "0" {
			allTaoreFlagArray[1] = true
		}
		if sArray[7] == "0" && sArray[1] == "0" {
			allTaoreFlagArray[2] = true
		}
		if sArray[4] == "0" && sArray[0] == "0" {
			allTaoreFlagArray[3] = true
		}
		if sArray[8] == "0" && sArray[2] == "0" {
			allTaoreFlagArray[4] = true
		}
		if sArray[5] == "0" {
			allTaoreFlagArray[5] = true
		}
		if sArray[9] == "0" {
			allTaoreFlagArray[6] = true
		}

		result := "No"
		count := 0
		for i := 0; i < 7; i++ {
			if count == 0 {
				if !allTaoreFlagArray[i] {
					count = 1
				} else {
					count = 0
				}

			} else if count == 1 {
				if allTaoreFlagArray[i] {
					count = 2
				} else {
					count = 1
				}
			} else if count == 2 {
				if !allTaoreFlagArray[i] {
					result = "Yes"
					break
				} else {
					count = 2
				}
			}
		}
		fmt.Println(result)
	}

}
