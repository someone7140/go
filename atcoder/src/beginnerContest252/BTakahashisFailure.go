package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BTakahashisFailureMain() {
	var n, k int
	fmt.Scan(&n, &k)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aLine := BTakahashisFailureRdr(rdr)
	aStrArray := strings.Split(aLine, " ")
	var aArray []int
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray = append(aArray, a)
	}
	bLine := BTakahashisFailureRdr(rdr)
	bStrArray := strings.Split(bLine, " ")
	var bArray []int
	for i := 0; i < k; i++ {
		b, _ := strconv.Atoi(bStrArray[i])
		bArray = append(bArray, b-1)
	}

	max := -9999
	result := "No"
	for i := 0; i < n; i++ {
		a := aArray[i]
		if max < a {
			result = "No"
			max = a
			for j := 0; j < k; j++ {
				if i == bArray[j] {
					result = "Yes"
				}
			}
		} else if max == a {
			if result == "No" {
				for j := 0; j < k; j++ {
					if i == bArray[j] {
						result = "Yes"
					}
				}
			}
		}
	}
	fmt.Println(result)
}

func BTakahashisFailureRdr(rdr *bufio.Reader) string {
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
