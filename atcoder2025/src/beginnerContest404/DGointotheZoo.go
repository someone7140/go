package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DGointotheZooMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	cStrList := strings.Split(DGointotheZooRdr(rdr), " ")
	cList := make([]int64, n)

	for i := 0; i < n; i++ {
		c, _ := strconv.Atoi(cStrList[i])
		cList[i] = int64(c)
	}
	enSetMap := map[int]map[int]struct{}{}

	result := int64(-1)
	for i := 0; i < m; i++ {
		ka := strings.Split(DGointotheZooRdr(rdr), " ")
		k, _ := strconv.Atoi(ka[0])
		for j := 1; j <= k; j++ {
			a, _ := strconv.Atoi(ka[j])
			enSet, ok := enSetMap[a]
			if ok {
				enSet[i] = struct{}{}
				enSetMap[a] = enSet
			} else {
				newSet := make(map[int]struct{})
				newSet[i] = struct{}{}
				enSetMap[a] = newSet
			}
		}
	}

	fmt.Println(result)
}

func DGointotheZooRdr(rdr *bufio.Reader) string {
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
