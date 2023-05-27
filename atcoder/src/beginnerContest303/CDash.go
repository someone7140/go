package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CDashMain() {
	var n, m, h, k int
	fmt.Scan(&n, &m, &h, &k)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := CDashReadLine(rdr)
	mMap := map[string]bool{}

	for i := 0; i < m; i++ {
		mArray := strings.Split(CDashReadLine(rdr), " ")
		x, _ := strconv.Atoi(mArray[0])
		y, _ := strconv.Atoi(mArray[1])
		key := strconv.FormatInt(int64(x), 10) + "-" + strconv.FormatInt(int64(y), 10)
		mMap[key] = true
	}

	result := "Yes"
	x := 0
	y := 0

	for i, c := range s {
		houkou := string([]rune{c})
		h = h - 1
		if houkou == "L" {
			x = x - 1
		} else if houkou == "R" {
			x = x + 1
		} else if houkou == "U" {
			y = y + 1
		} else {
			y = y - 1
		}

		if h < k {
			key := strconv.FormatInt(int64(x), 10) + "-" + strconv.FormatInt(int64(y), 10)
			v, ok := mMap[key]
			if v && ok {
				h = k
				mMap[key] = false
			}
		}

		if h == 0 && i != n-1 {
			result = "No"
			break
		}
	}
	fmt.Println(result)

}

func CDashReadLine(rdr *bufio.Reader) string {
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
