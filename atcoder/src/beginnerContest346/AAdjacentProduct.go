package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AAdjacentProductMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var resultSlice []string
	aStrArray := strings.Split(AAdjacentProductReadLine(rdr), " ")
	for i := 0; i < n-1; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		a2, _ := strconv.Atoi(aStrArray[i+1])
		res := a * a2
		resultSlice = append(resultSlice, strconv.FormatInt(int64(res), 10))
	}

	fmt.Println(strings.Join(resultSlice, " "))
}

func AAdjacentProductReadLine(rdr *bufio.Reader) string {
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
