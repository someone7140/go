package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DTeleportMazeMain() {
	var h, w int
	fmt.Scan(&h, &w)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	sArrayArray := make([][]string, h)
	countArrayArray := make([][]int, h)
	warpMap := make(map[string][][]int)
	for i := 0; i < h; i++ {
		sStrArray := strings.Split(DTeleportMazeRdr(rdr), "")
		sArrayArray[i] = sStrArray
		countArrayArray[i] = make([]int, w)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if sArrayArray[i][j] != "." && sArrayArray[i][j] != "#" {
				zahyouArray, ok := warpMap[sArrayArray[i][j]]
				if !ok {
					warpMap[sArrayArray[i][j]] = [][]int{[]int{i, j}}
				} else {
					warpMap[sArrayArray[i][j]] = append(zahyouArray, []int{i, j})
				}
			}
			countArrayArray[i][j] = -1
		}
	}

	var calcFunc func(hParam int, wParam int, nowCount int)
	calcFunc = func(hParam int, wParam int, nowCount int) {
		nextCount := nowCount + 1
		if nowCount == 0 {
			if sArrayArray[hParam][wParam+1] != "#" {
				countArrayArray[hParam][wParam+1] = nextCount
				if sArrayArray[hParam][wParam+1] == "." {
					calcFunc(hParam, wParam+1, nextCount)
				} else {
					zahyouArrray, ok := warpMap[sArrayArray[hParam][wParam+1]]
					if ok {
						for _, z := range zahyouArrray {
							if strconv.FormatInt(int64(hParam), 10)+"-"+strconv.FormatInt(int64(wParam), 10) !=
								strconv.FormatInt(int64(z[0]), 10)+"-"+strconv.FormatInt(int64(z[1]), 10) && sArrayArray[z[0]][z[1]] != "#" {
								if countArrayArray[z[0]][z[1]] == -1 || countArrayArray[z[0]][z[1]] > nextCount {
									countArrayArray[z[0]][z[1]] = nextCount
									calcFunc(z[0], z[1], nextCount+1)
								}
							}
						}
					}
				}
			}
			if sArrayArray[hParam+1][wParam] != "#" {
				if countArrayArray[hParam+1][wParam] == -1 || countArrayArray[hParam+1][wParam] > nextCount {
					countArrayArray[hParam+1][wParam] = nextCount
					if sArrayArray[hParam+1][wParam] == "." {
						calcFunc(hParam+1, wParam, nextCount)
					} else {
						zahyouArrray, ok := warpMap[sArrayArray[hParam+1][wParam]]
						if ok {
							for _, z := range zahyouArrray {
								if strconv.FormatInt(int64(hParam), 10)+"-"+strconv.FormatInt(int64(wParam), 10) !=
									strconv.FormatInt(int64(z[0]), 10)+"-"+strconv.FormatInt(int64(z[1]), 10) && sArrayArray[z[0]][z[1]] != "#" {
									if countArrayArray[z[0]][z[1]] == -1 || countArrayArray[z[0]][z[1]] > nextCount {
										countArrayArray[z[0]][z[1]] = nextCount
										calcFunc(z[0], z[1], nextCount+1)
									}
								}
							}
						}
					}
				}
			}
		} else {
			if hParam > 0 && sArrayArray[hParam-1][wParam] != "#" {
				if countArrayArray[hParam-1][wParam] == -1 || countArrayArray[hParam-1][wParam] > nextCount {
					countArrayArray[hParam-1][wParam] = nextCount
					if sArrayArray[hParam-1][wParam] == "." {
						calcFunc(hParam-1, wParam, nextCount)
					} else {
						zahyouArrray, ok := warpMap[sArrayArray[hParam-1][wParam]]
						if ok {
							for _, z := range zahyouArrray {
								if strconv.FormatInt(int64(hParam), 10)+"-"+strconv.FormatInt(int64(wParam), 10) !=
									strconv.FormatInt(int64(z[0]), 10)+"-"+strconv.FormatInt(int64(z[1]), 10) && sArrayArray[z[0]][z[1]] != "#" {
									if countArrayArray[z[0]][z[1]] == -1 || countArrayArray[z[0]][z[1]] > nextCount {
										countArrayArray[z[0]][z[1]] = nextCount
										calcFunc(z[0], z[1], nextCount+1)
									}
								}
							}
						}
					}
				}
			}
			if wParam > 0 && sArrayArray[hParam][wParam-1] != "#" {
				if countArrayArray[hParam][wParam-1] == -1 || countArrayArray[hParam][wParam-1] > nextCount {
					countArrayArray[hParam][wParam-1] = nextCount
					if sArrayArray[hParam][wParam-1] == "." {
						calcFunc(hParam, wParam-1, nextCount)
					} else {
						zahyouArrray, ok := warpMap[sArrayArray[hParam][wParam-1]]
						if ok {
							for _, z := range zahyouArrray {
								if strconv.FormatInt(int64(hParam), 10)+"-"+strconv.FormatInt(int64(wParam), 10) !=
									strconv.FormatInt(int64(z[0]), 10)+"-"+strconv.FormatInt(int64(z[1]), 10) && sArrayArray[z[0]][z[1]] != "#" {
									if countArrayArray[z[0]][z[1]] == -1 || countArrayArray[z[0]][z[1]] > nextCount {
										countArrayArray[z[0]][z[1]] = nextCount
										calcFunc(z[0], z[1], nextCount+1)
									}
								}
							}
						}
					}
				}
			}
			if hParam < h-1 && sArrayArray[hParam+1][wParam] != "#" {
				if countArrayArray[hParam+1][wParam] == -1 || countArrayArray[hParam+1][wParam] > nextCount {
					countArrayArray[hParam+1][wParam] = nextCount
					if sArrayArray[hParam+1][wParam] == "." {
						calcFunc(hParam+1, wParam, nextCount)
					} else {
						zahyouArrray, ok := warpMap[sArrayArray[hParam+1][wParam]]
						if ok {
							for _, z := range zahyouArrray {
								if strconv.FormatInt(int64(hParam), 10)+"-"+strconv.FormatInt(int64(wParam), 10) !=
									strconv.FormatInt(int64(z[0]), 10)+"-"+strconv.FormatInt(int64(z[1]), 10) && sArrayArray[z[0]][z[1]] != "#" {
									if countArrayArray[z[0]][z[1]] == -1 || countArrayArray[z[0]][z[1]] > nextCount {
										countArrayArray[z[0]][z[1]] = nextCount
										calcFunc(z[0], z[1], nextCount+1)
									}
								}
							}
						}
					}
				}
			}
			if wParam < w-1 && sArrayArray[hParam][wParam+1] != "#" {
				if countArrayArray[hParam][wParam+1] == -1 || countArrayArray[hParam][wParam+1] > nextCount {
					countArrayArray[hParam][wParam+1] = nextCount
					if sArrayArray[hParam][wParam+1] == "." {
						calcFunc(hParam, wParam+1, nextCount)
					} else {
						zahyouArrray, ok := warpMap[sArrayArray[hParam][wParam+1]]
						if ok {
							for _, z := range zahyouArrray {
								if strconv.FormatInt(int64(hParam), 10)+"-"+strconv.FormatInt(int64(wParam), 10) !=
									strconv.FormatInt(int64(z[0]), 10)+"-"+strconv.FormatInt(int64(z[1]), 10) && sArrayArray[z[0]][z[1]] != "#" {
									if countArrayArray[z[0]][z[1]] == -1 || countArrayArray[z[0]][z[1]] > nextCount {
										countArrayArray[z[0]][z[1]] = nextCount
										calcFunc(z[0], z[1], nextCount+1)
									}
								}
							}
						}
					}
				}
			}
		}
	}
	countArrayArray[0][0] = 0
	calcFunc(0, 0, 0)

	result := countArrayArray[h-1][w-1]
	if result < 1 {
		result = -1
	}
	fmt.Println(result)
}

func DTeleportMazeRdr(rdr *bufio.Reader) string {
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
