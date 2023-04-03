package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CPathGraphMain() {
	var n, m int
	fmt.Scan(&n, &m)

	uvMap := map[int][]int{}
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < m; i++ {
		uvArr := strings.Split(CPathGraphReadLine(rdr), " ")
		u, _ := strconv.Atoi(uvArr[0])
		v, _ := strconv.Atoi(uvArr[1])

		uRes, uOk := uvMap[u]
		if uOk {
			uvMap[u] = append(uRes, v)
		} else {
			uvMap[u] = []int{v}
		}

		vRes, vOk := uvMap[v]
		if vOk {
			uvMap[v] = append(vRes, u)
		} else {
			uvMap[v] = []int{u}
		}
	}

	ikkoCount := 0
	ikkoValue := -1
	twoOverCount := 0
	for k, v := range uvMap {
		lenV := len(v)
		if lenV == 1 {
			ikkoCount = ikkoCount + 1
			ikkoValue = k
		} else if lenV > 2 {
			twoOverCount = twoOverCount + 1
			break
		}
	}

	if ikkoCount != 2 || twoOverCount != 0 {
		fmt.Println("No")
	} else {
		routeCount := 2
		prevPos := ikkoValue
		nowPos := uvMap[ikkoValue][0]
		nextArr := uvMap[nowPos]
		for {
			if routeCount == n {
				break
			}
			routeCount = routeCount + 1
			if len(nextArr) != 2 {
				break
			} else {
				if nextArr[0] == prevPos {
					prevPos = nowPos
					nowPos = nextArr[1]
					nextArr = uvMap[nowPos]
				} else if nextArr[1] == prevPos {
					prevPos = nowPos
					nowPos = nextArr[0]
					nextArr = uvMap[nowPos]
				} else {
					break
				}

			}
		}

		if routeCount == n {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}

func CPathGraphReadLine(rdr *bufio.Reader) string {
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
