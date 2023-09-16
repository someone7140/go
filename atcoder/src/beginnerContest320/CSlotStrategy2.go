package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CSlotStrategy2Main() {
	var m int
	fmt.Scan(&m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	s1 := strings.Split(CSlotStrategy2ReadLine(rdr), "")
	s2 := strings.Split(CSlotStrategy2ReadLine(rdr), "")
	s3 := strings.Split(CSlotStrategy2ReadLine(rdr), "")

	result := -1
	s1Map := map[int][]int{}
	s2Map := map[int][]int{}
	s3Map := map[int][]int{}

	for i := 0; i < m; i++ {
		s1Num, _ := strconv.Atoi(s1[i])
		s1Arr, ok1 := s1Map[s1Num]
		if !ok1 {
			s1Map[s1Num] = []int{i}
		} else {
			s1Map[s1Num] = append(s1Arr, i)
		}

		s2Num, _ := strconv.Atoi(s2[i])
		s2Arr, ok2 := s2Map[s2Num]
		if !ok2 {
			s2Map[s2Num] = []int{i}
		} else {
			s2Map[s2Num] = append(s2Arr, i)
		}

		s3Num, _ := strconv.Atoi(s3[i])
		s3Arr, ok3 := s3Map[s3Num]
		if !ok3 {
			s3Map[s3Num] = []int{i}
		} else {
			s3Map[s3Num] = append(s3Arr, i)
		}

	}

	var judgeFunc func(s1Arr []int, s2Arr []int, s3Arr []int)
	judgeFunc = func(s1Arr []int, s2Arr []int, s3Arr []int) {
		// s1の秒数
		sec := s1Arr[0]

		// s2の秒数
		s2Ok := false
		plusFlugS2 := false
		for {
			tempSec := sec % m
			for _, sr := range s2Arr {
				if sr > tempSec && !plusFlugS2 || sr >= tempSec && plusFlugS2 {
					sec = (sec/m)*m + sr
					s2Ok = true
					if s2Ok {
						break
					}
				}
			}
			if s2Ok {
				break
			}
			sec = ((sec / m) + 1) * m
			plusFlugS2 = true
		}

		// s3の秒数
		s3Ok := false
		plusFlugS3 := false
		for {

			tempSec := sec % m
			for _, sr := range s3Arr {
				if sr > tempSec && !plusFlugS3 || sr >= tempSec && plusFlugS3 {
					sec = (sec/m)*m + sr
					s3Ok = true
					if s3Ok {
						break
					}
				}
			}
			if s3Ok {
				break
			}
			sec = ((sec / m) + 1) * m
			plusFlugS3 = true
		}

		if result == -1 {
			result = sec
		} else if result > sec {
			result = sec
		}
	}

	for i := 0; i < 10; i++ {
		s1Arr, ok1 := s1Map[i]
		s2Arr, ok2 := s2Map[i]
		s3Arr, ok3 := s3Map[i]

		if ok1 && ok2 && ok3 {
			// 6通り試す
			judgeFunc(s1Arr, s2Arr, s3Arr)
			judgeFunc(s1Arr, s3Arr, s2Arr)
			judgeFunc(s2Arr, s1Arr, s3Arr)
			judgeFunc(s2Arr, s3Arr, s1Arr)
			judgeFunc(s3Arr, s2Arr, s1Arr)
			judgeFunc(s3Arr, s1Arr, s2Arr)
		}

	}
	fmt.Println(result)

}

func CSlotStrategy2ReadLine(rdr *bufio.Reader) string {
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
