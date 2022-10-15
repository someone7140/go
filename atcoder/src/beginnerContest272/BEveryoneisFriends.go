package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BEveryoneisFriendsMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var aSliceSlice = make([][]int, n)
	for i := 0; i < n; i++ {
		var aSlice = make([]int, n)
		aSliceSlice[i] = aSlice
	}

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < m; i++ {
		xLine := BEveryoneisFriendsRdr(rdr)
		xStrArray := strings.Split(xLine, " ")
		k, _ := strconv.Atoi(xStrArray[0])

		for j := 1; j < k; j++ {
			x1, _ := strconv.Atoi(xStrArray[j])
			for l := j + 1; l <= k; l++ {
				x2, _ := strconv.Atoi(xStrArray[l])
				aSliceSlice[x1-1][x2-1] = 1
				aSliceSlice[x2-1][x1-1] = 1
			}
		}

	}

	result := "Yes"
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				val := aSliceSlice[i][j]
				if val < 1 {
					result = "No"
					break
				}
			}
		}
	}

	fmt.Println(result)
}

func BEveryoneisFriendsRdr(rdr *bufio.Reader) string {
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
