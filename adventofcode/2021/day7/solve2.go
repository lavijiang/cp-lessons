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

func Calc(x int, nums []int) int {
	res := 0
	for _, num := range nums {
		res += (abs(x-num) * (abs(x-num) + 1)) / 2
	}
	return res
}

//三分法 凹函数
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

	l := nums[0]
	r := nums[len(nums)-1]
	for {
		if l >= r {
			break
		}
		mid := (l + r) >> 1
		mid2 := (mid + r) >> 1
		if Calc(mid, nums) > Calc(mid2, nums) {
			l = mid + 1
		} else {
			r = mid2
		}
	}

	fmt.Println(l)
	fmt.Println(Calc(l, nums))
}
