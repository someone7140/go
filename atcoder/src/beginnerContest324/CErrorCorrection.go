package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CErrorCorrectionMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	inputStr := strings.Split(CErrorCorrectionReadLine(rdr), " ")
	n, _ := strconv.Atoi(inputStr[0])
	t := inputStr[1]
	tLen := len(t)
	var resultSlice []string

	for i := 0; i < n; i++ {
		s := CErrorCorrectionReadLine(rdr)
		if s == t {
			resultSlice = append(resultSlice, strconv.FormatInt(int64(i+1), 10))
		} else {
			sLen := len(s)
			// 長さが同じ場合
			if tLen == sLen {
				// 文字が違う数
				differentCount := 0
				for j := 0; j < tLen; j++ {
					tMoji := t[j : j+1]
					sMoji := s[j : j+1]
					if tMoji != sMoji {
						differentCount = differentCount + 1
					}
					if differentCount > 1 {
						break
					}

				}
				if differentCount <= 1 {
					resultSlice = append(resultSlice, strconv.FormatInt(int64(i+1), 10))
				}
				// 1文字tが短い場合
			} else if (sLen - 1) == tLen {
				// 文字が違う数
				differentCount := 0
				sIndex := 0
				for j := 0; j < tLen; j++ {
					tMoji := t[j : j+1]
					sMoji := s[sIndex : sIndex+1]
					if tMoji != sMoji {
						differentCount = differentCount + 1
						if differentCount <= 1 {
							sIndex = sIndex + 1
							sMoji = s[sIndex : sIndex+1]
							if tMoji != sMoji {
								differentCount = differentCount + 1
							}
							sIndex = sIndex + 1
						}
					} else {
						sIndex = sIndex + 1
					}
					if differentCount > 1 {
						break
					}
				}
				if differentCount <= 1 {
					resultSlice = append(resultSlice, strconv.FormatInt(int64(i+1), 10))
				}
				// 1文字sが短い場合
			} else if (tLen - 1) == sLen {
				// 文字が違う数
				differentCount := 0
				tIndex := 0
				for j := 0; j < sLen; j++ {
					tMoji := t[tIndex : tIndex+1]
					sMoji := s[j : j+1]
					if tMoji != sMoji {
						differentCount = differentCount + 1
						if differentCount <= 1 {
							tIndex = tIndex + 1
							tMoji = t[tIndex : tIndex+1]
							if tMoji != sMoji {
								differentCount = differentCount + 1
							}
							tIndex = tIndex + 1
						}
					} else {
						tIndex = tIndex + 1
					}
					if differentCount > 1 {
						break
					}
				}
				if differentCount <= 1 {
					resultSlice = append(resultSlice, strconv.FormatInt(int64(i+1), 10))
				}
			}
		}
	}

	fmt.Println(len(resultSlice))
	fmt.Println(strings.Join(resultSlice, " "))
}

func CErrorCorrectionReadLine(rdr *bufio.Reader) string {
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
