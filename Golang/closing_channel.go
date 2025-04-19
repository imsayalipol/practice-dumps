package main
import "fmt"

func main(){
	jobs :=make(chan int,5)
	done :=make(chan bool)

	go func (){
		for{
			j, ok := <-jobs
			if ok{
				fmt.Println("Job received :", j)
			}else{
				fmt.Println("No more jobs")
				done<-true
				return
			}			
		}
	}()

	for j:=1; j<4; j++{
		jobs<-j
		fmt.Println("Job sent :", j)
	}
	
	// ****** VIMP ******
	close(jobs)
	
	fmt.Println("All jobs sent !!!")

	<-done

	_, ok := <-jobs
	fmt.Println("More jobs available? : ", ok)
}