package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fi, _ := os.Open("data.txt")
	defer fi.Close()
	br := bufio.NewReader(fi)

	var s string
	var ss [3]string
	st := 0
	res := 0
	for {
		_, c := fmt.Fscanln(br, &s)
		if c == io.EOF {
			break
		}
		ss[0] = ss[1]
		ss[1] = ss[2]
		ss[2] = s
		if st == 1 {
			for i := range ss[1] {
				if ss[1][i] < ss[2][i] && ((i > 0 && ss[1][i] < ss[1][i-1]) || i == 0) &&
					((i+1 < len(ss[1]) && ss[1][i] < ss[1][i+1]) || i+1 == len(ss[1])) {
					res += int(ss[1][i]-'0') + 1
					//fmt.Println(int(ss[1][i] - '0'))
				}
			}
		}
		if st > 1 {
			for i := range ss[1] {
				if ss[1][i] < ss[0][i] && ss[1][i] < ss[2][i] &&
					((i > 0 && ss[1][i] < ss[1][i-1]) || i == 0) &&
					((i+1 < len(ss[1]) && ss[1][i] < ss[1][i+1]) || i+1 == len(ss[1])) {
					res += int(ss[1][i]-'0') + 1
					//fmt.Println(int(ss[1][i] - '0'))
				}
			}
		}
		st++
	}

	for i := range ss[2] {
		if ss[2][i] < ss[1][i] && ((i > 0 && ss[2][i] < ss[2][i-1]) || i == 0) &&
			((i+1 < len(ss[2]) && ss[2][i] < ss[2][i+1]) || i+1 == len(ss[2])) {
			res += int(ss[2][i]-'0') + 1
			//fmt.Println(int(ss[2][i] - '0'))
		}
	}
	fmt.Println(res)
}
