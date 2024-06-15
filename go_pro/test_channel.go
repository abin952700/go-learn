package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	//c := make(chan int)
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		c <- i
	//	}
	//	close(c)
	//}()
	//for {
	//	if data, ok := <-c; ok {
	//		fmt.Println(data)
	//	} else {
	//		break
	//	}
	//}
	//fmt.Println("main结束")

	// 当通过通道发送有限的数据时，我们可以通过close函数关闭通道来告知从该通道接收值的goroutine停止等待。当通道被关闭时，往该通道发送值会引发panic,从该通道里的接收的值一直都是类型零值
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//// 开启goroutine将0~100的数发送到ch1中
	//go func() {
	//	for i := 0; i < 100; i++ {
	//		ch1 <- i
	//	}
	//	close(ch1)
	//}()
	//// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	//go func() {
	//	for {
	//		i, ok := <-ch1 // 通道关闭后再取值ok=false
	//		if !ok {
	//			break
	//		}
	//		ch2 <- i * i
	//	}
	//	close(ch2)
	//}()
	//// 在主goroutine中从ch2中接收值打印
	//for i := range ch2 { // 通道关闭后会退出for range循环
	//	fmt.Println(i)
	//}

	//有的时候我们会将通道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用通道都会对其进行限制，比如限制通道在函数中只能发送或接收。
	// 1.chan<- int是一个只能发送的通道，可以发送但是不能接收；
	// 2.<-chan int是一个只能接收的通道，可以接收但是不能发送。
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
