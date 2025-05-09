package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	// create an out channel that will be returned
	out := make(chan int)

	// create a goroutine to push all the nums into the out channel
	go func() {
		for _, num := range nums {
			// push the num into the channel
			out <- num
		}
		close(out)
	}()

	return out
}

func squareNumsChannel(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for chItem := range in {
			// square the value and then push the item into the channel
			out <- chItem * chItem
		}
		close(out)
	}()

	return out
}

func main() {
	// a pipeline that takes a slice of numbers and returns the the square of each number
	nums := []int{2, 4, 7, 9, 5, 15}

	// first stage of the pipeline
	numsChannel := sliceToChannel(nums)

	// second state of the pipeline
	squaredChannel := squareNumsChannel(numsChannel)

	// final stage, print all the squared nums!
	for num := range squaredChannel {
		fmt.Println(num)
	}

}
