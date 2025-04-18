package main
import (
	"fmt"
	"errors"
)

type argError struct{
	args int
	msg string
} 

func (e *argError) Error() string{
	return fmt.Sprintf("%d - %s", e.args, e.msg)
}

func f(arg int)(int, error){
	if arg==42 {
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg+3, nil
}

func main(){
s := [2]int{2, 42} 
for i:=0; i<2; i++ {
	i, err := f(s[i])
	var ae *argError
	if errors.As(err, &ae){
		fmt.Println(ae.args)
		fmt.Println(ae.msg)
	}else{
		fmt.Println("No error matched, found : ", i)
	}
}
}