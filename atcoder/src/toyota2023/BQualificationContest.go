package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func BQualificationContestMain() {
	var n, k int
	fmt.Scan(&n, &k)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var results []string
	for i := 0; i < n; i++ {
		s := BQualificationContestReadLine(rdr)
		if i < k {
			results = append(results, s)
		}
	}
	sort.Strings(results)
	fmt.Println(strings.Join(results, "\n"))
}

func BQualificationContestReadLine(rdr *bufio.Reader) string {
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
