package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type item struct {
	name  string
	price int
}

func solve(items []item, s string) bool {
	prices := make(map[int]bool, len(items))
	names := make(map[string]int, len(items))

	for _, it := range items {
		prices[it.price] = false
		names[it.name] = it.price
	}

	for _, it := range strings.Split(s, ",") {
		tok := strings.Split(it, ":")
		if len(tok) != 2 {
			return false
		}

		name, priceA := tok[0], tok[1]

		if name == "" || len(name) > 10 {
			return false
		}
		if priceA == "" || priceA[0] == '0' {
			return false
		}

		price, err := strconv.Atoi(priceA)
		if err != nil {
			return false
		}
		if used, ok := prices[price]; !ok || used {
			return false
		}
		if wantPrice, ok := names[name]; !ok || price != wantPrice {
			return false
		}
		prices[price] = true
	}

	for _, used := range prices {
		if !used {
			return false
		}
	}

	return true
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

		items := make([]item, n)
		for j := 0; j < n; j++ {
			if _, err := fmt.Fscanln(br, &items[j].name, &items[j].price); err != nil {
				panic(fmt.Sprintf("%d.%d: %v", i, j, err))
			}
		}

		s, err := br.ReadString('\n')
		if err != nil {
			panic(fmt.Sprintf("%d.s: %v", i, err))
		}
		s = strings.TrimSuffix(s, "\n")

		if debugEnable {
			log.Println(i, "items:", items, "s:", s)
		}

		ans := solve(items, s)
		if ans {
			fmt.Fprintln(bw, "YES")
		} else {
			fmt.Fprintln(bw, "NO")
		}
	}
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}
