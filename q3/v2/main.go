package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func solve(ss []string) int {
	even := make(map[string]int, len(ss))
	odd := make(map[string]int, len(ss))
	all := make(map[string]int, len(ss))
	count := 0

	for _, s := range ss {
		var e, o string
		{
			var sb strings.Builder
			sb.Grow(len(s)/2 + 1)
			for i := 0; i < len(s); i += 2 {
				sb.WriteByte(s[i])
			}
			e = sb.String()
		}
		{
			var sb strings.Builder
			sb.Grow(len(s)/2 + 1)
			for i := 1; i < len(s); i += 2 {
				sb.WriteByte(s[i])
			}
			o = sb.String()
		}

		count += even[e]
		if o != "" {
			count += odd[o]
			count -= all[s]
		}

		even[e]++
		odd[o]++
		all[s]++
	}

	return count
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
		var n int
		if _, err := fmt.Fscanln(br, &n); err != nil {
			panic(err)
		}

		ss := make([]string, n)
		for j := range ss {
			s, err := br.ReadString('\n')
			if err != nil {
				panic(fmt.Sprintf("%d.%d: %v", i, j, err))
			}
			ss[j] = strings.TrimSpace(s)
		}

		ans := solve(ss)
		fmt.Fprintln(bw, ans)
	}
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}
