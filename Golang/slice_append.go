package main
import "fmt"

func appendInt(x []int, y int){
	var z []int
	zlen:=len(x)+1
	if zlen <= cap(x){
		z=x[:zlen]
	} else{
		zcap := zlen
		if zcap < 2*len(x){
			zcap=2*len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)	
	}
	z[len(x)] = y
	fmt.Println(z)
}

func main() {
	x := []int{1, 2, 3}
	appendInt(x, 10)
}