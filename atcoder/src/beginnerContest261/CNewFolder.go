package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CNewFolderMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	resultArray := make([]string, n)
	resultMap := map[string]int{}

	for i := 0; i < n; i++ {
		s := CNewFolderRdr(rdr)
		v, ok := resultMap[s]
		if ok {
			resultArray[i] = s + "(" + strconv.FormatInt(int64(v), 10) + ")"
			resultMap[s] = v + 1
		} else {
			resultArray[i] = s
			resultMap[s] = 1
		}

	}

	fmt.Println(strings.Join(resultArray, "\n"))
}

func CNewFolderRdr(rdr *bufio.Reader) string {
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
