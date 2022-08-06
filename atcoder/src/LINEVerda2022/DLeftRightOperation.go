package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var n, l, r int
	fmt.Scan(&n, &l, &r)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aLine := DLeftRightOperationRdr(rdr)

	aArray := make([]int64, n)
	aStrArray := strings.Split(aLine, " ")
	// 通常の累積和
	aSumArray := make([]int64, n)
	var aSum int64
	for i, aMoji := range aStrArray {
		aInt, _ := strconv.Atoi(aMoji)
		aArray[i] = int64(aInt)
		aSum = aSum + int64(aInt)
		aSumArray[i] = aSum
	}
	// 逆からの累積和
	aSumReverseArray := make([]int64, n)
	var tempASumReverse int64
	for i := n - 1; i >= 0; i-- {
		tempASumReverse = tempASumReverse + aArray[i]
		aSumReverseArray[i] = tempASumReverse
	}

	// Lの累積和
	var lSumMin int64
	lSumMin = 999999999999999999
	lSumIndex := -1
	tempLSum := aSum
	for i := 0; i < n; i++ {
		tempLSum = tempLSum + (int64(l) - aArray[i])
		if tempLSum < lSumMin {
			lSumMin = tempLSum
			if lSumMin < aSumArray[i] {
				lSumIndex = i
			}
		}
	}
	// Rの累積和
	var rSumMin int64
	rSumMin = 999999999999999999
	rSumIndex := -1
	tempRSum := aSum
	for i := n - 1; i >= 0; i-- {
		tempRSum = tempRSum + (int64(r) - aArray[i])
		if tempRSum < rSumMin {
			rSumMin = tempRSum
			if rSumMin < aSumReverseArray[i] {
				rSumIndex = i
			}
		}
	}

	if lSumIndex == -1 && rSumIndex == -1 {
		fmt.Println(aSum)
	} else if lSumIndex == -1 {
		fmt.Println(rSumMin)
	} else if rSumIndex == -1 {
		fmt.Println(lSumMin)
	} else {
		resultArray := make([]int64, n)
		if l > r {
			for i := 0; i <= lSumIndex; i++ {
				resultArray[i] = int64(l)
			}
			for i := lSumIndex + 1; i < n; i++ {
				resultArray[i] = aArray[i]
			}
			for i := n - 1; i >= rSumIndex; i-- {
				resultArray[i] = int64(r)
			}
		} else {
			for i := n - 1; i >= rSumIndex; i-- {
				resultArray[i] = int64(r)
			}
			for i := rSumIndex - 1; i >= 0; i-- {
				resultArray[i] = aArray[i]
			}
			for i := 0; i <= lSumIndex; i++ {
				resultArray[i] = int64(l)
			}
		}

		var sum int64
		for _, x := range resultArray {
			sum += x
		}
		fmt.Println(sum)
	}

}

func DLeftRightOperationRdr(rdr *bufio.Reader) string {
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
