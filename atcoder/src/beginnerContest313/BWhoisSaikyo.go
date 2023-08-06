package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BWhoisSaikyoMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	baMap := map[int][]int{}

	for i := 0; i < m; i++ {
		abListStr := strings.Split(BWhoisSaikyoReadLine(rdr), " ")
		a, _ := strconv.Atoi(abListStr[0])
		b, _ := strconv.Atoi(abListStr[1])
		aList, ok := baMap[b]
		if !ok {
			baMap[b] = []int{a}
		} else {
			baMap[b] = append(aList, a)
		}
	}

	result := -90
	for i := 1; i <= n; i++ {
		_, ok := baMap[i]
		if !ok {
			if result == -90 {
				result = i
			} else {
				result = -1
				break
			}
		}
	}

	fmt.Println(result)

}

func BWhoisSaikyoReadLine(rdr *bufio.Reader) string {
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
