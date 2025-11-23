package main

import (
	"fmt"
)

func AHappyBirthday4Main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)

	if y >= x {
		fmt.Println("No")
		return
	}

	if x/y == z && x%y == 0 {
		fmt.Println("Yes")
		return
	}

	result := "No"

	for {
		y = y + 1
		x = x + 1
		if x/y == z && x%y == 0 {
			result = "Yes"
			break
		}
		if x/y < z {
			result = "No"
			break
		}
	}
	fmt.Println(result)
}
