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

	var pre int64 = -1
	var no int64
	var res int64

	br := bufio.NewReader(fi)
	for {
		_, c := Fscanln(br, &no)
		if c == io.EOF {
			break
		}
		if pre != -1 && no > pre {
			res++
		}
		pre = no
	}

	Println(res)
}
