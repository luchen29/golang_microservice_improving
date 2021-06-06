package practice

import (
	"fmt"
	"sync"
	"time"
)

type sig struct {}

type f func() error

type Pool struct {
	// pool 容量
	capacity int32
	// 当前正在执行任务的goroutine数量
	running int32
	// 空闲队列中worker最近一次执行与当前时间之差 如果大于过期时间则由定时清理任务回收
	expiryDuration time.Duration
	// worker执行完成放回队列 使用freeSignal告知被阻塞的请求绑定哪一个空闲worker
	freeSignal chan sig
	// 存放当前全部空闲workers 新来任务会在这绑定空闲worker；没有时 if超过容量则阻塞/ 否则开新worker
	workers []*Worker
	// 关闭pool时 通知所有goroutine推出运行
	release chan sig
	// 用来支持pool的同步操作
	lock sync.Mutex
	// 保证pool的关闭只执行一次
	once sync.Once
}

//type Worker struct {
//	pool *Pool
//	task chan f
//	recycleTime time.Time
//}

type Data interface {
	length() int
	less(i, j int) bool
	swap(i, j int)
}

// 把sort操作抽象出来
func Sort(data Data) {
	// ....不关心具体的排序算法使用哪种
	//只要实现不同类型的data 都可以使用同一个sort逻辑实现排序
}

func main() {
	//var n int32
	//var wg sync.WaitGroup
	//for i := 0; i < 100; i++ {
	//	wg.Add(1)   // Add(1) 通知程序有一个需要等待完成的任务
	//	go func() {
	//		//atomic.AddInt32(&n, 1)
	//		n++
	//		wg.Done() // 调用wg.Done() 表示正在等待的程序已经执行完成了
	//	}()
	//}
	//wg.Wait() //wg.Wait会阻塞当前程序直到等待的程序都执行完成为止
	//fmt.Println(atomic.LoadInt32(&n)) // output:1000

	var mu sync.Mutex
	go func(){
		fmt.Println("你好, 世界")
		mu.Lock()
	}()
	mu.Unlock()
}

//type StructA struct{}
//
//func NewStructA() Data {
//	return &StructA{}
//}
//
//func (a *StructA)swap(i, j int){
//}
//
//func (a *StructA)length() int{
//}
//func (a *StructA)less(i, j int) bool {
//}
//
//type Writer interface {
//	Write(b []byte)(int, error)
//}
//
//func WriteTo(w Writer)(n int64, err error) {
//	// ...
//}


