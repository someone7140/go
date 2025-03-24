package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CUniquenessMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrs := strings.Split(CUniquenessRdr(rdr), " ")
	aArray := make([]int, n)
	aMap := map[int][]int{}
	for i, aMoji := range aStrs {
		a, _ := strconv.Atoi(aMoji)
		aArray[i] = a
		val, ok := aMap[a]
		if !ok {
			aMap[a] = []int{i}
		} else {
			aMap[a] = append(val, i)
		}
	}

	result := -1
	maxNumber := -1
	for key, value := range aMap {
		if len(value) == 1 {
			if maxNumber < key {
				result = value[0] + 1
				maxNumber = key
			}
		}
	}

	fmt.Println(result)

}

func CUniquenessRdr(rdr *bufio.Reader) string {
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
