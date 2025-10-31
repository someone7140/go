package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CRotateandSumQueryMain() {
	var n, q int
	fmt.Scan(&n, &q)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CRotateandSumQueryRdr(rdr), " ")
	aArray := make([]int, 2*n)
	aArraySum := make([]int, 2*n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray[i] = a
		if i == 0 {
			aArraySum[i] = a
		} else {
			aArraySum[i] = aArraySum[i-1] + a
		}
	}
	for i := 0; i < n; i++ {
		aArray[n+i] = aArray[i]
		aArraySum[n+i] = aArraySum[n+i-1] + aArray[i]
	}

	var resultSlice []string
	nowIndex := 0
	for i := 0; i < q; i++ {
		qStrArray := strings.Split(CRotateandSumQueryRdr(rdr), " ")
		if qStrArray[0] == "1" {
			proceed, _ := strconv.Atoi(qStrArray[1])
			nowIndex = nowIndex + proceed
			if nowIndex >= n {
				nowIndex = nowIndex - n
			}
		} else {
			start, _ := strconv.Atoi(qStrArray[1])
			end, _ := strconv.Atoi(qStrArray[2])
			sabunIndex := nowIndex + start - 1
			result := aArraySum[nowIndex+end-1]
			if sabunIndex > 0 {
				result = result - aArraySum[sabunIndex-1]
			}
			resultSlice = append(resultSlice, strconv.FormatInt(int64(result), 10))
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func CRotateandSumQueryRdr(rdr *bufio.Reader) string {
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
