package main

import "fmt"

var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }

func main() {
	Add([]string{"apple", "banana"})
	Add([]string{"apple", "banana"})
	Add([]string{"orange"})

	fmt.Println(Count([]string{"apple", "banana"})) // Output: 2
	fmt.Println(Count([]string{"orange"}))          // Output: 1
	fmt.Println(Count([]string{"banana", "apple"})) // Output: 0 (different order -> different string key)

}
