package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func C2x2PlacingMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	rcSet := make(map[string]struct{})
	for i := 0; i < m; i++ {
		aStrArray := strings.Split(C2x2PlacingRdr(rdr), " ")
		r, _ := strconv.Atoi(aStrArray[0])
		c, _ := strconv.Atoi(aStrArray[1])
		r = r - 1
		c = c - 1
		rc1 := strconv.FormatInt(int64(r), 10) + "-" + strconv.FormatInt(int64(c), 10)
		rc2 := strconv.FormatInt(int64(r+1), 10) + "-" + strconv.FormatInt(int64(c), 10)
		rc3 := strconv.FormatInt(int64(r), 10) + "-" + strconv.FormatInt(int64(c+1), 10)
		rc4 := strconv.FormatInt(int64(r+1), 10) + "-" + strconv.FormatInt(int64(c+1), 10)
		_, ok1 := rcSet[rc1]
		_, ok2 := rcSet[rc2]
		_, ok3 := rcSet[rc3]
		_, ok4 := rcSet[rc4]
		if !ok1 && !ok2 && !ok3 && !ok4 {
			rcSet[rc1] = struct{}{}
			rcSet[rc2] = struct{}{}
			rcSet[rc3] = struct{}{}
			rcSet[rc4] = struct{}{}
		}
	}

	fmt.Println(len(rcSet) / 4)
}

func C2x2PlacingRdr(rdr *bufio.Reader) string {
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
