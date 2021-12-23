package main

import (
	"bufio"
	"fmt"
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

	br := bufio.NewReader(fi)
	var s string
	_, _ = fmt.Fscanln(br, &s)
	numsStrList := strings.Split(s, ",")
	nums := []int{}
	for _, numStr := range numsStrList {
		num, _ := strconv.Atoi(numStr)
		nums = append(nums, num)
	}

	var matrixs [][5][5]int
	var matrixsTag [][5][5]int
	var matrixsRow [][5]int
	var matrixsCol [][5]int
	for {
		_, err := fmt.Fscanln(br, &s)
		if err == io.EOF {
			break
		}
		var matrix [5][5]int
		for i := 0; i < 5; i++ {
			_, err = fmt.Fscanln(br, &matrix[i][0], &matrix[i][1], &matrix[i][2], &matrix[i][3], &matrix[i][4])
			if err == io.EOF {
				break
			}
		}
		var matrixRow [5]int
		matrixs = append(matrixs, matrix)
		matrixsRow = append(matrixsRow, matrixRow)
		matrixsCol = append(matrixsCol, matrixRow)
		var matrixTag [5][5]int
		matrixsTag = append(matrixsTag, matrixTag)
	}

	var tagNum, ind, sum int
	for _, num := range nums {
		if tagNum > 0 {
			break
		}
		for index, matrix := range matrixs {
			if tagNum > 0 {
				break
			}
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if matrix[i][j] == num && matrixsTag[index][i][j] == 0 {
						matrixsTag[index][i][j] = 1
						matrixsRow[index][i]++
						matrixsCol[index][j]++
						if matrixsRow[index][i] >= 5 || matrixsCol[index][j] >= 5 {
							tagNum = num
							ind = index
							break
						}
					}
				}
			}
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrixsTag[ind][i][j] == 0 {
				sum += matrixs[ind][i][j]
			}
		}
	}
	fmt.Println(sum * tagNum)
}
