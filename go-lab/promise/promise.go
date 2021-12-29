package promise

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func PromiseAll() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	c := make(chan int)
	quit := make(chan int)

	wg := new(sync.WaitGroup)
	go func() {
		wg.Wait()
		close(c)
		close(quit)
	}()

	for k := range make([]int64, 5) {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rt := r.Int63n(10)
			time.Sleep(time.Duration(rt) * time.Second)
			fmt.Printf("sleep %v second, index is %v \n", rt, i)
			if false {
				quit <- i
			} else {
				c <- i
			}
		}(k)
	}

	for {
		select {
		case res, ok := <-c:
			fmt.Printf("res is %v ok is %v\n", res, ok)
		case <-quit:
			fmt.Println("exit")
		default:
			break
		}
	}

}
