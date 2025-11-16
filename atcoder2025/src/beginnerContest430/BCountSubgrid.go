package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BCountSubgridMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	hwArrayArray := make([][]string, n)

	for i := 0; i < n; i++ {
		sStrArray := strings.Split(BCountSubgridRdr(rdr), "")
		hwArrayArray[i] = sStrArray
	}
	sSet := make(map[string]struct{})

	for i := 0; i <= n-m; i++ {
		for j := 0; j <= n-m; j++ {
			tempS := ""
			for h := i; h < i+m; h++ {
				for k := j; k < j+m; k++ {
					tempS = tempS + hwArrayArray[h][k]
				}
			}
			sSet[tempS] = struct{}{}
		}
	}
	fmt.Println(len(sSet))
}

func BCountSubgridRdr(rdr *bufio.Reader) string {
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
