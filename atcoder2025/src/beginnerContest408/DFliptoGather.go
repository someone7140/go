package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DFliptoGatherRenzokuCount struct {
	val   int
	count int
}

type DFliptoGatherDp struct {
	sousa    int
	valCount int
}

func DFliptoGatherMain() {
	var t int
	fmt.Scan(&t)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var resultSlice []string

	for i := 0; i < t; i++ {
		DFliptoGatherRdr(rdr)
		sStrList := strings.Split(DFliptoGatherRdr(rdr), "")
		// まずはラングレス圧縮する
		var renzokuSlice []DFliptoGatherRenzokuCount
		before := -1
		renzokuIndex := 0
		for i, sStr := range sStrList {
			s, _ := strconv.Atoi(sStr)
			if i == 0 {
				renzokuSlice = append(renzokuSlice, DFliptoGatherRenzokuCount{
					val:   s,
					count: 1,
				})
			} else {
				if s == before {
					renzokuSlice[renzokuIndex] = DFliptoGatherRenzokuCount{
						val:   s,
						count: renzokuSlice[renzokuIndex].count + 1,
					}
				} else {
					renzokuSlice = append(renzokuSlice, DFliptoGatherRenzokuCount{
						val:   s,
						count: 1,
					})
					renzokuIndex = renzokuIndex + 1
				}
			}
			before = s
		}

		// ラングレス圧縮の配列からdpを行い最小の数を導き出す
		result := 0
		var dp [][]DFliptoGatherDp
		renzokuLen := len(renzokuSlice)
		for i, renzoku := range renzokuSlice {
			counts := make([]DFliptoGatherDp, 2)
			if i == 0 {
				if renzoku.val == 0 {
					//1を続けない時
					counts[0] = DFliptoGatherDp{
						sousa:    0,
						valCount: renzoku.count,
					}
					// 1を続ける時
					counts[1] = DFliptoGatherDp{
						sousa:    renzoku.count,
						valCount: renzoku.count,
					}
				} else {
					//1を続けない時
					counts[0] = DFliptoGatherDp{
						sousa:    renzoku.count,
						valCount: renzoku.count,
					}
					// 1を続ける時
					counts[1] = DFliptoGatherDp{
						sousa:    0,
						valCount: renzoku.count,
					}
				}
				dp = append(dp, counts)
			} else {
				before := dp[i-1]
				if renzoku.val == 0 {
					//1を続けない時
					if before[1].sousa+before[1].valCount < before[0].sousa {
						counts[0] = DFliptoGatherDp{
							sousa:    before[1].sousa + before[1].valCount,
							valCount: before[1].valCount + renzoku.count,
						}
					} else {
						counts[0] = DFliptoGatherDp{
							sousa:    before[0].sousa,
							valCount: before[0].valCount + renzoku.count,
						}
					}
					//1を続ける時
					if before[0].sousa+before[0].valCount+renzoku.count < before[1].sousa {
						counts[1] = DFliptoGatherDp{
							sousa:    before[0].sousa + before[0].valCount + renzoku.count,
							valCount: before[0].valCount + renzoku.count,
						}
					} else {
						counts[1] = DFliptoGatherDp{
							sousa:    before[1].sousa + renzoku.count,
							valCount: before[1].valCount + renzoku.count,
						}
					}
				} else {
					//1を続けない時
					if before[1].sousa+before[1].valCount+renzoku.count < before[0].sousa {
						counts[0] = DFliptoGatherDp{
							sousa:    before[1].sousa + before[1].valCount + renzoku.count,
							valCount: before[1].valCount + renzoku.count,
						}
					} else {
						counts[0] = DFliptoGatherDp{
							sousa:    before[0].sousa + renzoku.count,
							valCount: before[0].valCount + renzoku.count,
						}
					}
					//1を続ける時
					if before[0].sousa < before[1].sousa {
						counts[1] = DFliptoGatherDp{
							sousa:    before[0].sousa,
							valCount: renzoku.count,
						}
					} else {
						counts[1] = DFliptoGatherDp{
							sousa:    before[1].sousa,
							valCount: before[1].valCount + renzoku.count,
						}
					}
				}
				dp = append(dp, counts)
			}

		}
		zeroVal := dp[renzokuLen-1][0]
		oneVal := dp[renzokuLen-1][1]
		if zeroVal.sousa < oneVal.sousa {
			result = zeroVal.sousa
		} else {
			result = oneVal.sousa
		}
		resultSlice = append(resultSlice, strconv.FormatInt(int64(result), 10))
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func DFliptoGatherRdr(rdr *bufio.Reader) string {
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
