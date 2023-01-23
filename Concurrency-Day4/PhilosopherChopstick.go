/*
Solution :
Correctness properties it needs to satisfy are :

Mutual Exclusion Principle –

	No two Philosophers can have the two forks simultaneously.

Free from Deadlock –

	Each philosopher can get the chance to eat in a certain finite time.

Free from Starvation –When few Philosophers are waiting then one gets a chance to eat in a while.

	No strict Alternation.
	Proper utilization of time.
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// chopsticks struct created
type ChopS struct {
	sync.Mutex
}

// philosopher struct created where referencing
// chopstick datatype for right n left chopsticks
type Philosopher struct {
	leftChop, rightChop *ChopS
}

// built-in func in golang
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// wen philosopher gets both the chopsticks he proceeds to eat
// i.e wen he gets both resources he executes eat() function and releses after certain period
func (p Philosopher) eat() {
	for i := 0; i < 3; i++ {
		p.leftChop.Lock()
		p.rightChop.Lock()
		fmt.Println("Philospher", i, "Eating")
		time.Sleep(time.Second)
		p.leftChop.Unlock()
		p.rightChop.Unlock()
		fmt.Println("Philospher", i, "finished eating")
		time.Sleep(time.Second)
		wg.Done()

	}
}

var wg sync.WaitGroup

func main() {

	wg.Add(1)

	//5 chopsticks present
	ChopSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		ChopSticks[i] = new(ChopS)

	}

	//5 philosophers
	philos := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philosopher{ChopSticks[i], ChopSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}

	wg.Wait()

}
