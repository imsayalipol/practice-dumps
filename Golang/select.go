package main
import(
	"fmt"
	"time"
)

func main(){
	start:=time.Now()

	c1:=make(chan string)
	c2:=make(chan string)

	go func(){
		time.Sleep(2*time.Second)
		c1<-"one"
	}()

	go func(){
		time.Sleep(1*time.Second)
		c2<-"two"
	}()

	for i:=0; i<2; i++ {
		select{
		case msg1:=<-c1:
			fmt.Println("msg 1 :", msg1)

		case msg2:=<-c2:
			fmt.Println("msg 2 :", msg2)
		}
	}
	
	duration:=time.Since(start)
	fmt.Println("Time : ", duration)
}