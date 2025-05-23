package main

import "fmt"

func main() {
    messages := make(chan string,1)
    signals := make(chan bool,1)
    
	messages <-"Hii all"
	//fmt.Println(<-messages)

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    
    }

    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    
    }

    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}