package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := DScopeReadLine(rdr)

	kakkkoDeque := list.New()
	alphabetDeque := list.New()
	sMap := map[string]bool{}

	result := "Yes"
	count := 0
	for _, c := range s {
		sMoji := string([]rune{c})
		if sMoji == "(" {
			if count != 0 {
				kakkkoDeque.PushBack(count)
				count = 0
			}
		} else if sMoji == ")" {
			if count != 0 {
				for i := 0; i < count; i++ {
					alphabet := alphabetDeque.Back().Value.(string)
					sMap[alphabet] = false
					alphabetDeque.Remove(alphabetDeque.Back())
				}
				if kakkkoDeque.Len() > 0 {
					count = kakkkoDeque.Back().Value.(int)
					kakkkoDeque.Remove(kakkkoDeque.Back())
				} else {
					count = 0
				}
			}
		} else {
			count = count + 1
			flag, ok := sMap[sMoji]
			if ok && flag {
				result = "No"
				break
			}
			sMap[sMoji] = true
			alphabetDeque.PushBack(sMoji)
		}
	}

	fmt.Println(result)
}

func DScopeReadLine(rdr *bufio.Reader) string {
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
