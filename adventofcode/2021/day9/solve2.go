package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var dx, dy []int
var flag [][]int
var ss []string

func main() {
	dx = []int{1, -1, 0, 0}
	dy = []int{0, 0, 1, -1}

	fi, _ := os.Open("data.txt")
	defer fi.Close()
	br := bufio.NewReader(fi)

	var s string
	flag = make([][]int, 100)
	ss = make([]string, 100)
	for i := range ss {
		_, c := fmt.Fscanln(br, &s)
		if c == io.EOF {
			break
		}
		ss[i] = s
		flag[i] = make([]int, 100)
	}

	ress := [3]int{}
	for i, s := range ss {
		for j := range s {
			//fmt.Println(i, j)
			tmp := FloodFill(i, j)
			if ress[0] <= tmp {
				ress[2] = ress[1]
				ress[1] = ress[0]
				ress[0] = tmp
			} else if ress[1] <= tmp {
				ress[2] = ress[1]
				ress[1] = tmp
			} else if ress[2] <= tmp {
				ress[2] = tmp
			}
		}
	}
	fmt.Println(ress[0] * ress[1] * ress[2])
}

func FloodFill(x, y int) int {
	if flag[x][y] == 1 || ss[x][y] == '9' {
		return 0
	}
	flag[x][y] = 1
	tot := 1
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		if nx >= 0 && nx < len(flag) && ny >= 0 && ny < len(flag[0]) && ss[nx][ny] != ss[x][y] {
			tot += FloodFill(nx, ny)
		}
	}
	return tot
}
