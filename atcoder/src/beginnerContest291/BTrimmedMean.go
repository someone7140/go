package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BTrimmedMeanMain() {
	var n int
	fmt.Scan(&n)

	var xArray []int
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	xStrArray := strings.Split(BTrimmedMeanReadLine(rdr), " ")

	for i := 0; i < 5*n; i++ {
		x, _ := strconv.Atoi(xStrArray[i])
		xArray = append(xArray, x)
	}
	sort.Ints(xArray)
	xArray2 := xArray[n : 4*n]
	sum := 0

	for _, x := range xArray2 {
		sum += x
	}

	fmt.Println(float64(sum) / float64(3*n))

}

func BTrimmedMeanReadLine(rdr *bufio.Reader) string {
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
