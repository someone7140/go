package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DABCPuzzleMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	cArray := strings.Split(DABCPuzzleReadLine(rdr), "")
	rArray := strings.Split(DABCPuzzleReadLine(rdr), "")
	rFirstFlagArray := make([]int, n)
	cFirstFlagArray := make([]int, n)
	rcArrayArray := make([][]string, n)

	if rArray[0] != cArray[0] {
		fmt.Println("No")
	} else {
		// 初期配置
		for i := 0; i < n; i++ {
			if i == 0 {
				rArrayCopy := make([]string, n)
				copy(rArrayCopy, rArray)
				rcArrayArray[0] = rArrayCopy
			} else {
				rArray := make([]string, n)
				for j := 0; j < n; j++ {
					if j == 0 {
						rArray[0] = cArray[i]
					} else {
						rArray[j] = "."
					}
				}
				rcArrayArray[i] = rArray
			}
		}

		result := true
		rFirstFlagArray[0] = 1
		cFirstFlagArray[0] = 1
		// 行と列でそれぞれ判定していく
		for i := 0; i < n; i++ {
			// 行判定
			aFlag := false
			bFlag := false
			cFlag := false
			for j := 0; j < n; j++ {
				target := rcArrayArray[i][j]

				if target == "A" {
					if !aFlag {
						aFlag = true
					} else {
						if i < n-1 && rcArrayArray[i+1][j] == "." {
							rcArrayArray[i+1][j] = "A"
							rcArrayArray[i][j] = "."
						} else {
							result = false
						}
					}
				}
				if target == "B" {
					if !bFlag {
						bFlag = true
					} else {
						if i < n-1 && rcArrayArray[i+1][j] == "." {
							rcArrayArray[i+1][j] = "B"
							rcArrayArray[i][j] = "."
						} else {
							result = false
						}
					}
				}
				if target == "C" {
					if !cFlag {
						cFlag = true
					} else {
						if i < n-1 && rcArrayArray[i+1][j] == "." {
							rcArrayArray[i+1][j] = "C"
							rcArrayArray[i][j] = "."
						} else {
							result = false
						}
					}
				}

				if rFirstFlagArray[j] != 1 {
					if rArray[j] == rcArrayArray[i][j] {
						rFirstFlagArray[j] = 1
					}
				}
			}
			if !result {
				break
			}
			if !aFlag || !bFlag || !cFlag {
				for j := i; j < n; j++ {
					if rFirstFlagArray[j] == 1 && rcArrayArray[i][j] == "." {
						if !aFlag {
							rcArrayArray[i][j] = "A"
							aFlag = true
						} else if !bFlag {
							rcArrayArray[i][j] = "B"
							bFlag = true
						} else if !cFlag {
							rcArrayArray[i][j] = "C"
							cFlag = true
						}
					}
				}
			}
			if !aFlag || !bFlag || !cFlag {
				result = false
				break
			}

			// 列判定
			aFlag = false
			bFlag = false
			cFlag = false
			for j := 0; j < n; j++ {
				target := rcArrayArray[j][i]

				if target == "A" {
					if !aFlag {
						aFlag = true
					} else {
						if i < n-1 && rcArrayArray[j][i+1] == "." {
							rcArrayArray[j][i+1] = "A"
							rcArrayArray[j][i] = "."
						} else {
							result = false
						}
					}
				}
				if target == "B" {
					if !bFlag {
						bFlag = true
					} else {
						if i < n-1 && rcArrayArray[j][i+1] == "." {
							rcArrayArray[j][i+1] = "B"
							rcArrayArray[j][i] = "."
						} else {
							result = false
						}
					}
				}
				if target == "C" {
					if !cFlag {
						cFlag = true
					} else {
						if i < n-1 && rcArrayArray[j][i+1] == "." {
							rcArrayArray[j][i+1] = "C"
							rcArrayArray[j][i] = "."
						} else {
							result = false
						}
					}
				}

				if cFirstFlagArray[j] != 1 {
					if cArray[j] == rcArrayArray[j][i] {
						cFirstFlagArray[j] = 1
					}
				}
			}
			if !result {
				break
			}
			if !aFlag || !bFlag || !cFlag {
				for j := i + 1; j < n; j++ {
					if cFirstFlagArray[j] == 1 && rcArrayArray[j][i] == "." {
						if !aFlag {
							rcArrayArray[j][i] = "A"
							aFlag = true
						} else if !bFlag {
							rcArrayArray[j][i] = "B"
							bFlag = true
						} else if !cFlag {
							rcArrayArray[j][i] = "C"
							cFlag = true
						}
					}
				}
			}
			if !aFlag || !bFlag || !cFlag {
				result = false
				break
			}
		}

		if !result {
			fmt.Println("No")
		} else {
			var resultSlice = make([]string, n)
			for i := 0; i < n; i++ {
				resultSlice = append(resultSlice, strings.Join(rcArrayArray[i], ""))
			}
			fmt.Println("Yes")
			fmt.Println(strings.Join(resultSlice, "\n"))
		}
	}

}

func DABCPuzzleReadLine(rdr *bufio.Reader) string {
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
