package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BColoringMatrixMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aArrayArray1 := make([][]int, n)
	for i := 0; i < n; i++ {
		aStrArray := strings.Split(BColoringMatrixReadLine(rdr), " ")
		aArray := make([]int, n)
		for j := 0; j < n; j++ {
			a, _ := strconv.Atoi(aStrArray[j])
			aArray[j] = a
		}
		aArrayArray1[i] = aArray
	}

	bArrayArray := make([][]int, n)
	for i := 0; i < n; i++ {
		bStrArray := strings.Split(BColoringMatrixReadLine(rdr), " ")
		bArray := make([]int, n)
		for j := 0; j < n; j++ {
			b, _ := strconv.Atoi(bStrArray[j])
			bArray[j] = b
		}
		bArrayArray[i] = bArray
	}

	var judgeFunc func(a [][]int, b [][]int) bool
	judgeFunc = func(a [][]int, b [][]int) bool {
		result := true
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				aValue := a[i][j]
				if aValue == 1 {
					if b[i][j] != 1 {
						result = false
						break
					}
				}
			}
			if !result {
				break
			}
		}
		return result
	}

	var rollingFunc func(arrayArray [][]int) [][]int
	rollingFunc = func(arrayArray [][]int) [][]int {
		newArrayArray := make([][]int, n)
		for i := 0; i < n; i++ {
			newArray := make([]int, n)
			for j := 0; j < n; j++ {
				newArray[j] = arrayArray[n-1-j][i]
			}
			newArrayArray[i] = newArray
		}
		return newArrayArray
	}

	targetArrayArray := aArrayArray1
	if judgeFunc(targetArrayArray, bArrayArray) {
		fmt.Println("Yes")
	} else {
		count := 0
		yesFlag := false
		for {
			targetArrayArray = rollingFunc(targetArrayArray)
			if judgeFunc(targetArrayArray, bArrayArray) {
				yesFlag = true
				break
			}
			count = count + 1
			if count > 50 {
				break
			}
		}
		if yesFlag {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func BColoringMatrixReadLine(rdr *bufio.Reader) string {
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
