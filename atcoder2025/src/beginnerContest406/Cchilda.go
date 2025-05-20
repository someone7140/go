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
	pStrList := strings.Split(CChildaRdr(rdr), " ")

	result := 0
	tempPList := []int{}
	len := 0
	a1a2Flag := false
	centerLargeIndex := -1
	centerSmallIndex := -1

	var resultJudge func() = func() {
		if len > 3 && a1a2Flag && centerLargeIndex > -1 && centerSmallIndex > -1 {
			result = result + 1
			fmt.Println(tempPList)
		}
	}
	var deleteSentou func() = func() {
		if centerLargeIndex == 1 {
			centerLargeIndex = -1
		} else {
			centerLargeIndex = centerLargeIndex - 1
		}
		if centerSmallIndex == 1 {
			centerSmallIndex = -1
		} else {
			centerSmallIndex = centerSmallIndex - 1
		}
		tempPList = tempPList[1:]
		len--
		if len < 2 || tempPList[0] >= tempPList[1] {
			a1a2Flag = false
		}
	}

	var lastJudge func() = func() {
		if len > 2 {
			last := tempPList[len-1]
			last1 := tempPList[len-2]
			last2 := tempPList[len-3]
			centerLargeFlag := last < last1 && last2 < last1
			centerSmallFlag := last > last1 && last2 > last1
			if centerLargeIndex == len-2 {
				// 何もしない
			} else if centerLargeFlag && centerLargeIndex < 0 {
				centerLargeIndex = len - 2
			} else if centerLargeFlag && centerLargeIndex > -1 {
				// 現状の手前まで削除する
				loopLen := centerLargeIndex
				for i := 0; i < loopLen; i++ {
					deleteSentou()
				}
				centerLargeIndex = len - 2 - loopLen
			}
			if centerSmallIndex == len-2 {
				// 何もしない
			} else if centerSmallFlag && centerSmallIndex < 0 {
				centerSmallIndex = len - 2
			} else if centerSmallFlag && centerSmallIndex > -1 {
				// 現状の手前まで削除する
				loopLen := centerSmallIndex
				for i := 0; i < loopLen; i++ {
					deleteSentou()
				}
				centerSmallIndex = len - 2 - loopLen
			}
			resultJudge()
		}
	}

	for _, pStr := range pStrList {
		p, _ := strconv.Atoi(pStr)
		if len < 2 {
			tempPList = append(tempPList, p)
			len++
			lastJudge()
			continue
		}
		if !a1a2Flag {
			if tempPList[0] < tempPList[1] {
				a1a2Flag = true
			} else if len > 1 {
				deleteSentou()
				lastJudge()
			}
			tempPList = append(tempPList, p)
			len++
			lastJudge()
			continue
		}
		tempPList = append(tempPList, p)
		len++
		lastJudge()
	}

	// 先頭から削除してジャッジ
	for i := 0; i < len; i++ {
		deleteSentou()
		lastJudge()
	}
	fmt.Println(result)
}

func CChildaRdr(rdr *bufio.Reader) string {
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
