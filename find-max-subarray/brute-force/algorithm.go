package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"time"
	"os"
	"strconv"
	"sync"
	"math"
)
func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func generateRandomArray(length int) []int {
	array := make([]int, length, length)
	for i:=0; i < length; i++ {
		array[i] = random(-1000, 1000)
	}
	return array
}
func algorithm(array []int) (int,int,int) {
	bestStartIndex := 0
	bestEndIndex := 0
	bestSum := 0
	for i:=0; i < len(array); i++ {
		currentSum := array[i]
		for j:=i+1; j < len(array); j++ {
			currentSum += array[j]
			if currentSum > bestSum {
				bestSum = currentSum
				bestStartIndex = i
				bestEndIndex = j
			}
		}
	}
	return bestStartIndex, bestEndIndex, bestSum
}
func timing(array []int, result chan int, wg *sync.WaitGroup) {
	start := time.Now()
	algorithm(array)
	result <- int(time.Since(start).Nanoseconds())
	wg.Done()
}
func makeTest(length int) float64 {
	array := generateRandomArray(length)
	resultChannel := make(chan int, 4)
	wg := sync.WaitGroup{}
	for i:=0; i < 4; i++ {
		wg.Add(1)
		go timing(array, resultChannel, &wg)
	}
	wg.Wait()
	close(resultChannel)
	min := math.MaxInt32
	for resultTiming := range resultChannel {
		if resultTiming < min {
			min = resultTiming
		}
	}
	return float64(min)
}
func main() {

	testArray := []int{13,-3,-25,20,-3,-16,-23,18,20,-7,12,-5,-22,15,-4,7}
	bestStartIndex, bestEndIndex, bestSum := algorithm(testArray)
	if bestStartIndex != 7 {
		panic("fail!")
	}
	if bestEndIndex != 10 {
		panic("fail!")
	}
	if bestSum != 43 {
		panic("fail!")
	}

	max := 8000
	f, err := os.OpenFile("algorithm.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for i := 0; i < max; i++ {
		result := makeTest(i)
		if _,err := f.WriteString(strconv.Itoa(i) + ","+ strconv.FormatFloat(result, 'f', 6,64) +"\n"); err != nil {
			panic(err)
		}
		clearScreen()
		fmt.Println("Progress: " + strconv.Itoa(i) + "/" + strconv.Itoa(max))
	}

}
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}