package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

var wg sync.WaitGroup

func sortSlice1(slice1 []float64, wg *sync.WaitGroup) {
	fmt.Println("Before sorting:", slice1)
	sort.Float64s(slice1)
	fmt.Println("After sorting:", slice1)
	wg.Done()
}
func sortSlice2(slice2 []float64, wg *sync.WaitGroup) {
	fmt.Println("Before sorting:", slice2)
	sort.Float64s(slice2)
	fmt.Println("After sorting:", slice2)
	wg.Done()
}
func sortSlice3(slice3 []float64, wg *sync.WaitGroup) {
	fmt.Println("Before sorting:", slice3)
	sort.Float64s(slice3)
	fmt.Println("After sorting:", slice3)
	wg.Done()
}
func sortSlice4(slice4 []float64, wg *sync.WaitGroup) {
	fmt.Println("Before sorting:", slice4)
	sort.Float64s(slice4)
	fmt.Println("After sorting:", slice4)
	wg.Done()
}

func main() {
	wg.Add(4)
	floatarray := []float64{109.32, 19.21, 2.93, 49.23, 4.32, 49.11, 33.9, 1.9, 52.5, 81.0, 23.65}
	fmt.Println(len(floatarray))

	chunks := len(floatarray) / 4
	fmt.Println("Dividing array into 4 equal parts and ",
		"how many elements each part will contain :", math.Round(float64(chunks)))

	slice1 := floatarray[:11/4]
	fmt.Println(slice1)
	go sortSlice1(slice1, &wg)

	slice2 := floatarray[2 : 9/4]
	fmt.Println(slice2)
	go sortSlice2(slice2, &wg)

	wg.Wait()
	/*slice1 := make([]int, 0)

	n := len(floatarray) / 4

	for i := 0; i < 4; i++ {

		min := i * n
		max := (i + 1) * n

		slice1 = append(slice1, floatarray[min:max])

	}

	go sortSlice3(slice3, &wg)
	go sortSlice4(slice4, &wg)


	fmt.Println("Sorted Array : ")

	*/
}
