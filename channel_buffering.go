package main
import "fmt"

func main(){

	msg := make(chan string, 2)
	msg <- "hello"
	msg <-"all"

	fmt.Println(<-msg)
	fmt.Println(<-msg)
}