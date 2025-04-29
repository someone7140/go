package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DLineCrossingMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var result int64
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	pointArrayArray := make([][]int, n)
	mArrayArray := make([][]int, m)
	for i := 0; i < n; i++ {
		pointArrayArray[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		ab := strings.Split(DLineCrossingRdr(rdr), " ")
		a, _ := strconv.Atoi(ab[0])
		b, _ := strconv.Atoi(ab[1])
		pointArrayArray[a-1][b-1] = 1
		mArrayArray[i] = []int{a - 1, b - 1}
	}

	for i := 0; i < m; i++ {
		a := mArrayArray[i][0]
		b := mArrayArray[i][1]

		// aがプラスでbがマイナス
		tempA1 := a
		tempB1 := b
		for {
			tempA1 = tempA1 + 1
			if tempA1 == n {
				tempA1 = 0
			}
			tempB1 = tempB1 - 1
			if tempB1 == -1 {
				tempB1 = n - 1
			}
			if tempA1 == tempB1 || tempA1 == a || tempB1 == b {
				break
			} else {
				if tempA1 > tempB1 {
					if pointArrayArray[tempB1][tempA1] == 1 {
						result = result + 1
					}
				} else {
					if tempA1 > tempB1 {
						if pointArrayArray[tempA1][tempB1] == 1 {
							result = result + 1
						}
					}
				}
			}
		}

		// bがプラスでaがマイナス
		tempA2 := a
		tempB2 := b
		for {
			tempA2 = tempA2 - 1
			if tempA2 == -1 {
				tempA2 = n - 1
			}
			tempB2 = tempB2 + 1
			if tempB2 == n {
				tempB2 = 0
			}
			if tempA2 == tempB2 || tempA1 == a || tempB1 == b {
				break
			} else {
				if tempA2 > tempB2 {
					if pointArrayArray[tempB2][tempA2] == 1 {
						result = result + 1
					}
				} else {
					if tempA2 > tempB2 {
						if pointArrayArray[tempA2][tempB2] == 1 {
							result = result + 1
						}
					}
				}
			}
		}
	}

	fmt.Println(result)
}

func DLineCrossingRdr(rdr *bufio.Reader) string {
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
