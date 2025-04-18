package main
import (
	"fmt"
	"time"
)

func f(c chan bool){
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	c <- true
}
func main(){
	c :=make(chan bool, 1)
	go f(c)
	<-c
}