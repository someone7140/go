package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func BNumberBoxMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var result int64
	resultArrayArray := make([][]int, n)

	for i := 0; i < n; i++ {
		resultArray := make([]int, n)
		s := BNumberBoxRdr(rdr)
		for i2, c := range s {
			aMoji := string([]rune{c})
			a, _ := strconv.Atoi(aMoji)
			resultArray[i2] = a
		}
		resultArrayArray[i] = resultArray
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result2 := getResultBNumberBox(n, i, j, resultArrayArray, 0, 1)
			if result2 > result {
				result = result2
			}
			result3 := getResultBNumberBox(n, i, j, resultArrayArray, 0, -1)
			if result3 > result {
				result = result3
			}
			result4 := getResultBNumberBox(n, i, j, resultArrayArray, 1, 0)
			if result4 > result {
				result = result4
			}
			result5 := getResultBNumberBox(n, i, j, resultArrayArray, 1, 1)
			if result5 > result {
				result = result5
			}
			result6 := getResultBNumberBox(n, i, j, resultArrayArray, 1, -1)
			if result6 > result {
				result = result6
			}
			result7 := getResultBNumberBox(n, i, j, resultArrayArray, -1, 0)
			if result7 > result {
				result = result7
			}
			result8 := getResultBNumberBox(n, i, j, resultArrayArray, -1, 1)
			if result8 > result {
				result = result8
			}
			result9 := getResultBNumberBox(n, i, j, resultArrayArray, -1, -1)
			if result9 > result {
				result = result9
			}
		}
	}
	fmt.Println(result)
}

func getResultBNumberBox(n int, i int, j int, aArrayArray [][]int, directionYoko int, directionTate int) int64 {
	resultStr := strconv.FormatInt(int64(aArrayArray[i][j]), 10)
	yoko := j
	tate := i
	for l := 0; l < n-1; l++ {
		yoko = yoko + directionYoko
		if yoko < 0 {
			yoko = n - 1
		}
		if yoko > n-1 {
			yoko = 0
		}
		tate = tate + directionTate
		if tate < 0 {
			tate = n - 1
		}
		if tate > n-1 {
			tate = 0
		}
		a := aArrayArray[tate][yoko]
		resultStr = resultStr + strconv.FormatInt(int64(a), 10)
	}
	result, _ := strconv.ParseInt(resultStr, 10, 64)
	return result
}

func BNumberBoxRdr(rdr *bufio.Reader) string {
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
