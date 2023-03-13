package concurrent
import (
	"sync"
	"math/rand"
)
type Future1 struct{
	Queue * DEQueueObj
	Arr * [] *DEQueueObj
	Finished * int
	cond *sync.Cond
	threshold * int
	qn * int
	wg *sync.WaitGroup
}


func (f Future1) Get() interface{} {
	tasks := f.Queue
	sum := 0.0
	for true {
			size := *tasks.size
			if size > *f.threshold{
				res := tasks.PopTop()
				res.Run()
				continue
			}
		queueCount := len(*f.Arr)
		queue_num := rand.Intn(queueCount)
			//steal from queue_num
		if ! (*f.Arr)[queue_num].IsEmpty() {
			for i := 0; i < 10; i++ {
				if *(*f.Arr)[queue_num].size < *f.threshold  {
					break
				} 
				temp := (*f.Arr)[queue_num].PopBottom()
				tasks.PushBottom(temp)
			}
		}
		if *tasks.size <= *f.threshold{
			for ! tasks.IsEmpty() {
				res := tasks.PopTop()
				res.Run()
			}
			break
		}
	}
	f.wg.Done()
	//attempt to steal 
	return sum
}

type ExecutorService1 struct{
	Arr * [] *DEQueueObj
	idx * int
	queueCount * int
	Finished * int
	Cond *sync.Cond
	Threshold * int
	wg *sync.WaitGroup
}

func (e ExecutorService1) Submit(task interface{}) Future {

	idx := *(e.idx)
	res := task.(*Interval)
	a := res.A
	b := res.B
	interval_range := (b-a) / float64(100)
	start := a
	finish := start + interval_range
	queue := (*e.Arr)[idx:idx+1][0]
	for i := 0; i < 100; i ++ {
		add_interval := Interval{start, finish, *e.queueCount, 100, res.sum}
		queue.PushBottom(add_interval)
		start = finish
		finish = finish + interval_range
	}
	temp := *(e.idx)
	idx = idx + 1
	if idx == *(e.queueCount) {
		idx = 0
	}
	(*e.idx) = idx
	return Future1{queue, e.Arr, e.Finished, e.Cond, e.Threshold, &temp, e.wg}
}

func (e ExecutorService1) Shutdown() {
	e.wg.Wait()
}

func NewWorkBalancingExecutor(capacity, threshold int, wg *sync.WaitGroup) ExecutorService {
	var arr [] *DEQueueObj
	for i := 0; i < capacity; i++{
		dq := NewUnBoundedDEQueue()
		arr = append(arr, &dq)
	}
	// queuecount := capacity
	Finish := 0
	idx := 0
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)
	return ExecutorService1{&arr, &idx, &capacity, &Finish, cond, &threshold, wg}

}
