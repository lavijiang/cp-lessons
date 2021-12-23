package main

import (
	"bufio"
	"fmt"
	. "fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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
	nums := [2][]int{}
	for _, numStr := range numStrList {
		num, _ := strconv.Atoi(numStr)
		nums[0] = append(nums[0], num)
	}

	var res [257]int64
	res[0] = int64(len(nums[0]))
	for i := 1; i < 9; i++ {
		nums[i&1] = []int{}
		for _, num := range nums[1-i&1] {
			if num == 0 {
				nums[i&1] = append(nums[i&1], 8, 6)
			} else {
				nums[i&1] = append(nums[i&1], num-1)
			}
		}
		res[i] = int64(len(nums[i&1]))
	}

	for i := 9; i <= 256; i++ {
		res[i] = res[i-9] + res[i-7] //找规律
	}
	fmt.Println(res[256])
}
