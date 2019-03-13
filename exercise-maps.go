package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	counter := make(map[string]int)
	sf := strings.Fields(s)
	for _, v := range sf {
		counter[v] = counter[v] + 1
	}
	return counter
}

func main() {
	c := WordCount("I am a dog I am your dog I am cool")
	fmt.Println(c)
}
