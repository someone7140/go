package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BNoDivisibleRangeMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(BNoDivisibleRangeRdr(rdr), " ")
	aArray := make([]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		sum = sum + a
		aArray[i] = a
	}

	result := 0
	for i := 0; i < n; i++ {
		aTempSum := aArray[i]
		var aTempArray []int
		aTempArray = append(aTempArray, aArray[i])
		for j := i + 1; j < n; j++ {
			aTempArray = append(aTempArray, aArray[j])
			aTempSum = aTempSum + aArray[j]
			tempResult := true
			for _, aTemp := range aTempArray {
				if aTempSum%aTemp == 0 {
					tempResult = false
					break
				}
			}
			if tempResult {
				result = result + 1
			}
		}
	}

	fmt.Println(result)
}

func BNoDivisibleRangeRdr(rdr *bufio.Reader) string {
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
