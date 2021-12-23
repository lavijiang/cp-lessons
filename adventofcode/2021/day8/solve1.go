package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fi, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	res := 0
	sc := bufio.NewScanner(fi)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for sc.Scan() {
		s := sc.Text()
		ss := strings.Split(s, "|")
		if len(ss) < 2 {
			continue
		}
		ss1 := strings.Split(ss[1], " ")
		for _, s := range ss1 {
			if len(s) == 2 || len(s) == 3 || len(s) == 4 || len(s) == 7 {
				res++
			}
		}
	}

	fmt.Println(res)
}
