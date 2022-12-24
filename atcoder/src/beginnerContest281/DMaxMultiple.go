package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DMaxMultipleResult int64
var DMaxMultipleAArray []int64
var DMaxMultipleAArrayArraySet [][]int64
var DMaxMultipleAArrayN, DMaxMultipleAArrayK, DMaxMultipleAArrayD int

func main() {
	fmt.Scan(&DMaxMultipleAArrayN, &DMaxMultipleAArrayK, &DMaxMultipleAArrayD)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(DMaxMultipleReadLine(rdr), " ")
	DMaxMultipleAArray = make([]int64, DMaxMultipleAArrayN)

	for i := 0; i <= DMaxMultipleAArrayK; i++ {
		dMaxMultipleAArraySet := make([]int64, DMaxMultipleAArrayN)
		DMaxMultipleAArrayArraySet = append(DMaxMultipleAArrayArraySet, dMaxMultipleAArraySet)
	}

	for i := 0; i < DMaxMultipleAArrayN; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		DMaxMultipleAArray[i] = int64(a)
	}

	DMaxMultipleLoop(0, -1, 1)
	DMaxMultipleResult = -1
	for i := 0; i < DMaxMultipleAArrayN; i++ {
		temp := DMaxMultipleAArrayArraySet[DMaxMultipleAArrayK][i]
		if temp != int64(0) && DMaxMultipleResult < temp {
			DMaxMultipleResult = temp
		}
	}

	fmt.Println(DMaxMultipleResult)
}

func DMaxMultipleLoop(index int, prevIndex int, count int) {
	if index < DMaxMultipleAArrayN {
		if count == 1 {
			DMaxMultipleAArrayArraySet[1][index] = DMaxMultipleAArray[index]
			DMaxMultipleLoop(index+1, -1, 1)
			DMaxMultipleLoop(index+1, index, 2)
		} else if count == (DMaxMultipleAArrayK) {
			temp := DMaxMultipleAArrayArraySet[count][prevIndex] + DMaxMultipleAArray[index]
			if DMaxMultipleAArrayArraySet[count][index] < temp && temp%int64(DMaxMultipleAArrayD) == 0 {
				DMaxMultipleAArrayArraySet[count][index] = temp
			}
		} else {
			temp := DMaxMultipleAArrayArraySet[count][prevIndex] + DMaxMultipleAArray[index]
			if DMaxMultipleAArrayArraySet[count][index] < temp {
				DMaxMultipleAArrayArraySet[count][index] = temp
			}
			DMaxMultipleLoop(index+1, prevIndex, count)
			DMaxMultipleLoop(index+1, index, count+1)
		}
	}
}

func DMaxMultipleReadLine(rdr *bufio.Reader) string {
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
