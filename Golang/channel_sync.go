package main
import (
	"fmt"
	"time"
)
//This is the function we’ll run in a goroutine. 
// The done channel will be used to notify another goroutine that this function’s work is done.
func f(c chan bool){
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	c <- true
}
func main(){
	c :=make(chan bool, 1)
	go f(c)
	<-c		//Block until we receive a notification from the worker on the channel.
}