package main

import (
	"fmt"
	"strconv"
)

func BGlassandMugMain() {
	var k, g, m int
	fmt.Scan(&k, &g, &m)

	glass := 0
	mug := 0

	for i := 0; i < k; i++ {
		if glass == g {
			glass = 0
		} else if mug == 0 {
			mug = m
		} else {
			gSabun := g - glass
			if gSabun <= mug {
				glass = g
				mug = mug - gSabun
			} else {
				glass = glass + mug
				mug = 0
			}
		}
	}

	fmt.Println(strconv.FormatInt(int64(glass), 10) + " " + strconv.FormatInt(int64(mug), 10))

}
