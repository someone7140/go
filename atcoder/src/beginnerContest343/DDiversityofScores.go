package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DDiversityofScoresMain() {
	var n, t int
	fmt.Scan(&n, &t)

	var resultSlice []string
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	ninzuuMap := map[int]int{}
	hitoMap := map[int]int{}

	// 最初は0だけ
	syurui := 1
	ninzuuMap[0] = n
	for i := 1; i <= n; i++ {
		hitoMap[i] = 0
	}

	for i := 0; i < t; i++ {
		abArray := strings.Split(DDiversityofScoresReadLine(rdr), " ")
		a, _ := strconv.Atoi(abArray[0])
		b, _ := strconv.Atoi(abArray[1])

		nowScore := hitoMap[a]
		newScore := nowScore + b

		// 今のスコアを減らす
		ninzuu := ninzuuMap[nowScore]
		if ninzuu == 1 {
			delete(ninzuuMap, nowScore)
			syurui = syurui - 1
		} else {
			ninzuuMap[nowScore] = ninzuu - 1
		}

		// 新しいスコアを増やす
		ninzuuNew, ok := ninzuuMap[newScore]
		if ok {
			ninzuuMap[newScore] = ninzuuNew + 1
		} else {
			ninzuuMap[newScore] = 1
			syurui = syurui + 1
		}
		hitoMap[a] = newScore

		resultSlice = append(resultSlice, strconv.FormatInt(int64(syurui), 10))
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func DDiversityofScoresReadLine(rdr *bufio.Reader) string {
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
