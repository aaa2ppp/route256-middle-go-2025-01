package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unsafe"
)

func solve(ss []string) int {
	even := make(map[string]int, len(ss))
	odd := make(map[string]int, len(ss))
	all := make(map[string]int, len(ss))

	var count int

	for _, s := range ss {
		e := extractEvenChars(s)
		o := extractOddChars(s)

		count += even[e]
		even[e]++

		if o == "" {
			continue
		}

		odd[o]++
		count += odd[o]

		all[s]++
		count -= all[s]
}

	return count
}

func extractEvenChars(s string) string {
	b := make([]byte, (len(s)+1)/2)
	for i, j := 0, 0; i < len(s); i, j = i+2, j+1 {
		b[j] = s[i]
	}
	return unsafeString(b)
}

func extractOddChars(s string) string {
	b := make([]byte, len(s)/2)
	for i, j := 1, 0; i < len(s); i, j = i+2, j+1 {
		b[j] = s[i]
	}
	return unsafeString(b)
}

func unsafeString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
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
			ss[j] = strings.TrimRight(s, " \t\r\n")
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
