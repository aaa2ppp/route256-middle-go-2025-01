package main

import (
	"bufio"
	"fmt"
	"hash/maphash"
	"io"
	"os"
)

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

		var evenHash, oddHash maphash.Hash
		even := make(map[uint64]int, n)
		odd := make(map[uint64]int, n)
		all := make(map[uint64]int, n)
		count := 0

		for j := 0; j < n; j++ {
			evenHash.Reset()
			oddHash.Reset()
			var isOdd, oddExists bool

			for {
				line, isPrefix, err := br.ReadLine()
				if err != nil {
					panic(err)
				}

				for _, c := range line {
					if isOdd {
						isOdd = false
						oddExists = true
						oddHash.WriteByte(c)
					} else {
						isOdd = true
						evenHash.WriteByte(c)
					}
				}

				if !isPrefix {
					break
				}
			}

			evenSum := evenHash.Sum64()
			count += int(even[evenSum])
			even[evenSum]++

			if oddExists {
				sum := oddHash.Sum64()
				count += int(odd[sum])
				odd[sum]++

				sum ^= evenSum
				count -= int(all[sum])
				all[sum]++
			}
		}

		fmt.Fprintln(bw, count)
	}
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}
