package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	s := CXXtoXXXRdr(rdr)
	sLength := len(s)
	sArray := make([]string, sLength)
	for i, c := range s {
		sMoji := string([]rune{c})
		sArray[i] = sMoji
	}

	t := CXXtoXXXRdr(rdr)
	tLength := len(t)
	tArray := make([]string, tLength)
	for i, c := range t {
		tMoji := string([]rune{c})
		tArray[i] = tMoji
	}

	result := "Yes"

	if sLength > tLength {
		result = "No"
	} else if sLength == tLength {
		if s != t {
			result = "No"
		}
	} else {
		sIndex := 0
		tKaisuu := 0
		tMoji := ""
		for i := 0; i < tLength; i++ {
			if i == 0 {
				tMoji = tArray[i]
				tKaisuu = 1
			} else {
				if i == tLength-1 {
					if sIndex == sLength-1 {
						if tMoji == tArray[i] {
							result = "No"
						} else {
							if sArray[sIndex] != tArray[i] {
								result = "No"
							}
						}
						break
					} else if tMoji == tArray[i] {
						tKaisuu = tKaisuu + 1
						sKaisuu := 0
						breakFlag := false
						for {
							sTemp := sArray[sIndex]
							if sTemp == tMoji && sIndex < sLength-1 {
								sKaisuu = sKaisuu + 1
							} else {
								if tKaisuu == 1 {
									if sKaisuu != 1 {
										result = "No"
									}
								} else {
									if sKaisuu == 1 {
										result = "No"
									}
								}
								breakFlag = true
							}
							if sIndex == sLength-1 || breakFlag {
								break
							} else {
								sIndex = sIndex + 1
							}
						}
					} else {
						sKaisuu := 0
						breakFlag := false
						for {
							sTemp := sArray[sIndex]
							if sTemp == tMoji && sIndex < sLength-1 {
								sKaisuu = sKaisuu + 1
							} else {
								if tKaisuu == 1 {
									if sKaisuu != 1 {
										result = "No"
									}
								} else {
									if sKaisuu == 1 {
										result = "No"
									}
								}
								breakFlag = true
							}
							if sIndex == sLength-1 || breakFlag {
								break
							} else {
								sIndex = sIndex + 1
							}
						}
						tMoji = tArray[i]
						sTemp := sArray[sIndex]
						if tMoji != sTemp {
							result = "No"
						} else {
							if sIndex != sLength-1 {
								result = "No"
							}
						}
					}
				} else if tMoji == tArray[i] {
					tKaisuu = tKaisuu + 1
				} else {
					loopFlag := true
					sKaisuu := 0
					breakFlag := false
					for {
						sTemp := sArray[sIndex]
						if sTemp == tMoji && sIndex < sLength-1 {
							sKaisuu = sKaisuu + 1
						} else {
							if tKaisuu == 1 {
								loopFlag = sKaisuu == 1
							} else {
								loopFlag = sKaisuu > 1
							}
							breakFlag = true
						}
						if sIndex == sLength-1 || breakFlag {
							break
						} else {
							sIndex = sIndex + 1
						}
					}
					if !loopFlag {
						result = "No"
						break
					}
					tMoji = tArray[i]
					tKaisuu = 1
				}
			}
		}
	}

	fmt.Println(result)
}

func CXXtoXXXRdr(rdr *bufio.Reader) string {
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
