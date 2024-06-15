package main

import (
	"fmt"
	"sync"
	"time"
)

//并发安全和锁
//有时候在go代码中可能会存在多个goroutine同时操作一个资源（临界区），这中情况会发生竞态问题（数据竞态）。类比现实生活中的例子有十字路口被各个方向的汽车竞争；还有火车上的卫生间被车厢里的人竞争。

//var x int64
//var wg1 sync.WaitGroup
//
//func add2() {
//	for i := 0; i < 5000; i++ {
//		x = x + 1
//	}
//	wg1.Done()
//}
//
//func main() {
//	wg1.Add(2)
//	go add2()
//	go add2()
//	wg1.Wait()
//	fmt.Println(x)
//}

//上面的代码中我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的时候就会存在数据竞争，导致最后的结果与期待的不符

// 互斥锁
// 互斥锁是一种常用的控制共享资源访问的方法，它能保证同时只有一个goroutine可以访问共享资源。Go语言中使用sync包的Mutex类型来实现互斥锁。使用互斥锁来修复上面代码的问题：
//var x int64
//var wg2 sync.WaitGroup //并发
//var lock sync.Mutex    //锁
//
//func add3() {
//	for i := 0; i < 5000; i++ {
//		lock.Lock() // 加锁
//		x = x + 1
//		lock.Unlock() // 解锁
//	}
//	wg2.Done()
//}
//
//func main() {
//	wg2.Add(2)
//	go add3()
//	go add3()
//	wg2.Wait()
//	fmt.Println(x)
//}

//使用互斥锁能够保证同一时间只有一个goroutine进入临界区，其他的goroutine则在等待锁；当互斥锁释放后，等待的goroutine才可以获取进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的。

// 读写互斥锁
// 互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没必要加锁的，这种场景下使用读写锁是更好的一种选择。读写锁在go语言中使用sync包中的RWMutex类型。

//读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。

// 读写锁示例：
var (
	x      int64
	wg2    sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write1() {
	//	lock.lock() 加互斥锁
	rwlock.Lock() //加写锁
	x = x + 1
	time.Sleep(10 * time.Microsecond) //假设读操作耗时10毫秒
	rwlock.Unlock()                   //解写锁
	//lock.Unlock() //解互斥锁
	wg.Done()
}

func read() {
	//	lock.Lock()  加互斥锁
	rwlock.RLock()               //加读锁
	time.Sleep(time.Microsecond) //假设读操作耗时1毫秒
	rwlock.RUnlock()             //解读锁
	//lock.Unlock() //解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write1()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

// 需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来
