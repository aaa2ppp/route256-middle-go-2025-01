package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func solve(matrix [][]byte) map[string]any {
	n, m := len(matrix), len(matrix[0])
	root := map[string]any{}
	searchSubBoxes(root, matrix, makeMatrix[int](n, m), 0, 0, n, m)
	return root
}

func searchSubBoxes(root map[string]any, matrix [][]byte, skeep [][]int, i1, j1, i2, j2 int) {
	for i := i1; i < i2; i++ {
		for j := j1; j < j2; {
			if skeep[i][j] != 0 {
				j += skeep[i][j]
				continue
			}

			if matrix[i][j] == '+' {
				box := map[string]any{}
				w := getBoxWidth(matrix, i, j+1)
				h := getBoxHeigh(matrix, i+1, j)
				name := getName(matrix, i+1, j+1)

				searchSubBoxes(box, matrix, skeep, i+1, j+1, i+1+h, j+1+w)
				if len(box) == 0 {
					root[name] = w * h
				} else {
					root[name] = box
				}

				setSkeep(matrix, skeep, i+1, j, w+2)
				j += w + 2
				continue
			}

			j++
		}
	}
}

func getBoxWidth(matrix [][]byte, i, j int) int {
	w := 0
	for matrix[i][j] != '+' {
		w++
		j++
	}
	return w
}

func getBoxHeigh(matrix [][]byte, i, j int) int {
	h := 0
	for matrix[i][j] != '+' {
		h++
		i++
	}
	return h
}

func getName(matrix [][]byte, i, j int) string {
	line := matrix[i][j:]
	k := 0
	for k < 3 {
		c := line[k]
		if !('0' <= c && c <= '9') && !('A' <= c && c <= 'Z') && !('a' <= c && c <= 'z') {
			break
		}
		k++
	}
	return string(line[:k])
}

func setSkeep(matrix [][]byte, skeep [][]int, i, j, s int) {
	for matrix[i][j] != '+' {
		skeep[i][j] = s
		i++
	}
	skeep[i][j] = s
}

func makeMatrix[T any](n, m int) [][]T {
	buf := make([]T, n*m)
	mtx := make([][]T, n)
	for i, j := 0, 0; i < n; i, j = i+1, j+m {
		mtx[i] = buf[j : j+m]
	}
	return mtx
}

func run(in io.Reader, out io.Writer) {
	br := bufio.NewReader(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()
	var t int

	if _, err := fmt.Fscanln(br, &t); err != nil {
		panic(err)
	}

	results := make([]map[string]any, 0, t)
	for i := 1; i <= t; i++ {
		var n, m int
		if _, err := fmt.Fscanln(br, &n, &m); err != nil {
			panic(err)
		}

		matrix := makeMatrix[byte](n, m)
		for i := 0; i < n; i++ {
			dst := matrix[i]
			for {
				buf, isPrefix, err := br.ReadLine()
				if err != nil {
					panic(err)
				}
				copy(dst, buf)
				dst = dst[len(buf):]
				if !isPrefix {
					break
				}
			}
		}

		ans := solve(matrix)
		results = append(results, ans)
	}

	json.NewEncoder(bw).Encode(results)
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}
