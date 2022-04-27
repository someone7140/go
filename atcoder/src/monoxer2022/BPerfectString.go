package main

import (
	"fmt"
	"regexp"
)

func BPerfectStringMain() {
	var s string
	fmt.Scan(&s)

	ooMojiFlag := false
	koMojiFlag := false
	chigauFlag := true
	sMap := map[string]int{}

	for _, c := range s {
		sMoji := string([]rune{c})
		if !ooMojiFlag {
			ooMojiFlag = BPerfectString_check_regexp(`[A-Z]`, sMoji)
		}
		if !koMojiFlag {
			koMojiFlag = BPerfectString_check_regexp(`[a-z]`, sMoji)
		}
		_, ok := sMap[sMoji]
		if ok {
			chigauFlag = false
			break
		} else {
			sMap[sMoji] = 1
		}
	}
	if ooMojiFlag && koMojiFlag && chigauFlag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func BPerfectString_check_regexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
