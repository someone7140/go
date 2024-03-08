package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BAdjacencyMatrixMain() {
	var n int
	fmt.Scan(&n)

	var resultSlice []string
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < n; i++ {
		var tempResults []string
		aArray := strings.Split(BAdjacencyMatrixReadLine(rdr), " ")
		for i, a := range aArray {
			if a == "1" {
				tempResults = append(tempResults, strconv.FormatInt(int64(i+1), 10))
			}
		}
		resultSlice = append(resultSlice, strings.Join(tempResults, " "))
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func BAdjacencyMatrixReadLine(rdr *bufio.Reader) string {
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
