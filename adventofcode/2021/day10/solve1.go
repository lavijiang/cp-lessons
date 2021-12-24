package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	s []int
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) push(v int) {
	s.s = append(s.s, v)
}

func (s *stack) size() int {
	return len(s.s)
}

func (s *stack) top() int {
	return s.s[len(s.s)-1]
}

func (s *stack) pop() {
	s.s = s.s[:len(s.s)-1]
}

func main() {
	fi, _ := os.Open("data.txt")
	defer fi.Close()

	br := bufio.NewReader(fi)
	var s string
	ress := [4]int64{}
	for {
		_, err := fmt.Fscan(br, &s)
		if err != nil {
			break
		}
		fmt.Println(s)
		st := NewStack()
		for _, c := range s {
			if c == '(' {
				st.push(0)
			} else if c == '[' {
				st.push(1)
			} else if c == '{' {
				st.push(2)
			} else if c == '<' {
				st.push(3)
			} else if c == ')' {
				if st.top() != 0 || st.size() == 0 {
					ress[0]++
					break
				}
				st.pop()
			} else if c == ']' {
				if st.top() != 1 || st.size() == 0 {
					ress[1]++
					break
				}
				st.pop()
			} else if c == '}' {
				if st.top() != 2 || st.size() == 0 {
					ress[2]++
					break
				}
				st.pop()
			} else if c == '>' {
				if st.top() != 3 || st.size() == 0 {
					ress[3]++
					break
				}
				st.pop()
			}
		}
	}
	fmt.Println(ress)
	fmt.Println(ress[0]*3 + ress[1]*57 + ress[2]*1197 + ress[3]*25137)
}
