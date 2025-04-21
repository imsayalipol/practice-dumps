package main
import "fmt"

func main() {
	x := make([]int, 3, 5)
	fmt.Println("x:", x, "len:", len(x), "cap:", cap(x))

	z := x[:4] // zlen = 4 <= cap(x), so it's OK
	fmt.Println("z:", z, "len:", len(z), "cap:", cap(z))
}
