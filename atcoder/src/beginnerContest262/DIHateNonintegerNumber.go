package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aLine := DIHateNonintegerNumberRdr(rdr)
	aArray := make([]int64, n)
	aStrArray := strings.Split(aLine, " ")
	for i, aMoji := range aStrArray {
		pNum, _ := strconv.Atoi(aMoji)
		aArray[i] = int64(pNum)
	}

	var result int64
	var aSliceSlice = make([][][]int64, n)
	for i := 0; i < n; i++ {
		a := aArray[i]
		result = (result + 1) % 998244353
		if i == 0 {
			aSlice1 := []int64{a}
			aSliceSlice[i] = [][]int64{aSlice1}
		} else {
			afterASlice := make([][]int64, i+1)
			beforeASlice := aSliceSlice[i-1]
			for j := 0; j <= i; j++ {
				if j == 0 {
					afterASlice[j] = append(beforeASlice[j], a)
				} else {
					var afterASlice2 []int64
					beforeASlice2 := beforeASlice[j-1]
					lenBeforeASlice2 := len(beforeASlice2)
					for k := 0; k < lenBeforeASlice2; k++ {
						beforeA := beforeASlice2[k]
						plus := beforeA + a
						afterASlice2 = append(afterASlice2, plus)
						amari := plus % int64(j+1)
						if amari == 0 {
							result = (result + 1) % 998244353
						}
					}
					if i > j {
						afterASlice2 = append(afterASlice2, beforeASlice[j]...)
					}
					afterASlice[j] = afterASlice2
				}
			}
			aSliceSlice[i] = afterASlice
		}
	}
	fmt.Println(result)

}

func DIHateNonintegerNumberRdr(rdr *bufio.Reader) string {
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
