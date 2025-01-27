package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type item struct {
	x, y int
	d    byte
}

func solve(n, m int) []item {
	if n == 1 {
		return []item{{1, 1, 'R'}}
	}
	if m == 1 {
		return []item{{1, 1, 'D'}}
	}
	if n < m {
		return []item{{1, 1, 'R'}, {n, m, 'L'}}
	}
	return []item{{1, 1, 'D'}, {n, m, 'U'}}
}

func run(in io.Reader, out io.Writer) {
	br := bufio.NewReader(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()
	var t int

	if _, err := fmt.Fscanln(br, &t); err != nil {
		panic(err)
	}

	for i := 1; i <= t; i++ {
		var n, m int
		if _, err := fmt.Fscan(br, &n, &m); err != nil {
			panic(err)
		}

		ans := solve(n, m)
		fmt.Fprintln(bw, len(ans))
		for _, v := range ans {
			fmt.Fprintf(bw, "%d %d %c\n", v.x, v.y, v.d)
		}
	}
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}
