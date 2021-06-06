package practice

import (
	"fmt"
	"runtime"
	"time"
)

var (
	//最大任务队列数
	MaxWorker = 10
)

//有效载荷
type Payload struct {
	Num int
	Test string
}

//待执行的工作
type Job struct {
	Payload Payload
}

//任务队列channel
var JobQueue chan Job

//执行任务的工作者单元
type Worker struct {
	WorkerPool chan chan Job //工作者池--每个元素是一个工作者的私有任务channel
	JobChannel chan Job //每个工作者单元包含一个任务管道 用于获取任务
	quit chan bool //退出信号
	no int //编号
}

// 停止信号
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

//调度中心
type Dispatcher struct {
	//工作者池(任务队列池)
	WorkerPool chan chan Job
	//工作者单元数量
	MaxWorkers int
}

//创建调度中心
func NewDispatcher(maxWorkers int) *Dispatcher {
	//创建工作者池，存放待处理的任务队列，maxWorkers为最大任务队列数
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}

//创建一个新工作者单元
func NewWorker(workerPool chan chan Job, no int) Worker {
	fmt.Println("创建一个新工作者单元")
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
		no:         no,
	}
}

//循环 监听任务和结束信号
func (w Worker) Start() {
	//启动协程，监听任务
	go func() {
		for {
			// register the current worker into the worker queue.
			//工作者放回工作者池
			w.WorkerPool <- w.JobChannel
			//fmt.Println("w.WorkerPool <- w.JobChannel", w)
			select {
			case job := <-w.JobChannel:
				// 收到任务,执行打印任务
				fmt.Println(job.Payload.Test)
				//执行任务需要1秒时间
				time.Sleep(500 * time.Microsecond)
			case <-w.quit:
				// 收到退出信号，停止监听，结束该协程
				return
			}
		}
	}()
}

//调度，任务分发
func (d *Dispatcher) dispatch() {
	for {
		select {
		//从任务队列中获取任务
		case job := <-JobQueue:
			//fmt.Println("job := <-JobQueue:")
			go func(job Job) {
				//等待空闲worker (任务多的时候会阻塞这里)
				//从（10个）工作者池中获取一个任务队列channel，
				jobChannel := <-d.WorkerPool
				//fmt.Println("jobChannel := <-d.WorkerPool", reflect.TypeOf(jobChannel))
				// 将任务放到上述woker的私有任务channal中，jobChannel是一个无缓冲信道，每次只能放一个任务
				jobChannel <- job
				//fmt.Println("jobChannel <- job")
			}(job)
		}
	}
}

//工作者池的初始化，注意Run为Dispatcher结构体指针的方法，所以此方法内对Dispathcer的修改在方法外也可见
func (d *Dispatcher) Run() {
	// starting n number of workers
	//创建10个工作者单元
	for i := 1; i < d.MaxWorkers+1; i++ {
		worker := NewWorker(d.WorkerPool, i)
		worker.Start()
	}
	go d.dispatch()
}

//新建任务并放入任务队列
func addQueue() {
	for i := 0; i < 1000000; i++ {
		time.Sleep(10*time.Microsecond)
		fmt.Println("当前请求数：",i)
		// 新建一个任务
		payLoad := Payload{Num: 1, Test:"this is Test string"}
		work := Job{Payload: payLoad}
		// 任务放入任务队列channel
		fmt.Println("新任务入队列！")
		JobQueue <- work
		//fmt.Println("队列总长度:",cap(JobQueue))
		//fmt.Println("队列未消费任务数量:",len(JobQueue))
		//fmt.Println("JobQueue <- work")
		fmt.Println("当前协程数:", runtime.NumGoroutine())
		//time.Sleep(1 * time.Second)
	}
}

func main() {
	//建立任务队列，每个任务队列中可以放10个任务
	JobQueue = make(chan Job, 10)
	fmt.Println("成功建立任务队列！")
	//新建任务分发器
	dispatcher := NewDispatcher(MaxWorker)
	fmt.Println("成功建立任务分发器！")
	dispatcher.Run()

	//time.Sleep(1 * time.Second)
	go addQueue()
	time.Sleep(1000 * time.Second)
}