package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(3*time.Second)
	//<-timer1.C

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	//time.Sleep(2 * time.Second)
	 <-timer1.C
	fmt.Println("Exiting")
}
