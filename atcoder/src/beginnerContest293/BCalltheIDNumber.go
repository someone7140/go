package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BCalltheIDNumberMain() {
	var n int
	fmt.Scan(&n)

	xCallArray := make([]int, n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	xStrArray := strings.Split(BCalltheIDNumberReadLine(rdr), " ")
	for i := 0; i < n; i++ {
		x, _ := strconv.Atoi(xStrArray[i])
		if xCallArray[i] != 1 {
			xCallArray[x-1] = 1
		}
	}

	count := 0
	var resultSlice []string
	for i := 0; i < n; i++ {
		if xCallArray[i] != 1 {
			count = count + 1
			resultSlice = append(resultSlice, strconv.FormatInt(int64(i+1), 10))
		}
	}
	fmt.Println(count)
	fmt.Println(strings.Join(resultSlice, " "))
}

func BCalltheIDNumberReadLine(rdr *bufio.Reader) string {
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
