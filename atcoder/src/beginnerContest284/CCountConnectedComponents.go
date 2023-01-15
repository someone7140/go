package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var CCountConnectedComponentsTottaArray []int
var CCountConnectedComponentsUvArrayArray [][]int

func CCountConnectedComponentsMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	CCountConnectedComponentsUvArrayArray = make([][]int, n)

	for i := 0; i < m; i++ {
		uvStrArray := strings.Split(CCountConnectedComponentsReadLine(rdr), " ")
		u, _ := strconv.Atoi(uvStrArray[0])
		v, _ := strconv.Atoi(uvStrArray[1])

		u = u - 1
		v = v - 1
		CCountConnectedComponentsUvArrayArray[u] = append(CCountConnectedComponentsUvArrayArray[u], v)
		CCountConnectedComponentsUvArrayArray[v] = append(CCountConnectedComponentsUvArrayArray[v], u)
	}

	result := 0
	CCountConnectedComponentsTottaArray = make([]int, n)
	for i := 0; i < n; i++ {
		tottaFlag := CCountConnectedComponentsTottaArray[i]
		if tottaFlag == 0 {
			result = result + 1
			CCountConnectedComponentsTottaArray[i] = 1
			uvArray := CCountConnectedComponentsUvArrayArray[i]
			CCountConnectedComponentsSaiki(uvArray)
		}
	}

	fmt.Println(result)

}

func CCountConnectedComponentsSaiki(uvArray []int) {
	for i := 0; i < len(uvArray); i++ {
		uv := uvArray[i]
		tottaFlag := CCountConnectedComponentsTottaArray[uv]
		if tottaFlag == 0 {
			CCountConnectedComponentsTottaArray[uv] = 1
			CCountConnectedComponentsSaiki(CCountConnectedComponentsUvArrayArray[uv])
		}
	}
}

func CCountConnectedComponentsReadLine(rdr *bufio.Reader) string {
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
