package main

import (
	"bufio"
	"fmt"
	. "fmt"
	"io"
	"os"
)

func main() {
	fi, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	var s string
	var r, gamma, epsilon int64
	var tot1 [13]int64
	br := bufio.NewReader(fi)
	for {
		_, err := Fscanln(br, &s)
		if err == io.EOF {
			break
		}
		r++
		for index, c := range s {
			if c == '1' {
				tot1[index]++
			}
		}
	}

	for i := 0; i < len(s); i++ {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if tot1[i]*2 >= r {
			gamma++
		} else {
			epsilon++
		}
	}

	Println(gamma * epsilon)
}
