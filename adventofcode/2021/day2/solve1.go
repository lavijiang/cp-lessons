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
	var hor, dep int64

	br := bufio.NewReader(fi)
	for {
		_, c := Fscanln(br, &order, &no)
		if c == io.EOF {
			break
		}
		//Println(order, no)
		if order == "forward" {
			hor += no
		} else if order == "down" {
			dep += no
		} else {
			dep -= no
		}
	}
	Println(hor * dep)
}
