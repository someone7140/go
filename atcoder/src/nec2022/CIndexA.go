package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CIndexAMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CIndexARdr(rdr), " ")
	var aSlice = make([]int64, n)
	var max int64
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aInt64 := int64(a)
		aSlice[i] = aInt64
		if aInt64 > max {
			max = aInt64
		}
	}

	if m == 1 {
		fmt.Println(max)
	} else {
		max = 0
		var sumKakeMae int64
		var tempSlice = make([]int64, m)
		for i := 0; i < m; i++ {
			tempSlice[i] = aSlice[i]
			max = max + aSlice[i]*int64(i+1)
			sumKakeMae = sumKakeMae + aSlice[i]
		}
		tempMax := max
		for i := m; i < n; i++ {
			next := aSlice[i]
			tempMax = tempMax - sumKakeMae + next*int64(m)
			if tempMax > max {
				max = tempMax
			}
			sumKakeMae = sumKakeMae + next - aSlice[i-m]
		}
		fmt.Println(max)
	}

}

func CIndexARdr(rdr *bufio.Reader) string {
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
