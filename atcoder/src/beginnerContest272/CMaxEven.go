package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CMaxEvenMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aLine := CMaxEvenRdr(rdr)
	aStrArray := strings.Split(aLine, " ")

	var guusuuSlice []int
	var kisuuSlice []int
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		if a%2 == 0 {
			guusuuSlice = append(guusuuSlice, a)
		} else {
			kisuuSlice = append(kisuuSlice, a)
		}
	}
	sort.Ints(guusuuSlice)
	sort.Ints(kisuuSlice)

	guusuuLen := len(guusuuSlice)
	kisuuLen := len(kisuuSlice)
	guusuuSumMax := -1
	kisuuSumMax := -1

	if guusuuLen > 1 {
		guusuuSumMax = guusuuSlice[guusuuLen-1] + guusuuSlice[guusuuLen-2]
	}
	if kisuuLen > 1 {
		kisuuSumMax = kisuuSlice[kisuuLen-1] + kisuuSlice[kisuuLen-2]
	}

	result := -1

	if guusuuSumMax > kisuuSumMax {
		result = guusuuSumMax
	} else {
		result = kisuuSumMax
	}

	fmt.Println(result)
}

func CMaxEvenRdr(rdr *bufio.Reader) string {
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
