package main

import (
	"fmt"
)

func nonempty(st []string) []string {
	i := 0
	for _, s := range st {
		if s != "" {
			st[i] = s
			i++
		}
		//fmt.Println("i=", i, "s=", s,"\n")
	}
	return st[:i]
}

func main() {

	s := []string{"one", "", "three"}

	fmt.Println(nonempty(s))
	fmt.Println(s)
}
