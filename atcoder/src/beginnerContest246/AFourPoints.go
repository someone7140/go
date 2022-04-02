package main

import (
	"fmt"
	"strconv"
)

func AFourPointsMain() {
	var x1, y1 int
	fmt.Scan(&x1, &y1)

	var x2, y2 int
	fmt.Scan(&x2, &y2)

	var x3, y3 int
	fmt.Scan(&x3, &y3)

	xMap := map[int]int{}
	yMap := map[int]int{}

	xMap[x1] = 1
	yMap[y1] = 1

	x2Count, x2ok := xMap[x2]
	if x2ok {
		xMap[x2] = x2Count + 1
	} else {
		xMap[x2] = 1
	}

	y2Count, y2ok := yMap[y2]
	if y2ok {
		yMap[y2] = y2Count + 1
	} else {
		yMap[y2] = 1
	}

	x3Count, x3ok := xMap[x3]
	if x3ok {
		xMap[x3] = x3Count + 1
	} else {
		xMap[x3] = 1
	}

	y3Count, y3ok := yMap[y3]
	if y3ok {
		yMap[y3] = y3Count + 1
	} else {
		yMap[y3] = 1
	}

	resultX := 0
	resultY := 0

	for k, v := range xMap {
		if v%2 != 0 {
			resultX = k
		}
	}

	for k, v := range yMap {
		if v%2 != 0 {
			resultY = k
		}
	}

	fmt.Println(strconv.Itoa(resultX) + " " + strconv.Itoa(resultY))
}
