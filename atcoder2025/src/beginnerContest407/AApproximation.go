package main

import (
	"fmt"
)

func AApproximationMain() {
	var a, b float32
	fmt.Scan(&a, &b)

	wari1 := a / b
	wariSeisuu := float32((int(a) / int(b)))
	wariSeisuuPlusOne := float32(wariSeisuu + 1)
	if (wari1 - wariSeisuu) < (wariSeisuuPlusOne - wari1) {
		fmt.Println(wariSeisuu)
	} else {
		fmt.Println(wariSeisuuPlusOne)
	}

}
