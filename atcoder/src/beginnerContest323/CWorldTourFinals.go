package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CWorldTourFinalsMondai struct {
	index int
	point int
}

func CWorldTourFinalsMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	// Aの配列を読み込み
	aArrayInput := make([]int, m)
	aArray := make([]CWorldTourFinalsMondai, m)
	aArrayStrs := strings.Split(CWorldTourFinalsReadLine(rdr), " ")
	for i, aStr := range aArrayStrs {
		a, _ := strconv.Atoi(aStr)
		aArray[i] = CWorldTourFinalsMondai{
			index: i,
			point: a,
		}
		aArrayInput[i] = a
	}
	sort.Slice(aArray, func(i, j int) bool {
		return aArray[i].point > aArray[j].point
	})

	// sの配列を読み込み
	sArrayArray := make([][]bool, n)
	sPointArray := make([]int, n)
	maxPoint := -11
	for i := 0; i < n; i++ {
		s := CWorldTourFinalsReadLine(rdr)

		sArray := make([]bool, m)
		point := i + 1
		for j, c := range s {
			sMoji := string([]rune{c})
			if sMoji == "o" {
				sArray[j] = true
				point = point + aArrayInput[j]
			} else {
				sArray[j] = false
			}
		}
		sPointArray[i] = point
		sArrayArray[i] = sArray
		if maxPoint < point {
			maxPoint = point
		}
	}

	var resultSlice []string
	for i := 0; i < n; i++ {
		point := sPointArray[i]
		sArray := sArrayArray[i]
		tempResult := 0
		if point < maxPoint {
			sabun := maxPoint - point
			for _, a := range aArray {
				if !sArray[a.index] {
					tempResult = tempResult + 1
					sabun = sabun - a.point
					if sabun < 0 {
						break
					}
				}
			}
		}
		resultSlice = append(resultSlice, strconv.FormatInt(int64(tempResult), 10))
	}

	fmt.Println(strings.Join(resultSlice, "\n"))

}

func CWorldTourFinalsReadLine(rdr *bufio.Reader) string {
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
