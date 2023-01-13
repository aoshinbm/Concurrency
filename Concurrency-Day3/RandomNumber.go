package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNum(chanl chan int) {
	for i := 0; i < 15; i++ {
		/*Intn returns, as an int,
		  a non-negative pseudo-random number in the half-open interval [0,n) from the default Source.
		  It panics if n <= 0.
		*/
		chanl <- rand.Intn(100)
		fmt.Println("Integer produced ", rand.Intn(100))
	}
	close(chanl)
}
func main() {

	chann := make(chan int, 3)
	go randomNum(chann)

	time.Sleep(2 * time.Second)
	for v := range chann {
		fmt.Println("Integer number received ", v)
		time.Sleep(2 * time.Second)
	}
}
