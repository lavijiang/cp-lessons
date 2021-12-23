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

	var a [3]int64
	var no, res int64
	br := bufio.NewReader(fi)
	for i := 0; ; i++ {
		_, err := Fscanln(br, &no)
		if err == io.EOF {
			break
		}
		if i < 3 {
			a[i] = no
		} else {
			if no > a[0] {
				res++
			}
			a[0] = a[1]
			a[1] = a[2]
			a[2] = no
		}
	}

	Println(res)
}
