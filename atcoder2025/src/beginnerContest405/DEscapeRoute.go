package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type DEscapeRouteRdrZahyou struct {
	h int
	w int
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	tListList := make([][]string, h)
	resultListList := make([][]string, h)
	var eSlice []DEscapeRouteRdrZahyou

	for i := 0; i < h; i++ {
		tList := strings.Split(DEscapeRouteRdr(rdr), "")
		tListList[i] = tList

		tList2 := make([]string, w)
		for j, t := range tList {
			tList2[j] = t
			if t == "E" {
				eSlice = append(eSlice, DEscapeRouteRdrZahyou{
					h: i,
					w: j,
				})
			}
		}
		resultListList[i] = tList2
	}

	// キューの初期設定
	tQue := list.New()
	for _, e := range eSlice {
		tQue.PushBack(e)
	}

	for {
		front := tQue.Front()
		if front == nil {
			break
		}
		val, ok := front.Value.(DEscapeRouteRdrZahyou)
		if ok {
			// 上
			if val.h > 0 {
				tResult := resultListList[val.h-1][val.w]
				// 壁でない&まだ通ってない
				if tResult == "." {
					resultListList[val.h-1][val.w] = "v"
					tQue.PushBack(DEscapeRouteRdrZahyou{
						h: val.h - 1,
						w: val.w,
					})
				}

			}
			// 下
			if val.h < h-1 {
				tResult := resultListList[val.h+1][val.w]
				// 壁でない&まだ通ってない
				if tResult == "." {
					resultListList[val.h+1][val.w] = "^"
					tQue.PushBack(DEscapeRouteRdrZahyou{
						h: val.h + 1,
						w: val.w,
					})
				}

			}
			// 左
			if val.w > 0 {
				tResult := resultListList[val.h][val.w-1]
				// 壁でない&まだ通ってない
				if tResult == "." {
					resultListList[val.h][val.w-1] = ">"
					tQue.PushBack(DEscapeRouteRdrZahyou{
						h: val.h,
						w: val.w - 1,
					})
				}

			}
			// 右
			if val.w < w-1 {
				tResult := resultListList[val.h][val.w+1]
				// 壁でない&まだ通ってない
				if tResult == "." {
					resultListList[val.h][val.w+1] = "<"
					tQue.PushBack(DEscapeRouteRdrZahyou{
						h: val.h,
						w: val.w + 1,
					})
				}

			}
			tQue.Remove(front)
		} else {
			break
		}
	}

	resultList := make([]string, h)
	for i, results := range resultListList {
		resultList[i] = strings.Join(results, "")
	}
	fmt.Println(strings.Join(resultList, "\n"))
}

func DEscapeRouteRdr(rdr *bufio.Reader) string {
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
