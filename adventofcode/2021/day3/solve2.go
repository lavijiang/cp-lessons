package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func getRes(s string) int64 {
	var res int64
	for _, c := range s {
		res <<= 1
		if c == '1' {
			res++
		}
	}
	return res
}

func main() {
	fi, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	var s string
	var gamma, epsilon int64
	var contt, cont, cont0, cont1 []string
	br := bufio.NewReader(fi)
	for {
		_, err := fmt.Fscanln(br, &s)
		if err == io.EOF {
			break
		}
		cont = append(cont, s)
		contt = append(contt, s)
	}

	var st, tot int
	for {
		if len(cont) == 1 {
			break
		}
		for _, s := range cont {
			if s[st] == '1' {
				tot++
				cont1 = append(cont1, s)
			} else {
				cont0 = append(cont0, s)
			}
		}
		if tot*2 >= len(cont) {
			cont = make([]string, len(cont1))
			copy(cont, cont1)
		} else {
			cont = make([]string, len(cont0))
			copy(cont, cont0)
		}
		st++
		tot = 0
		cont0 = make([]string, 0)
		cont1 = make([]string, 0)
	}

	gamma = getRes(cont[0])
	fmt.Println(cont[0])
	fmt.Println(gamma)

	st, tot = 0, 0
	for {
		if len(contt) == 1 {
			break
		}
		for _, s := range contt {
			if s[st] == '1' {
				tot++
				cont1 = append(cont1, s)
			} else {
				cont0 = append(cont0, s)
			}
		}
		if tot*2 >= len(contt) {
			contt = make([]string, len(cont0))
			copy(contt, cont0)
		} else {
			contt = make([]string, len(cont1))
			copy(contt, cont1)
		}
		st++
		tot = 0
		cont0 = make([]string, 0)
		cont1 = make([]string, 0)
	}

	epsilon = getRes(contt[0])
	fmt.Println(contt[0])
	fmt.Println(epsilon * gamma)
}
