package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := DMatchorNotReadLine(rdr)
	t := DMatchorNotReadLine(rdr)
	lenS := len(s)
	lenT := len(t)
	var resultSlice = make([]string, lenS)

	if lenS-1 != lenT {
		for i := 0; i < lenS; i++ {
			resultSlice[i] = "No"
		}
	} else {
		lastS := ""
		sArray := make([]string, lenT)
		for i, c := range s {
			sMoji := string([]rune{c})
			if i == lenS-1 {
				lastS = sMoji
			} else {
				sArray[i] = sMoji
			}
		}
		tArray := make([]string, lenT)
		for i, c := range t {
			tMoji := string([]rune{c})
			tArray[i] = tMoji
		}

		icchiCount := 0
		// 1回目の比較
		for i := 0; i < lenT; i++ {
			tempS := sArray[i]
			tempT := tArray[i]
			if tempS == tempT {
				icchiCount = icchiCount + 1
			} else if tempS == "?" || tempT == "?" {
				icchiCount = icchiCount + 1
			}
		}
		if icchiCount == lenT {
			resultSlice[lenT] = "Yes"
		} else {
			resultSlice[lenT] = "No"
		}
		// ここからループ
		for i := lenT - 1; i >= 0; i-- {
			tempLast := sArray[i]
			tempT := tArray[i]
			if icchiCount == lenT {
				if tempT == lastS {
					resultSlice[i] = "Yes"
				} else if tempT == "?" || lastS == "?" {
					resultSlice[i] = "Yes"
				} else {
					icchiCount = icchiCount - 1
					resultSlice[i] = "No"
				}
			} else {
				if tempT == lastS {
					if tempT == tempLast {
						// 何もしない
					} else if tempT == "?" || tempLast == "?" {
						// 何もしない
					} else {
						icchiCount = icchiCount - 1
					}
				} else if tempT == "?" || lastS == "?" {
					if tempT == tempLast {
						// 何もしない
					} else if tempT == "?" || tempLast == "?" {
						// 何もしない
					} else {
						icchiCount = icchiCount - 1
					}
				} else {
					if tempT == tempLast {
						icchiCount = icchiCount + 1
					} else if tempT == "?" || tempLast == "?" {
						icchiCount = icchiCount + 1
					} else {
						// 何もしない
					}
				}

				if icchiCount == lenT {
					resultSlice[i] = "Yes"
				} else {
					resultSlice[i] = "No"
				}
			}
			sArray[i] = lastS
			lastS = tempLast
		}

	}
	fmt.Println(strings.Join(resultSlice, "\n"))
}

func DMatchorNotReadLine(rdr *bufio.Reader) string {
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
