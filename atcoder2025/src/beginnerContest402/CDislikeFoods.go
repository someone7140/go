package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CDislikeFoodsMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var resultSlice []string
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	ryouriSyokuzaiCount := make([]int, m)
	tempResult := 0
	syokuzaiRyouriMap := map[int][]int{}

	for i := 0; i < m; i++ {
		ka := strings.Split(CDislikeFoodsRdr(rdr), " ")
		k, _ := strconv.Atoi(ka[0])
		aList := ka[1:]
		for j := 0; j < k; j++ {
			a, _ := strconv.Atoi(aList[j])
			ryouriSyokuzaiCount[i] = ryouriSyokuzaiCount[i] + 1
			ryouriList, ok := syokuzaiRyouriMap[a]
			if !ok {
				syokuzaiRyouriMap[a] = []int{i}
			} else {
				syokuzaiRyouriMap[a] = append(ryouriList, i)
			}
		}
	}

	bListStr := strings.Split(CDislikeFoodsRdr(rdr), " ")
	for i := 0; i < n; i++ {
		syookuzai, _ := strconv.Atoi(bListStr[i])
		ryouriList, ok := syokuzaiRyouriMap[syookuzai]
		if ok {
			lenRyouri := len(ryouriList)
			for j := 0; j < lenRyouri; j++ {
				ryouri := ryouriList[j]
				if ryouriSyokuzaiCount[ryouri] > 0 {
					ryouriSyokuzaiCount[ryouri] = ryouriSyokuzaiCount[ryouri] - 1
					if ryouriSyokuzaiCount[ryouri] == 0 {
						tempResult = tempResult + 1
					}
				}
			}
		}
		resultSlice = append(resultSlice, strconv.FormatInt(int64(tempResult), 10))
	}
	fmt.Println(strings.Join(resultSlice, "\n"))

}

func CDislikeFoodsRdr(rdr *bufio.Reader) string {
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
