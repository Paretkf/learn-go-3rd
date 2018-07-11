package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	a := strings.Fields(s)
	b := make(map[string]int)
	for i := 0; i < len(a); i++ {
		_, ok := b[string(a[i])]
		if ok {
			b[string(a[i])] = b[string(a[i])] + 1
		} else {
			b[string(a[i])] = 1
		}
	}
	return b
}

func main() {
	wc.Test(WordCount)
}
