package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CStandingsStruct struct {
	key int
	a   int64
	sum int64
}

func CStandingsMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var abSlice = make([]CStandingsStruct, n)
	for i := 0; i < n; i++ {
		abList := strings.Split(CStandingsReadLine(rdr), " ")
		aInt, _ := strconv.Atoi(abList[0])
		bInt, _ := strconv.Atoi(abList[1])
		a := int64(aInt)
		b := int64(bInt)
		sum := a + b
		abSlice[i] = CStandingsStruct{
			key: i + 1,
			sum: sum,
			a:   a,
		}
	}

	sort.Slice(abSlice, func(i, j int) bool {
		i1 := abSlice[i].a
		i2 := abSlice[i].sum
		j1 := abSlice[j].a
		j2 := abSlice[j].sum

		// Compute ad-bc
		comp := i1*j2 - i2*j1
		if comp == 0 {
			return abSlice[i].key < abSlice[j].key
		} else {
			return comp > 0
		}
	})

	var resultSlice = make([]string, n)
	for i := 0; i < n; i++ {
		resultSlice[i] = strconv.FormatInt(int64(abSlice[i].key), 10)
	}
	fmt.Println(strings.Join(resultSlice, " "))

}

func CStandingsReadLine(rdr *bufio.Reader) string {
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
