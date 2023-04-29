package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BSameMapintheRPGWorldMain() {
	var h, w int
	fmt.Scan(&h, &w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aArrayArray := make([][]string, h)

	// まずはAの設定
	for i := 0; i < h; i++ {
		aArray := make([]string, w)
		row := BSameMapintheRPGWorldLine(rdr)
		for j, c := range row {
			aArray[j] = string([]rune{c})
		}
		aArrayArray[i] = aArray
	}

	// Bは文字列連結するだけ
	bStr := ""
	for i := 0; i < h; i++ {
		bStr = bStr + BSameMapintheRPGWorldLine(rdr)
	}

	// AとBを比較
	result := "No"
	var compareAArrayArray [][]string
	for i := 0; i < h; i++ {
		if i == 0 {
			compareAArrayArray = aArrayArray
		} else if result == "Yes" {
			break
		} else {
			newCompareAArrayArray := make([][]string, h)
			for j := 0; j < h; j++ {
				if j == h-1 {
					newCompareAArrayArray[0] = compareAArrayArray[j]
				} else {
					newCompareAArrayArray[j+1] = compareAArrayArray[j]
				}
			}
			compareAArrayArray = newCompareAArrayArray
		}
		var compareAArrayArray2 [][]string

		for j := 0; j < w; j++ {
			if j == 0 {
				compareAArrayArray2 = compareAArrayArray
			} else {
				var newCompareAArrayArray2 [][]string
				newCompareAArrayArray2 = make([][]string, h)
				for k := 0; k < h; k++ {
					compareAArray2 := make([]string, w)
					for l := 0; l < w; l++ {
						compareAArray2[l] = compareAArrayArray2[k][l]
					}
					newCompareAArrayArray2[k] = compareAArray2
				}
				for k := 0; k < w; k++ {
					for l := 0; l < h; l++ {
						if k == w-1 {
							newCompareAArrayArray2[l][0] = compareAArrayArray2[l][k]
						} else {
							newCompareAArrayArray2[l][k+1] = compareAArrayArray2[l][k]
						}
					}
				}
				compareAArrayArray2 = newCompareAArrayArray2
			}
			// 比較
			var aStr = ""
			for k := 0; k < h; k++ {
				aStr = aStr + strings.Join(compareAArrayArray2[k], "")
			}
			if aStr == bStr {
				result = "Yes"
				break
			}
		}
	}
	fmt.Println(result)

}

func BSameMapintheRPGWorldLine(rdr *bufio.Reader) string {
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
