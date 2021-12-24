package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	ress := []int64{}
	for {
		_, err := fmt.Fscan(br, &s)
		if err != nil {
			break
		}
		st := NewStack()
		flag := 1
		for _, c := range s {
			if c == '(' {
				st.push(1)
			} else if c == '[' {
				st.push(2)
			} else if c == '{' {
				st.push(3)
			} else if c == '<' {
				st.push(4)
			} else if c == ')' {
				if st.top() != 1 || st.size() == 0 {
					flag = 0
					break
				}
				st.pop()
			} else if c == ']' {
				if st.top() != 2 || st.size() == 0 {
					flag = 0
					break
				}
				st.pop()
			} else if c == '}' {
				if st.top() != 3 || st.size() == 0 {
					flag = 0
					break
				}
				st.pop()
			} else if c == '>' {
				if st.top() != 4 || st.size() == 0 {
					flag = 0
					break
				}
				st.pop()
			}
		}

		if flag == 0 {
			continue
		}

		var tmpScore int64
		for st.size() > 0 {
			tmpScore *= 5
			tmpScore += int64(st.top())
			st.pop()
		}
		ress = append(ress, tmpScore)
	}
	sort.Slice(ress, func(i, j int) bool {
		return ress[i] < ress[j]
	})
	fmt.Println(ress[len(ress)/2])
}
