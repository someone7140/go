package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BCitationMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrs := strings.Split(BCitationRdr(rdr), " ")
	countMap := map[int]int{}
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrs[i])
		count, ok := countMap[a]
		if ok {
			countMap[a] = count + 1
		} else {
			countMap[a] = 1
		}
	}

	result := 0
	nowCount := n
	for i := 0; i <= n; i++ {
		if nowCount >= i {
			result = i
		}
		count, ok := countMap[i]
		if ok {
			nowCount = nowCount - count
		}

	}
	fmt.Println(result)
}

func BCitationRdr(rdr *bufio.Reader) string {
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
