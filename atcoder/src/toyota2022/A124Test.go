package main

import (
	"fmt"
)

func A124TestMain() {
	var a, b int
	fmt.Scan(&a, &b)

	mondai1Flag := false
	mondai2Flag := false
	mondai3Flag := false

	result := 0

	if a == 1 {
		mondai1Flag = true
	} else if a == 2 {
		mondai2Flag = true
	} else if a == 4 {
		mondai3Flag = true
	} else if a == 3 {
		mondai1Flag = true
		mondai2Flag = true
	} else if a == 5 {
		mondai1Flag = true
		mondai3Flag = true
	} else if a == 6 {
		mondai2Flag = true
		mondai3Flag = true
	} else if a == 7 {
		mondai1Flag = true
		mondai2Flag = true
		mondai3Flag = true
	}

	if b == 1 {
		mondai1Flag = true
	} else if b == 2 {
		mondai2Flag = true
	} else if b == 4 {
		mondai3Flag = true
	} else if b == 3 {
		mondai1Flag = true
		mondai2Flag = true
	} else if b == 5 {
		mondai1Flag = true
		mondai3Flag = true
	} else if b == 6 {
		mondai2Flag = true
		mondai3Flag = true
	} else if b == 7 {
		mondai1Flag = true
		mondai2Flag = true
		mondai3Flag = true
	}

	if mondai1Flag {
		result = result + 1
	}
	if mondai2Flag {
		result = result + 2
	}
	if mondai3Flag {
		result = result + 4
	}

	fmt.Println(result)
}
