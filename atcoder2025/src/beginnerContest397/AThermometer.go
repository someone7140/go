package main

import (
	"fmt"
)

func AThermometerMain() {
	var x float32
	fmt.Scan(&x)

	if x >= 38 {
		fmt.Println(1)
	} else if x >= 37.5 {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}

}
