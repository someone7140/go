package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PStruct struct {
	value int
	count int
	index int
}

func DDrawYourCardsMain() {
	var n, k int
	fmt.Scan(&n, &k)

	resultSlice := make([]string, n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	p := DDrawYourCardsRdr(rdr)
	pArray := make([]int, n)
	pStrArray := strings.Split(p, " ")
	for i, pMoji := range pStrArray {
		pNum, _ := strconv.Atoi(pMoji)
		pArray[i] = pNum
	}

	if k == 1 {
		for i, _ := range pArray {
			resultSlice[i] = strconv.FormatInt(int64(i+1), 10)
		}
	} else {
		tempLen := 0
		var tempSlice []PStruct
		for i, pValue := range pArray {
			if tempLen == 0 {
				tempSlice = append(tempSlice, PStruct{
					value: pValue,
					count: 1,
					index: i,
				})
				resultSlice[i] = "-1"
				tempLen = 1
			} else {
				max := tempSlice[tempLen-1]
				if max.value < pValue {
					tempSlice = append(tempSlice, PStruct{
						value: pValue,
						count: 1,
						index: i,
					})
					resultSlice[i] = "-1"
					tempLen = tempLen + 1
				} else {
					min := tempSlice[0]
					if min.value > pValue {
						if min.count+1 == k {
							if tempLen == 1 {
								tempSlice = []PStruct{}
							} else {
								resultSlice[min.index] = strconv.FormatInt(int64(i+1), 10)
								resultSlice[i] = strconv.FormatInt(int64(i+1), 10)
							}
						} else {
							resultSlice[min.index] = strconv.FormatInt(int64(i+1), 10)
						}
					}
				}
			}
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func DDrawYourCardsRdr(rdr *bufio.Reader) string {
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
