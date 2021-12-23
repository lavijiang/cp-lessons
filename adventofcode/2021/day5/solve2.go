package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/boljen/go-bitmap"
)

var bitMaps []bitmap.Bitmap
var crossMaps []bitmap.Bitmap

func Init() {
	for i := 0; i < 1010; i++ {
		bitMaps = append(bitMaps, bitmap.New(1010))
		crossMaps = append(crossMaps, bitmap.New(1010))
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func markMap(x1, y1, x2, y2 int) {
	if x1 == x2 {
		y1, y2 = minInt(y1, y2), maxInt(y1, y2)
		for i := y1; i <= y2; i++ {
			if bitMaps[x1].Get(i) {
				crossMaps[x1].Set(i, true)
			} else {
				bitMaps[x1].Set(i, true)
			}
		}
		return
	}

	if y1 == y2 {
		x1, x2 = minInt(x1, x2), maxInt(x1, x2)
		for i := x1; i <= x2; i++ {
			if bitMaps[i].Get(y1) {
				crossMaps[i].Set(y1, true)
			} else {
				bitMaps[i].Set(y1, true)
			}
		}
		return
	}

	if x2+y1 == x1+y2 {
		x1, x2 = minInt(x1, x2), maxInt(x1, x2)
		y1 = minInt(y1, y2)
		for i := x1; i <= x2; i++ {
			if bitMaps[i].Get(y1 + i - x1) {
				crossMaps[i].Set(y1+i-x1, true)
			} else {
				bitMaps[i].Set(y1+i-x1, true)
			}
		}
		return
	}

	x1, x2 = minInt(x1, x2), maxInt(x1, x2)
	y1 = maxInt(y1, y2)
	for i := x1; i <= x2; i++ {
		if bitMaps[i].Get(y1 - i + x1) {
			crossMaps[i].Set(y1-i+x1, true)
		} else {
			bitMaps[i].Set(y1-i+x1, true)
		}
	}
}

func GetMark() int {
	tot := 0
	for i := 0; i < 1010; i++ {
		for j := 0; j < 1010; j++ {
			if crossMaps[i].Get(j) {
				tot++
			}
		}
	}
	return tot
}

func main() {
	Init()

	fi, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var x1, y1, x2, y2 int
	for {
		_, err := fmt.Fscanf(br, "%d,%d -> %d,%d\n", &x1, &y1, &x2, &y2)
		if err == io.EOF {
			break
		}
		if x1 != x2 && y1 != y2 && x1+y1 != x2+y2 && x1+y2 != x2+y1 {
			continue
		}
		markMap(x1, y1, x2, y2)
	}
	fmt.Println(GetMark())
}
