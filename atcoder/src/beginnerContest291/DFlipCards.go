package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DFlipCardsMain() {
	var n int
	fmt.Scan(&n)

	var mapSlice = make([]map[int]int64, n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < n; i++ {
		abLine := strings.Split(DFlipCardsReadLine(rdr), " ")
		a, _ := strconv.Atoi(abLine[0])
		b, _ := strconv.Atoi(abLine[1])

		abMap := map[int]int64{}
		if i == 0 {
			if a == b {
				abMap[a] = int64(2)
			} else {
				abMap[a] = int64(1)
				abMap[b] = int64(1)
			}
			mapSlice[i] = abMap
		} else {
			prevMap := mapSlice[i-1]
			for k, v := range prevMap {
				if k != a {
					v2, ok := abMap[a]
					if !ok {
						abMap[a] = v
					} else {
						abMap[a] = (v2 + v) % 998244353
					}
				}
				if k != b {
					v2, ok := abMap[b]
					if !ok {
						abMap[b] = v
					} else {
						abMap[b] = (v2 + v) % 998244353
					}
				}
			}
		}
		mapSlice[i] = abMap
	}

	result := int64(0)
	resultMap := mapSlice[n-1]
	for _, v := range resultMap {
		result = (result + v) % 998244353
	}
	fmt.Println(result)

}

func DFlipCardsReadLine(rdr *bufio.Reader) string {
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
