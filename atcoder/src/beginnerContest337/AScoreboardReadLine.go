package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AScoreboardReadLineMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	takahashi := 0
	aoki := 0

	for i := 0; i < n; i++ {
		aStrArray := strings.Split(AScoreboardReadLine(rdr), " ")
		takahashiTemp, _ := strconv.Atoi(aStrArray[0])
		aokiTemp, _ := strconv.Atoi(aStrArray[1])
		takahashi = takahashi + takahashiTemp
		aoki = aoki + aokiTemp
	}

	if takahashi > aoki {
		fmt.Println("Takahashi")
	} else if aoki > takahashi {
		fmt.Println("Aoki")
	} else {
		fmt.Println("Draw")
	}
}

func AScoreboardReadLine(rdr *bufio.Reader) string {
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
