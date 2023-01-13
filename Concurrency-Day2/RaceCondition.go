/*
package main

import (

	"fmt"
	"sync"

)

	func numbInt(wg *sync.WaitGroup) {
		a = a + 2
		fmt.Println(a)
		wg.Done()
	}

	func main() {
		var wg sync.WaitGroup
		wg.Add(1)
		go numbInt(&wg)
		wg.Wait()
		fmt.Println("Finishing - In Main ")
	}
*/
package main

import (
	"fmt"
	"sync"
)

var a int

func add(wg *sync.WaitGroup) {
	fmt.Println("a := ", a)
	wg.Done()
}
func sum(wg *sync.WaitGroup) {
	a := 20
	a = a + 20
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go add(&wg)
	go sum(&wg)
	wg.Wait()
	fmt.Println("Race Condition")
}
