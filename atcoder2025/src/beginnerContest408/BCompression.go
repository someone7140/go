package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BCompressionMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	cStrList := strings.Split(BCompressionRdr(rdr), " ")

	cSet := make(map[int]struct{})
	for _, cStr := range cStrList {
		c, _ := strconv.Atoi(cStr)
		cSet[c] = struct{}{}
	}

	var resultList []int
	for k, _ := range cSet {
		resultList = append(resultList, k)
	}
	sort.Ints(resultList)

	var resultStrList []string
	for _, result := range resultList {
		resultStrList = append(resultStrList, strconv.FormatInt(int64(result), 10))
	}

	fmt.Println(strconv.FormatInt(int64(len(resultStrList)), 10))
	fmt.Println(strings.Join(resultStrList, " "))
}

func BCompressionRdr(rdr *bufio.Reader) string {
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
