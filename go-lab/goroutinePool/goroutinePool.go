package goroutinepool

import (
	"log"
	"time"
)

func pool() {

	pool := make(chan int, 5)

	arrTest := make([]int, 0)

	for i := 0; i < 100; i++ {
		pool <- 1
		go func(i int) {
			arrTest = append(arrTest, i)
			time.Sleep(2 * 1e9)
			<-pool
		}(i)
	}

	log.Printf("%+v\n", arrTest)

}
