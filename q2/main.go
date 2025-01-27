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
	prices := make(map[int]int, len(items))
	names := make(map[string]int, len(items))

	for _, it := range items {
		prices[it.price] = 0
		names[it.name] = it.price
	}

	for _, it := range strings.Split(s, ",") {
		tok := strings.Split(it, ":")
		if len(tok) != 2 {
			// log.Printf("oops! %q", tok)
			return false
		}

		// check name
		name := tok[0]
		wantPrice, ok := names[name]
		if !ok {
			// log.Printf("oops! double name: %q v:%d ok:%v", name, v, ok)
			return false
		}

		// check price
		if tok[1] == "" || tok[1][0] == '0' {
			// log.Printf("oops! price: %q", tok[1])
			return false
		}
		price, err := strconv.Atoi(tok[1])
		if err != nil {
			// log.Printf("oops! price: %q", tok[1])
			return false
		}

		if price != wantPrice {
			return false
		}

		if v, ok := prices[price]; !ok || v != 0 {
			// log.Printf("oops! double price: %d v:%d ok:%v", price, v, ok)
			return false
		}
		prices[price]++
	}

	for price, v := range prices {
		_ = price
		if v != 1 {
			// log.Printf("oops! price not found: %v", price)
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
