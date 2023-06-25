package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DPoisonousFullCourseMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	dpArray := make([][]int64, n)

	for i := 0; i < n; i++ {
		xy := strings.Split(DPoisonousFullCourseReadLine(rdr), " ")
		pointInt, _ := strconv.Atoi(xy[1])
		point := int64(pointInt)
		tempResult := make([]int64, 2)
		if i == 0 {
			if xy[0] == "0" {
				if point > 0 {
					tempResult[0] = point
				} else {
					tempResult[0] = 0
				}
				tempResult[1] = -999999999999999
			} else {
				tempResult[0] = 0
				tempResult[1] = point
			}
		} else {
			beforeInfo := dpArray[i-1]
			if xy[0] == "0" {
				if point > 0 {
					if beforeInfo[0] > beforeInfo[1] {
						tempResult[0] = beforeInfo[0] + point
					} else {
						tempPoint := beforeInfo[1] + point
						if tempPoint > beforeInfo[0] {
							tempResult[0] = tempPoint
						}
					}
				} else {
					tempPoint := beforeInfo[1] + point
					if tempPoint > beforeInfo[0] {
						tempResult[0] = tempPoint
					} else {
						tempResult[0] = beforeInfo[0]
					}
				}
				tempResult[1] = beforeInfo[1]
			} else {
				if point > 0 {
					tempResult[0] = beforeInfo[0]
					tempPoint := beforeInfo[0] + point
					if tempPoint > beforeInfo[1] {
						tempResult[1] = tempPoint
					} else {
						tempResult[1] = beforeInfo[1]
					}
				} else {
					tempResult[0] = beforeInfo[0]
					tempResult[1] = beforeInfo[1]
				}
			}
		}
		dpArray[i] = tempResult
	}

	dokuNashi := dpArray[n-1][0]
	doku := dpArray[n-1][1]
	if dokuNashi < doku {
		fmt.Println(doku)
	} else {
		fmt.Println(dokuNashi)
	}
}

func DPoisonousFullCourseReadLine(rdr *bufio.Reader) string {
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
