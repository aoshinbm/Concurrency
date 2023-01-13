package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
)

var wg sync.WaitGroup

func sortSlice1(slice1 []float64, channel chan int) {
	fmt.Println("Before sorting:", slice1)
	sort.Float64s(slice1)
	fmt.Println("After sorting:", slice1)
	wg.Done()
}
func sortSlice2(slice2 []float64, channel chan int) {
	fmt.Println("Before sorting:", slice2)
	sort.Float64s(slice2)
	fmt.Println("After sorting:", slice2)
	wg.Done()
}
func sortSlice3(slice3 []float64, channel chan int) {
	fmt.Println("Before sorting:", slice3)
	sort.Float64s(slice3)
	fmt.Println("After sorting:", slice3)
	wg.Done()
}
func sortSlice4(slice4 []float64, channel chan int) {
	fmt.Println("Before sorting:", slice4)
	sort.Float64s(slice4)
	fmt.Println("After sorting:", slice4)
	wg.Done()
}

func main() {

	chanel := make(chan int)
	floatarray := []float64{109.32, 19.21, 2.93, 49.23, 4.32, 49.11, 33.9, 1.9, 52.5, 81.0, 23.65}
	fmt.Println(len(floatarray))

	chunks := len(floatarray) / 4
	fmt.Println("Dividing array into 4 equal parts and ",
		"how many elements each part will contain :", math.Round(float64(chunks)))

	slice1 := floatarray[:len(floatarray)/4]
	fmt.Println(slice1)
	slice2 := floatarray[len(floatarray)/4 : 2*len(floatarray)/4]
	fmt.Println(slice2)
	slice3 := floatarray[2*len(floatarray)/4 : 3*len(floatarray)/4]
	fmt.Println(slice3)
	slice4 := floatarray[3*len(floatarray)/4 : 4*len(floatarray)/4]
	fmt.Println(slice4)

	go sortSlice1(slice1, chanel)
	go sortSlice2(slice2, chanel)
	go sortSlice3(slice3, chanel)
	go sortSlice4(slice4, chanel)

	<-chanel
	<-chanel
	<-chanel
	<-chanel

	fmt.Println("Sorted Array : ")

	finalfloatArray := make([]float64, 0)
	//... expansion syntax
	finalfloatArray = append(finalfloatArray, slice1...)
	finalfloatArray = append(finalfloatArray, slice2...)
	finalfloatArray = append(finalfloatArray, slice3...)
	finalfloatArray = append(finalfloatArray, slice4...)
	fmt.Println(finalfloatArray)
}
