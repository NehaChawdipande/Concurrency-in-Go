package main
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(getNumber())
}

func getNumber() int {
	var i int 
	go func() {
		i = i+5
	}()
	time.Sleep(50* time.Millisecond) /* this statement is added to make goroutine g1: getnumber() longer than
	groutine g2:func(). If this sleep is removed the value of i returns to be 0 and not 5 since g2:func() is 
	now slower than g1*/
	return i
	
}
/* Thus we have a code for two goroutines g1: getnumber() and g2:func() 
now depending upon whichever goroutine is faster, there are two outcomes to the program,
 i = 0 or i= 5
 This is called as the race condition where the outcome of integer variable i is dependent upon
 the order of execution of the two goroutines g1 and g2*/