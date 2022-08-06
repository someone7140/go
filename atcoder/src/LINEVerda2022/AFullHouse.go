package main

import (
	"fmt"
)

func AFullHouseMain() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)

	value1 := a
	value1Count := 1
	value2 := -1
	value2Count := 0

	if value1 == b {
		value1Count = value1Count + 1
	} else {
		value2 = b
		value2Count = value2Count + 1
	}

	if value1 == c {
		value1Count = value1Count + 1
	} else if value2 == c || value2 == -1 {
		value2 = c
		value2Count = value2Count + 1
	}

	if value1 == d {
		value1Count = value1Count + 1
	} else if value2 == d || value2 == -1 {
		value2 = d
		value2Count = value2Count + 1
	}

	if value1 == e {
		value1Count = value1Count + 1
	} else if value2 == e || value2 == -1 {
		value2 = e
		value2Count = value2Count + 1
	}

	if (value1Count == 3 && value2Count == 2) || (value1Count == 2 && value2Count == 3) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
