package main

import (
	"fmt"
	"runtime"
	"sync"
)

func print(wg *sync.WaitGroup, n int) {
	x := 0
	for i := 1; i < 10000000; i++ {
		x += 1
	}
	fmt.Println(n, x)
	//c <- true
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 管道方式
	//c := make(chan bool, 10)

	//
	wg := sync.WaitGroup{}
	//设置任务数
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go print(&wg, i)
	}

	//for i := 0; i < 10; i++ {
	//	<-c
	//}
	//等待任务完成
	wg.Wait()

	fmt.Println("Done.")
}
