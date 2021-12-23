package main

import (
	"bufio"
	"fmt"
	. "fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func main() {
	fi, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	var s string

	br := bufio.NewReader(fi)
	_, c := Fscanln(br, &s)
	if c == io.EOF {
		return
	}

	numStrList := strings.Split(s, ",")
	nums := []int{}
	for _, numStr := range numStrList {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	tag := nums[len(nums)/2]

	res := 0
	for _, num := range nums {
		res += abs(num - tag)
	}
	fmt.Println(res)
}
