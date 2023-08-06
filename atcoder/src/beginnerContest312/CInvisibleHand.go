package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CInvisibleHandMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aArray := make([]int, n)
	for i, aMoji := range strings.Split(CInvisibleHandReadLine(rdr), " ") {
		a, _ := strconv.Atoi(aMoji)
		aArray[i] = a
	}
	sort.Sort(sort.Reverse(sort.IntSlice(aArray)))

	bArray := make([]int, m)
	for i, bMoji := range strings.Split(CInvisibleHandReadLine(rdr), " ") {
		b, _ := strconv.Atoi(bMoji)
		bArray[i] = b
	}
	sort.Sort(sort.Reverse(sort.IntSlice(bArray)))

	result := bArray[m-1] + 1
	nPos := 0
	mPos := 0
	nPosSaled := false
	// 買い手でループ
	for {
		if mPos == m {
			if mPos-1 != nPos {
				result = bArray[0] + 1
			}
			break
		}
		if nPos == n {
			if mPos != nPos-1 {
				result = bArray[0] + 1
			}
			break
		}
		b := bArray[mPos]
		saleAmount := aArray[nPos]
		if b >= saleAmount {
			if !nPosSaled {
				nPosSaled = true
				mPos = mPos + 1
			} else {
				nPos = nPos + 1
				nPosSaled = false
			}
			result = saleAmount
		} else {
			if mPos == nPos {
				result = bArray[0] + 1
			}
			break
		}
	}
	fmt.Println(result)

}

func CInvisibleHandReadLine(rdr *bufio.Reader) string {
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
