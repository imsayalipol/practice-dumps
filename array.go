package main
import "fmt"

func main(){

	b := [5]int{1, 2, 3, 4, 5}

	for i:=0; i<5;i++ {
		fmt.Println(b[i])
	}
}