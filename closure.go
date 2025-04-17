package main

import "fmt"

func count() func() int{
	a := 0
	return func()int{
		 a+=1
		 return a
	}
}

func main(){

	counter:=count()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	fmt.Println()
	
	new_count:=count()
	fmt.Println(new_count())
	fmt.Println(new_count())
	fmt.Println(new_count())
	fmt.Println(new_count())
	fmt.Println(new_count())
	
}