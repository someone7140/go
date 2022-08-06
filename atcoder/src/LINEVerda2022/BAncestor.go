package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BAncestorMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	pLine := BAncestorRdr(rdr)

	pArray := make([]int, n-1)
	pStrArray := strings.Split(pLine, " ")
	for i, pMoji := range pStrArray {
		pNum, _ := strconv.Atoi(pMoji)
		pArray[i] = int(pNum)
	}

	tempParent := pArray[n-2]
	result := 1

	for {
		if tempParent == 1 {
			break
		} else {
			tempParent = pArray[tempParent-2]
			result = result + 1
		}
	}

	fmt.Println(result)

}

func BAncestorRdr(rdr *bufio.Reader) string {
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
