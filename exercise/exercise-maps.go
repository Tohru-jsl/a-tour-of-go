package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sf := strings.Fields(s)
	v := make(map[string]int)
	for i := 0; i < len(sf); i++ {
		v[sf[i]] = v[sf[i]] + 1
	}
	return v
}

func main() {
	wc.Test(WordCount)
}
