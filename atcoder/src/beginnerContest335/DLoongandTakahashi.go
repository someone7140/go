package main

import (
	"fmt"
	"strconv"
	"strings"
)

func DLoongandTakahashiMain() {
	var n int
	fmt.Scan(&n)

	xArrayArray := make([][]string, n)
	for i := 0; i < n; i++ {
		xArray := make([]string, n)
		xArrayArray[i] = xArray
	}

	type DLoongandTakahashiXy struct {
		x int
		y int
	}

	tempNum := n*n - 1
	now := DLoongandTakahashiXy{
		x: (n+1)/2 - 1,
		y: (n+1)/2 - 1,
	}
	direction := "L"
	xArrayArray[now.y][now.x] = "T"

	for {
		if tempNum < 1 {
			break
		} else {
			if direction == "L" {
				now = DLoongandTakahashiXy{
					x: now.x - 1,
					y: now.y,
				}
				xArrayArray[now.y][now.x] = strconv.FormatInt(int64(tempNum), 10)
				if now.y < n-1 && xArrayArray[now.y+1][now.x] == "" {
					direction = "D"
				}
			} else if direction == "D" {
				now = DLoongandTakahashiXy{
					x: now.x,
					y: now.y + 1,
				}
				xArrayArray[now.y][now.x] = strconv.FormatInt(int64(tempNum), 10)
				if now.x < n-1 && xArrayArray[now.y][now.x+1] == "" {
					direction = "R"
				}
			} else if direction == "R" {
				now = DLoongandTakahashiXy{
					x: now.x + 1,
					y: now.y,
				}
				xArrayArray[now.y][now.x] = strconv.FormatInt(int64(tempNum), 10)
				if now.y > 0 && xArrayArray[now.y-1][now.x] == "" {
					direction = "U"
				}
			} else if direction == "U" {
				now = DLoongandTakahashiXy{
					x: now.x,
					y: now.y - 1,
				}
				xArrayArray[now.y][now.x] = strconv.FormatInt(int64(tempNum), 10)
				if now.x > 0 && xArrayArray[now.y][now.x-1] == "" {
					direction = "L"
				}
			}
			tempNum = tempNum - 1
		}
	}

	var resultSlice []string
	for i := 0; i < n; i++ {
		result := strings.Join(xArrayArray[i], " ")
		resultSlice = append(resultSlice, result)
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}
