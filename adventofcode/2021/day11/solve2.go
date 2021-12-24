package main

import (
	"bufio"
	"fmt"
	"os"
)

var g [][]int
var mark [][]int
var dx, dy []int

func FloodFill(x, y int) int {
	tot := 0
	if mark[x][y] == 1 {
		return tot
	}
	g[x][y]++
	if g[x][y] > 9 {
		mark[x][y] = 1
		tot++
		for i := 0; i < 8; i++ {
			nx := x + dx[i]
			ny := y + dy[i]
			if nx >= 0 && nx < 10 && ny >= 0 && ny < 10 && mark[nx][ny] == 0 {
				tot += FloodFill(nx, ny)
			}
		}
		g[x][y] = 0
	}
	return tot
}

func Flash() int {
	mark = make([][]int, 10)
	for i := range mark {
		mark[i] = make([]int, 10)
	}
	res := 0
	for i, ga := range g {
		for j, _ := range ga {
			res += FloodFill(i, j)
		}
	}
	return res
}

func main() {
	fi, _ := os.Open("data.txt")
	defer fi.Close()

	br := bufio.NewReader(fi)
	var s string
	g = make([][]int, 10)
	dx = []int{1, -1, 0, 0, -1, -1, 1, 1}
	dy = []int{0, 0, 1, -1, -1, 1, -1, 1}
	i := 0
	for {
		_, err := fmt.Fscan(br, &s)
		if err != nil {
			break
		}
		g[i] = make([]int, 10)
		for index, c := range s {
			g[i][index] = int(c - '0')
		}
		i++
	}

	for i := 0; i < 1000; i++ {
		if Flash() == 100 {
			fmt.Println(i + 1)
			return
		}
	}
}
