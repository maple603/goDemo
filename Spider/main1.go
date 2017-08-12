package main

import (
	"fmt"
	"runtime"
)

func print(c chan bool, n int) {
	x := 0
	for i := 1; i < 10000000; i++ {
		x += 1
	}
	fmt.Println(n, x)
	c <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 管道方式 有缓冲
	c := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go print(c, i)
	}

	for i := 0; i < 10; i++ {
		<-c
	}


	fmt.Println("Done.")
}
