package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func student(sum *int64, id int) {

	//random waiting time for each student
	var timeout time.Duration
	r := rand.Intn(1)

	timeout = time.Duration(r) * time.Millisecond
	time.Sleep(timeout)

	//Rating from each student
	rating := int64(rand.Intn(100))
	atomic.AddInt64(sum, rating)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//rand.Intn(10)
	//waitgroup , sum of all rating
	var wg sync.WaitGroup
	var sum int64
	sum = 0
	for i := 1; i <= 200; i++ {
		wg.Add(1)
		x := i

		go func() {
			defer wg.Done()
			student(&sum, x)
		}()
	}
	wg.Wait()

	//average of rating
	fmt.Println(sum / 200)
}
