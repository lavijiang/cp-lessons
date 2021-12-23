package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(s1, s2 string) int {
	res := 0
	s1Map := make(map[rune]bool)
	for _, c := range s1 {
		s1Map[c] = true
	}
	for _, c := range s2 {
		if s1Map[c] {
			res++
		}
	}
	return res
}

func main() {
	fi, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	res := 0
	sc := bufio.NewScanner(fi)
	var numStr [10]string
	for sc.Scan() {
		s := sc.Text()
		ss := strings.Split(s, "|")
		if len(ss) < 2 {
			continue
		}

		ss[0] = strings.Trim(ss[0], " ")
		ss0 := strings.Split(ss[0], " ")
		for _, s := range ss0 {
			if len(s) == 2 {
				numStr[1] = s
			} else if len(s) == 3 {
				numStr[7] = s
			} else if len(s) == 4 {
				numStr[4] = s
			} else if len(s) == 7 {
				numStr[8] = s
			}
		}

		for _, s := range ss0 {
			if len(s) == 5 {
				if check(s, numStr[1]) == 2 {
					numStr[3] = s
				} else if check(s, numStr[4]) == 2 {
					numStr[2] = s
				} else {
					numStr[5] = s
				}
			} else if len(s) == 6 {
				if check(s, numStr[1]) == 1 {
					numStr[6] = s
				} else if check(s, numStr[4]) == 4 {
					numStr[9] = s
				} else {
					numStr[0] = s
				}
			}
		}

		ss[1] = strings.TrimSpace(ss[1])
		ss1 := strings.Split(ss[1], " ")
		tmp := 0
		for _, s := range ss1 {
			for index, s1 := range numStr {
				if len(s1) == len(s) && check(s1, s) == len(s) {
					tmp = tmp*10 + index
				}
			}
		}
		res += tmp
	}

	fmt.Println(res)
}
