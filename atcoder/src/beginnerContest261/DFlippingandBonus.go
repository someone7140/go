package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CoinPoint struct {
	kaisuu int
	score  int64
}

func DFlippingandBonusMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	xArray := make([]int64, n)
	xStrArray := strings.Split(DFlippingandBonusRdr(rdr), " ")
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(xStrArray[i])
		xArray[i] = int64(x)
	}

	cyMap := map[int]int64{}

	var maxBonusPoint int64
	maxBonusPoint = 0
	maxBonusKaisuu := 0
	maxSyou := 0
	maxAmari := 0

	for i := 0; i < m; i++ {
		cyStrArray := strings.Split(DFlippingandBonusRdr(rdr), " ")
		c, _ := strconv.Atoi(cyStrArray[0])
		y, _ := strconv.Atoi(cyStrArray[1])
		cyMap[c] = int64(y)

		syou := y / c
		amari := y % c

		if syou > maxSyou {
			maxSyou = syou
			maxAmari = amari
			maxBonusPoint = int64(y)
			maxBonusKaisuu = c
		} else if syou == maxSyou && maxAmari < amari {
			maxSyou = syou
			maxAmari = amari
			maxBonusPoint = int64(y)
			maxBonusKaisuu = c
		}
	}

	var result int64
	nowKaisuu := 0
	for i := 0; i < n; i++ {
		if nowKaisuu >= maxBonusKaisuu {
			// 残り回数
			nokori := n - 1 - i
			if nokori < maxBonusKaisuu {
				result = result + xArray[i]
				nowKaisuu = nowKaisuu + 1
			} else {
				var tempSum int64
				tempSum = 0
				for j := 0; j < maxBonusKaisuu; j++ {
					tempSum = tempSum + xArray[i+j]
				}
				if tempSum < maxBonusPoint {
					nowKaisuu = 0
				} else {
					result = result + xArray[i]
					nowKaisuu = nowKaisuu + 1
				}
			}
		} else {
			result = result + xArray[i]
			nowKaisuu = nowKaisuu + 1
		}

		y, ok := cyMap[nowKaisuu]
		if ok {
			result = result + y
		}

	}

	fmt.Println(result)
}

func DFlippingandBonusRdr(rdr *bufio.Reader) string {
	buf := make([]byte, 0, 10000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
