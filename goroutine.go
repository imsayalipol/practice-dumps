package main
import ( 
	"fmt"
	"time"
)

func f(name string){
	for i:=0; i<3; i++ {
		fmt.Println(name,"=", i)
	}
}
func main(){
 go f("Golang")

 go func(){ 
	for i:=0;i<2;i++{
		fmt.Println("Hello all")
		}
	}()
time.Sleep(time.Second)
}