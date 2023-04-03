package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CSocksMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CSocksReadLine(rdr), " ")
	aMap := map[int]int{}

	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		v, ok := aMap[a]
		if ok {
			aMap[a] = v + 1
		} else {
			aMap[a] = 1
		}
	}

	result := 0
	for _, v := range aMap {
		result = result + (v / 2)
	}
	fmt.Println(result)

}

func CSocksReadLine(rdr *bufio.Reader) string {
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
