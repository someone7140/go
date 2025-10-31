package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CDistanceIndicatorsMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CDistanceIndicatorsRdr(rdr), " ")

	aArray := make([]int, n)
	iCountMap := map[int]int{}
	jCountMap := map[int]int{}
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray[i] = a

		ai := a + i
		count, ok := iCountMap[ai]
		if ok {
			iCountMap[ai] = count + 1
		} else {
			iCountMap[ai] = 1
		}

		aj := i - a
		count, ok = jCountMap[aj]
		if ok {
			jCountMap[aj] = count + 1
		} else {
			jCountMap[aj] = 1
		}
	}

	result := 0
	for key, value := range iCountMap {
		jCount, ok := jCountMap[key]
		if ok {
			result = result + value*jCount
		}
	}

	fmt.Println(result)
}

func CDistanceIndicatorsRdr(rdr *bufio.Reader) string {
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
