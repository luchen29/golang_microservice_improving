package main

import (
	myerrors "awesomeProject/errors"
	"fmt"
	"github.com/pkg/errors"
	_ "net/http/pprof"
	"os"
)

func main() {
/*
	// 21/06/04 practice1: implement error interface by yourself
	err := myerrors.Test()
	// 虽然Test return的是error interface； 但可以通过type类型断言方式 将err转换成该类型
	switch err := err.(type) {
	case nil:
		fmt.Printf("current error is nil\n")
	case *myerrors.MyError:
		fmt.Printf("current err type is myErr: %v %v \n", err.ErrorMessage, err.ErrorLine)
	case error:
		fmt.Printf("this is error type\n")
	default:
		fmt.Printf("this is default error type\n")
	}

	newErr := myerrors.New("this is a new error")
	fmt.Printf("this is a new error: %v\n", newErr)
*/

	// 21/06/06 practice2: wrap an err & get the err
	_, err := myerrors.ReadFile("testPath")
	if err != nil {
		// errors.Cause(err) is used to the the roooooot err;
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace of the err:\n %+v\n", err)
		os.Exit(1)
	}
	_, err = myerrors.ReadConfig()
	if err != nil {
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace of the err:\n %+v\n", err)
		os.Exit(1)
	}
	// if an error is already been handled, do not return error any longer


	/* eg1:
		ch := make(chan int)
		go func() {

			fmt.Println("print this later")
			ch <- 1
		} ()
		fmt.Println("print this first")
		x := <-ch
		fmt.Println(x)
	*/
	/* eg2:
		s := []int{7, 2, 8, -9, 4, 0}
		c := make(chan int)
		go sum(s[:len(s)/2], c)
		go sum(s[len(s)/2:], c)
		x, y := <-c, <-c // receive from c
		fmt.Println(x, y, x+y)
	*/
	/* eg3:
		c2 := make(chan int, 10)
		go fibonacci(cap(c2), c2)
		for i := range c2 {
			fmt.Println(i)
		}
	*/
	/* eg4
		c3 := make(chan int, 10)
		c3 <- 1
		c3 <- 2
		close(c3)
		fmt.Println(<-c3)
		fmt.Println(<-c3)
		fmt.Println(<-c3)
		fmt.Println(<-c3)
		fmt.Println(<-c3)
	*/

	//go func() {
	//	time.Sleep(10 * time.Second)
	//}()
	//c := make(chan int, 10)
	//go fibonacci(cap(c), c)
	//for i := range c {
	//	fmt.Println(i)
	//}
	//fmt.Println("Finished")

	//ch := make(chan int, 10)
	//ch <- 1
	//ch <- 2
	//ch <- 3

	// 关闭函数非常重要,若不执行close(),那么range将无法结束,造成死循环
	// close(ch)

	//for v := range ch {
	//	fmt.Println(v)
		//var mu sync.Mutex
	//go func(){
	//	fmt.Println("你好, 世界")
	//	mu.Lock()
	//}()
	//mu.Unlock()

	//var mu sync.Mutex
	//mu.Lock()
	//go func() {
	//	fmt.Println("inside>>>>")
	//	mu.Unlock()
	//}()
	//mu.Lock()

	//ch := make(chan int, 10) // 成果队列
	//go Producer(3, ch) // 生成 3 的倍数的序列
	//go Producer(5, ch) // 生成 5 的倍数的序列
	//go Consumer(ch)    // 消费 生成的队列
	//
	//// 在这里阻塞main函数 手动command+C退出保证consumer稳定输出
	//sig := make(chan os.Signal, 1)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//fmt.Printf("quit (%v)\n", <-sig)


	//http.ListenAndServe(":8888", nil)

	//persons:=make(map[string]int)
	//persons["张三"]=19
	////mp:=&persons
	//fmt.Printf("原始map的内存地址是：%p\n",&persons)
	//modify(persons)
	//fmt.Println("map值被修改了，新值为:",persons)

	//m := make(map[int]Bird)
	//fmt.Printf("map地址：%p\n", &m)
	//parrot1.Age = 1
	//m[0] = parrot1
	//parrot1.Age = 2
	//m[1] = parrot1
	//parrot1.Age = 3
	//m[2] = parrot1
	//parrot1.Age = 4
	//
	//for _, v := range m {
	//	fmt.Printf("parro本身：%v\n, parrot的年龄：%v\n, 内存地址：%p\n", v, v.Age, &v)
	//	newAge := 4
	//	v.Age = newAge
	//	//m[k] = v
	//}
	//fmt.Println("map打印出来： %v\n", m)



	//start := time.Now()
	//ch := make(chan string, 1)
	//var wg sync.WaitGroup
	//
	//// check header
	//wg.Add(1)
	//go HeaderCheck()
	//
	//// check url
	//wg.Add(1)
	//go UrlCheck()
	//
	//// check body
	//wg.Add(1)
	//go BodyCheck()
	//
	//wg.Wait()
	//
	//latency := time.Since(start).Seconds()
	//logs.CtxInfo(ctx, "三个goroutine全部执行完成的 最长耗时（seconds）：%v", latency)
}

//func modify(p map[string]int){
//	fmt.Printf("函数里接收到map的内存地址是：%p\n",&p)
//	p["张三"]=20
//}


//var parrot1 = Bird{Age: 1, Name: "Blue"}
//
//type Bird struct {
//	Age  int
//	Name string
//}
//
//// 生产者: 生成 factor 整数倍的序列
//func Producer(factor int, out chan<- int) {
//	for i := 0; ; i++ {
//		out <- i*factor
//	}
//}
//// 消费者
//func Consumer(in <-chan int) {
//	for v := range in {
//		fmt.Println(v)
//	}
//}
//
//func sum(s []int, c chan int) {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	c <- sum // send sum to c
//}
//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n; i++ {
//		c <- x
//		x, y = y, x+y
//	}
//	//close(c)
//
