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

	var no int64
	var order string
	var hor, dep, aim int64

	br := bufio.NewReader(fi)
	for {
		_, c := Fscanln(br, &order, &no)
		if c == io.EOF {
			break
		}
		if order == "forward" {
			hor += no
			dep += no * aim
		} else if order == "down" {
			aim += no
		} else {
			aim -= no
		}
	}
	Println(hor * dep)
}
