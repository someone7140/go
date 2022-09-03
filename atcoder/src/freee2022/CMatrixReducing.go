package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var CMatrixReducingResult string
var CMatrixReducingSuccessBRow int
var CMatrixReducingSuccessMaxBRow int
var CMatrixReducingLimitARowIndex int
var CMatrixReducingLimitAColumn int
var CMatrixReducingLimitBColumn int
var CMatrixReducingAArrayArray [][]int
var CMatrixReducingBArrayArray [][]int

func CMatrixReducingMain() {
	var h1, w1 int
	fmt.Scan(&h1, &w1)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	for i := 0; i < h1; i++ {
		aLine := CMatrixReducingRdr(rdr)
		aArray := make([]int, w1)
		aStrArray := strings.Split(aLine, " ")
		for j, aMoji := range aStrArray {
			pNum, _ := strconv.Atoi(aMoji)
			aArray[j] = pNum
		}
		CMatrixReducingAArrayArray = append(CMatrixReducingAArrayArray, aArray)
	}
	CMatrixReducingLimitARowIndex = h1
	CMatrixReducingLimitAColumn = w1

	var h2, w2 int
	h2w2Line := CMatrixReducingRdr(rdr)
	h2w2StrArray := strings.Split(h2w2Line, " ")
	h2, _ = strconv.Atoi(h2w2StrArray[0])
	w2, _ = strconv.Atoi(h2w2StrArray[1])
	CMatrixReducingSuccessMaxBRow = h2

	for i := 0; i < h2; i++ {
		bLine := CMatrixReducingRdr(rdr)
		bArray := make([]int, w2)
		bStrArray := strings.Split(bLine, " ")
		for j, bMoji := range bStrArray {
			pNum, _ := strconv.Atoi(bMoji)
			bArray[j] = pNum
		}
		CMatrixReducingBArrayArray = append(CMatrixReducingBArrayArray, bArray)
	}
	CMatrixReducingLimitBColumn = w2

	CMatrixReducingResult = "No"
	CMatrixReducingSuccessBRow = 0

	for i := 0; i < h1; i++ {
		judgeCMatrixReducing(i, []int{}, true)
		if CMatrixReducingResult == "Yes" {
			break
		}
	}

	fmt.Println(CMatrixReducingResult)
}

func judgeCMatrixReducing(aRowIndex int, deleteColumnIndexList []int, starFlag bool) {
	if aRowIndex < CMatrixReducingLimitARowIndex && CMatrixReducingSuccessBRow < CMatrixReducingSuccessMaxBRow {
		var tempDeleteColumnIndexList []int
		copy(tempDeleteColumnIndexList, deleteColumnIndexList)

		aRow := CMatrixReducingAArrayArray[aRowIndex]
		bRow := CMatrixReducingBArrayArray[CMatrixReducingSuccessBRow]

		successFlag := false
		bColumnIndex := 0
		for i, a := range aRow {
			deleteColumnInclude := false
			deleteColumnIndexListLen := len(deleteColumnIndexList)

			for j := 0; j < deleteColumnIndexListLen; j++ {
				if i == deleteColumnIndexList[j] {
					deleteColumnInclude = true
					break
				}
			}
			if !deleteColumnInclude {
				if bColumnIndex == CMatrixReducingLimitBColumn {
					tempDeleteColumnIndexList = append(tempDeleteColumnIndexList, i)
				} else {
					if a == bRow[bColumnIndex] {
						bColumnIndex = bColumnIndex + 1
					} else {
						tempDeleteColumnIndexList = append(tempDeleteColumnIndexList, i)
					}
					if bColumnIndex == CMatrixReducingLimitBColumn {
						CMatrixReducingSuccessBRow = CMatrixReducingSuccessBRow + 1
						successFlag = true
					}
				}
			}

		}
		if CMatrixReducingSuccessBRow == CMatrixReducingSuccessMaxBRow {
			CMatrixReducingResult = "Yes"
		} else if successFlag {
			judgeCMatrixReducing(aRowIndex+1, tempDeleteColumnIndexList, false)
		} else {
			if starFlag {
				CMatrixReducingSuccessBRow = 0
				tempDeleteColumnIndexList = []int{}
			} else {
				judgeCMatrixReducing(aRowIndex+1, deleteColumnIndexList, false)
			}
		}
	}

}

func CMatrixReducingRdr(rdr *bufio.Reader) string {
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
