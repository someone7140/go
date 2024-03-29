package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DIndexA struct {
	index int
	value int64
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(DIndexARdr(rdr), " ")
	var aSlice = make([]DIndexA, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aInt64 := int64(a)
		aSlice[i] = DIndexA{
			index: i,
			value: aInt64,
		}
	}

}

func DIndexARdr(rdr *bufio.Reader) string {
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
