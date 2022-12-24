package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CPreviousPermutationMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	pStrArray := strings.Split(CPreviousPermutationReadLine(rdr), " ")
	pArray := make([]int, n)

	for i := 0; i < n; i++ {
		p, _ := strconv.Atoi(pStrArray[i])
		pArray[i] = p
	}

	irekaeFrom := -1
	// 昇順になってるところまでを見つける
	for i := n - 2; i >= 0; i = i - 1 {
		if pArray[i] < pArray[i+1] {
			// 何もしない
		} else {
			irekaeFrom = i
			break
		}
	}

	sortTarget := pArray[irekaeFrom+1:]
	lenSortTarget := len(sortTarget)
	var irekaeSlice []int

	if lenSortTarget == 1 {
		irekaeSlice = append(irekaeSlice, sortTarget[0])
		irekaeSlice = append(irekaeSlice, pArray[irekaeFrom])
	} else {
		irekaeTaisyou := -1
		for i := 0; i < lenSortTarget; i++ {
			if sortTarget[i] > pArray[irekaeFrom] {
				irekaeTaisyou = sortTarget[i-1]
				break
			}
		}
		if irekaeTaisyou == -1 {
			irekaeTaisyou = sortTarget[lenSortTarget-1]
		}
		for i := 0; i < lenSortTarget; i++ {
			if sortTarget[i] == irekaeTaisyou {
				irekaeSlice = append(irekaeSlice, pArray[irekaeFrom])
			} else {
				irekaeSlice = append(irekaeSlice, sortTarget[i])
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(irekaeSlice)))
		irekaeSlice = append([]int{irekaeTaisyou}, irekaeSlice...)
	}

	newSlice := append(pArray[:irekaeFrom], irekaeSlice...)
	var resultSlice []string
	for i := 0; i < n; i++ {
		resultSlice = append(resultSlice, strconv.FormatInt(int64(newSlice[i]), 10))
	}

	fmt.Println(strings.Join(resultSlice, " "))
}

func CPreviousPermutationReadLine(rdr *bufio.Reader) string {
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
