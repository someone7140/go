package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CABCMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var aArray = make([]int, n)
	aStrArray := strings.Split(CABCReadLine(rdr), " ")
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray[i] = a
	}

	m, _ := strconv.Atoi(CABCReadLine(rdr))
	var bArray = make([]int, m)
	bStrArray := strings.Split(CABCReadLine(rdr), " ")
	for i := 0; i < m; i++ {
		b, _ := strconv.Atoi(bStrArray[i])
		bArray[i] = b
	}

	l, _ := strconv.Atoi(CABCReadLine(rdr))
	var cArray = make([]int, l)
	cStrArray := strings.Split(CABCReadLine(rdr), " ")
	for i := 0; i < l; i++ {
		c, _ := strconv.Atoi(cStrArray[i])
		cArray[i] = c
	}

	q, _ := strconv.Atoi(CABCReadLine(rdr))
	var xArray = make([]int, q)
	xStrArray := strings.Split(CABCReadLine(rdr), " ")
	for i := 0; i < q; i++ {
		x, _ := strconv.Atoi(xStrArray[i])
		xArray[i] = x
	}

	sumSet := make(map[int]struct{})

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for o := 0; o < l; o++ {
				sumSet[aArray[i]+bArray[j]+cArray[o]] = struct{}{}
			}
		}
	}

	var resultSlice []string
	for i := 0; i < q; i++ {
		_, ok := sumSet[xArray[i]]
		if ok {
			resultSlice = append(resultSlice, "Yes")
		} else {
			resultSlice = append(resultSlice, "No")
		}
	}
	fmt.Println(strings.Join(resultSlice, "\n"))
}

func CABCReadLine(rdr *bufio.Reader) string {
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
