package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CSigmaMain() {
	var n, k int
	fmt.Scan(&n, &k)

	result := (int64(k) + 1) * int64(k) / 2

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CSigmaReadLine(rdr), " ")
	aSet := make(map[int64]struct{})
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		a64 := int64(a)
		_, ok := aSet[a64]
		if !ok {
			if a64 <= int64(k) {
				result = result - a64
			}
			aSet[a64] = struct{}{}
		}
	}

	fmt.Println(result)
}

func CSigmaReadLine(rdr *bufio.Reader) string {
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
