package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DStringRotationMain() {
	var t int
	fmt.Scan(&t)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var resultSlice []string

	for i := 0; i < t; i++ {
		nStr := DStringRotationRdr(rdr)
		n, _ := strconv.Atoi(nStr)
		sStrs := strings.Split(DStringRotationRdr(rdr), "")
		if n == 1 {
			resultSlice = append(resultSlice, sStrs[0])
		} else {
			tempResult := ""
			lDeleteTarget := ""
			irekaeDone := false
			for j := 0; j < n; j++ {
				if irekaeDone {
					tempResult = tempResult + sStrs[j]
					continue
				}
				if lDeleteTarget != "" {
					if lDeleteTarget < sStrs[j] {
						tempResult = tempResult + lDeleteTarget + sStrs[j]
						lDeleteTarget = ""
						irekaeDone = true
					} else {
						tempResult = tempResult + sStrs[j]
					}
					continue
				}
				if j < n-1 && sStrs[j] > sStrs[j+1] {
					lDeleteTarget = sStrs[j]
				} else {
					tempResult = tempResult + sStrs[j]
				}
			}
			if lDeleteTarget != "" {
				tempResult = tempResult + lDeleteTarget
			}

			resultSlice = append(resultSlice, tempResult)
		}
	}
	fmt.Println(strings.Join(resultSlice, "\n"))
}

func DStringRotationRdr(rdr *bufio.Reader) string {
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
