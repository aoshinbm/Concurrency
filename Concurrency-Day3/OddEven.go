package main

import (
	"fmt"
)

func even(numbers []int, chnll chan bool) {
	fmt.Println("Even numbers:")
	for _, eve := range numbers {
		if eve%2 == 0 {
			fmt.Println(eve)
		}
	}
	chnll <- true
}

func odd(numbers []int, chnll chan bool) {
	fmt.Println("Odd numbers:")
	for _, od := range numbers {
		if od%2 != 0 {
			fmt.Println(od)
		}
	}
	chnll <- true
}

func main() {
	chnel := make(chan bool)
	numbers := []int{22, 35, 64, 1, 29, 48, 56, 19, 10, 67}
	fmt.Println(numbers)
	fmt.Println("OddEven Program")
	go odd(numbers, chnel)
	go even(numbers, chnel)
	<-chnel
	<-chnel
}
