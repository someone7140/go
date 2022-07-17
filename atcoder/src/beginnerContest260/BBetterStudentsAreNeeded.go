package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	index int
	score int
}

func BBetterStudentsAreNeededMain() {
	var n, x, y, z int
	fmt.Scan(&n, &x, &y, &z)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	a := BBetterStudentsAreNeededRdr(rdr)
	aArray := make([]Student, n)
	aStrArray := strings.Split(a, " ")
	for i, aMoji := range aStrArray {
		aNum, _ := strconv.Atoi(aMoji)
		aArray[i] = Student{
			index: i,
			score: aNum,
		}
	}

	b := BBetterStudentsAreNeededRdr(rdr)
	bArray := make([]Student, n)
	bStrArray := strings.Split(b, " ")
	sumArray := make([]Student, n)
	for i, bMoji := range bStrArray {
		bNum, _ := strconv.Atoi(bMoji)
		bArray[i] = Student{
			index: i,
			score: bNum,
		}
		sumArray[i] = Student{
			index: i,
			score: aArray[i].score + bNum,
		}
	}

	resultMap := map[int]int{}

	aSorted := make([]Student, n)
	copy(aSorted, aArray)
	sort.Slice(aSorted, func(i, j int) bool {
		if aSorted[i].score != aSorted[j].score {
			return aSorted[i].score > aSorted[j].score
		} else {
			return aSorted[i].index < aSorted[j].index
		}
	})
	for i := 0; i < x; i++ {
		resultMap[aSorted[i].index] = 1
	}

	bSorted := make([]Student, n)
	copy(bSorted, bArray)
	sort.Slice(bSorted, func(i, j int) bool {
		if bSorted[i].score != bSorted[j].score {
			return bSorted[i].score > bSorted[j].score
		} else {
			return bSorted[i].index < bSorted[j].index
		}
	})
	yCount := 0
	if y > 0 {
		for i := 0; i < n; i++ {
			_, ok := resultMap[bSorted[i].index]
			if !ok {
				resultMap[bSorted[i].index] = 1
				yCount = yCount + 1
			}
			if yCount == y {
				break
			}
		}
	}

	sumSorted := make([]Student, n)
	copy(sumSorted, sumArray)
	sort.Slice(sumSorted, func(i, j int) bool {
		if sumSorted[i].score != sumSorted[j].score {
			return sumSorted[i].score > sumSorted[j].score
		} else {
			return sumSorted[i].index < sumSorted[j].index
		}
	})
	sumCount := 0
	if z > 0 {
		for i := 0; i < n; i++ {
			_, ok := resultMap[sumSorted[i].index]
			if !ok {
				resultMap[sumSorted[i].index] = 1
				sumCount = sumCount + 1
			}
			if sumCount == z {
				break
			}
		}
	}

	var resultSlice []string
	for i := 0; i < n; i++ {
		_, ok := resultMap[i]
		if ok {
			iStr := strconv.FormatInt(int64(i+1), 10)
			resultSlice = append(resultSlice, iStr)
		}

	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func BBetterStudentsAreNeededRdr(rdr *bufio.Reader) string {
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
