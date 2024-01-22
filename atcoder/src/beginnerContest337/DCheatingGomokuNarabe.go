package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DCheatingGomokuNarabeMain() {
	var h, w, k int
	fmt.Scan(&h, &w, &k)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var sArrayArray = make([][]string, h)
	for i := 0; i < h; i++ {
		sArrayArray[i] = strings.Split(DCheatingGomokuNarabeReadLine(rdr), "")
	}

	result := -1
	// 横の判定
	for i := 0; i < h; i++ {
		maruKosuu := 0
		kakikae := 0
		for j := 0; j < w; j++ {
			s := sArrayArray[i][j]
			if s == "x" {
				maruKosuu = 0
				kakikae = 0
			} else if s == "o" {
				maruKosuu = maruKosuu + 1
				if maruKosuu == k {
					if result == -1 || result > kakikae {
						result = kakikae
					}
					maruKosuu = maruKosuu - 1
					syoukyo := sArrayArray[i][j-k+1]
					if syoukyo == "." {
						kakikae = kakikae - 1
					}
				}
			} else {
				maruKosuu = maruKosuu + 1
				kakikae = kakikae + 1
				if maruKosuu == k {
					if result == -1 || result > kakikae {
						result = kakikae
					}
					maruKosuu = maruKosuu - 1
					syoukyo := sArrayArray[i][j-k+1]
					if syoukyo == "." {
						kakikae = kakikae - 1
					}
				}
			}
		}
	}

	// 縦の判定
	for i := 0; i < w; i++ {
		maruKosuu := 0
		kakikae := 0
		for j := 0; j < h; j++ {
			s := sArrayArray[j][i]
			if s == "x" {
				maruKosuu = 0
				kakikae = 0
			} else if s == "o" {
				maruKosuu = maruKosuu + 1
				if maruKosuu == k {
					if result == -1 || result > kakikae {
						result = kakikae
					}
					maruKosuu = maruKosuu - 1
					syoukyo := sArrayArray[j-k+1][i]
					if syoukyo == "." {
						kakikae = kakikae - 1
					}
				}
			} else {
				maruKosuu = maruKosuu + 1
				kakikae = kakikae + 1
				if maruKosuu == k {
					if result == -1 || result > kakikae {
						result = kakikae
					}
					maruKosuu = maruKosuu - 1
					syoukyo := sArrayArray[j-k+1][i]
					if syoukyo == "." {
						kakikae = kakikae - 1
					}
				}
			}
		}

	}

	fmt.Println(result)
}

func DCheatingGomokuNarabeReadLine(rdr *bufio.Reader) string {
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
