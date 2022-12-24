package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BFirstQueryProblemMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(BFirstQueryProblemReadLine(rdr), " ")
	aArray := make([]int, n)

	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray[i] = a
	}
	q, _ := strconv.Atoi(BFirstQueryProblemReadLine(rdr))

	var results []string
	for i := 0; i < q; i++ {
		queryStrArr := strings.Split(BFirstQueryProblemReadLine(rdr), " ")
		index, _ := strconv.Atoi(queryStrArr[1])
		if queryStrArr[0] == "2" {
			results = append(results, strconv.FormatInt(int64(aArray[index-1]), 10))
		} else {
			update, _ := strconv.Atoi(queryStrArr[2])
			aArray[index-1] = update
		}
	}

	fmt.Println(strings.Join(results, "\n"))
}

func BFirstQueryProblemReadLine(rdr *bufio.Reader) string {
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
