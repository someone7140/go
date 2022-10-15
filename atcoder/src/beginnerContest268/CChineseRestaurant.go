package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func CChineseRestaurantMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	pLine := CChineseRestaurantRdr(rdr)
	pStrArray := strings.Split(pLine, " ")
	pArray := make([]int, n)

	for i := 0; i < n; i++ {
		pArray[i], _ = strconv.Atoi(pStrArray[i])
	}

	max := 0

	for i := 0; i < n; i++ {
		nowIndex := i
		temp := 0
		for j := 0; j < n; j++ {
			if nowIndex == n {
				nowIndex = 0
			}
			nokori := n - j
			if (temp + nokori) < max {
				break
			}
			sabun := pArray[nowIndex] - j
			if math.Abs(float64(sabun)) <= 1 || math.Abs(float64(sabun)) == float64(n-1) {
				temp = temp + 1
			}
			nowIndex = nowIndex + 1
		}
		if temp > max {
			max = temp
		}
		if max == n {
			break
		}
	}

	fmt.Println(max)

}

func CChineseRestaurantRdr(rdr *bufio.Reader) string {
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
