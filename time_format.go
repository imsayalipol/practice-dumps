package main
import (
	"fmt"
	"time"
)

func main(){

	t:= time.Date(1995,time.September,22,13,0,0,0,time.UTC)
	//FormattedTime:=t.Format("Monday, 01/02/2006, 15:04")
	newFormat := t.Format("January 2, 2006, 3:04pm")
	fmt.Println(newFormat)
}