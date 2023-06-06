package main

import (
	"fmt"
	"time"
)

/*
统计1~8000的数字中，哪些是素数？
1、写入8000个数
2、开4个协程读取
*/

func isPrime(in int) (isPrime bool) {

	isPrime = true

	if in > 1 {
		for i := 2; i < in; i++ {
			if in%i == 0 {
				isPrime = false
				break
			}
		}
	}

	return isPrime
}

func main() {
	total := 8000
	gorountin := 6
	intChan := make(chan int, 1000)
	primChan := make(chan int, 2000)
	exitChan := make(chan bool, 4)

	go func(intChain chan int) {
		for i := 1; i < total; i++ {
			intChain <- i
		}
		close(intChain)
		fmt.Printf("intChain 共初试化了%d个数\n", total)
	}(intChan)

	start := time.Now()

	for i := 0; i < gorountin; i++ {
		go func(intChan chan int, primChan chan int, exitChan chan bool) {
			for {
				v, ok := <-intChan
				if !ok {
					//取完了，往exitChan里塞入一个成功
					exitChan <- true
					break
				}
				if isPrime(v) {
					fmt.Printf("成功找到一个素数：%d\n", v)
					primChan <- v
				}
			}

			// close(exitChan)
		}(intChan, primChan, exitChan)
	}

	// exitChan没有clos，但也没有报错.
	for i := 1; i <= gorountin; i++ {
		<-exitChan
		// fmt.Printf("管道%d成功结束%v\n", i, v)
	}
	close(primChan)

	// go func() {
	// label:
	// 	for {
	// 		select {
	// 		case v := <-exitChan:
	// 			fmt.Printf("成功关闭一个管道%v", v)
	// 		default:
	// 			fmt.Println("end!~~~")
	// 			break label
	// 		}
	// 	}

	// 	close(primChan)
	// }()
	// for {
	// 	_, ok := <-primChan
	// 	if !ok {
	// 		break
	// 	}
	// }

	end := time.Now().UnixNano() / 1e6
	diff := time.Since(start)

	fmt.Printf("共找到%d个素数,耗时%v毫秒,%v\n", len(primChan), diff.Milliseconds(), end-(start.UnixNano()/1e6))

}
