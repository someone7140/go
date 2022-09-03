package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DIrohaandHaikuMain() {
	var nInt64, p, q, r int64
	fmt.Scan(&nInt64, &p, &q, &r)

	n := int(nInt64)
	aArray := make([]int64, n)
	aSumArray := make([]int64, n)
	var aSum int64

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aLine := DIrohaandHaikuRdr(rdr)
	aStrArray := strings.Split(aLine, " ")

	for i, aStr := range aStrArray {
		aInt64, _ := strconv.ParseInt(aStr, 10, 64)
		aSum = aSum + aInt64

		aArray[i] = aInt64
		aSumArray[i] = aSum
	}

	result := "No"
	var tempPSum int64
	for i := 0; i < n; i++ {
		if result == "Yes" {
			break
		}
		tempPSum = aSumArray[i]
		if tempPSum >= p {
			start := i
			if tempPSum > p {
				if result == "Yes" {
					break
				}
				for s := 0; s < i; s++ {
					tempPSum = aSumArray[i] - aSumArray[s]
					if tempPSum == p || tempPSum < p {
						break
					}
				}
			}
			if tempPSum == p {
				minusSum := aSumArray[start]
				for j := start + 1; j < n; j++ {
					if result == "Yes" {
						break
					}
					tempQSum := aSumArray[j] - minusSum
					if tempQSum == q {
						minusSum = aSumArray[j]
						for k := j + 1; k < n; k++ {
							tempRSum := aSumArray[k] - minusSum
							if tempRSum == r {
								result = "Yes"
								break
							}
						}
					}
				}
			}
		}

	}
	fmt.Println(result)
}

func DIrohaandHaikuRdr(rdr *bufio.Reader) string {
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
