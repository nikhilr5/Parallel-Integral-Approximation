package concurrent
import (
	"math/rand"
	"math"
	"time"
	"sync"
)
type Task interface{}

type DEQueue interface {
	PushBottom(task Task)
	IsEmpty() bool //returns whether the queue is empty
	PopTop() Task
	PopBottom() Task
}

type Interval struct {
	A float64
	B float64
	threadCount int
	breakUpCount int
	sum * float64
}
type DEQueueObj struct {
	Arr * [] *Interval
	mu *sync.Mutex
	size *int
	
	
}
func NewUnBoundedDEQueue() DEQueueObj {
	var arr [] *Interval
	lock := sync.Mutex{}
	size := 0
	dq := DEQueueObj{&arr, &lock, &size}
	return dq
}

func (queue *DEQueueObj) PushBottom(task Interval) {
	(*queue).mu.Lock()
	*(*queue).size++
	*(*queue).Arr = append(*(*queue).Arr, &task)
	(*queue).mu.Unlock()
	return
}

func (queue *DEQueueObj) IsEmpty() bool {
	(*queue).mu.Lock()
	size := *(*queue).size
	if size == 0 {
		(*queue).mu.Unlock()
		return true
	}
	(*queue).mu.Unlock()
	return false
}

func (queue *DEQueueObj)  PopTop() Interval {
	(*queue).mu.Lock()
	queue1 := (*queue)
	arr := *((*queue).Arr)

	if len(*queue.Arr) > 0 {
		ele := arr[0]
		*(*queue).Arr = arr[1:]

		queue1.mu.Unlock()
		(*queue1.size)-= 1
		return *ele
	}
	queue1.mu.Unlock()
	return Interval{}
}

func (queue *DEQueueObj) PopBottom() Interval {
	(*queue).mu.Lock()
	queue1 := (*queue)
	arr := *((*queue).Arr)

	if len(*queue.Arr) > 0 {
		ele := arr[len(*queue.Arr) -1 ]
		*(*queue).Arr = arr[:len(*queue.Arr) -1 ]

		queue1.mu.Unlock()
		(*queue1.size)-= 1
		return *ele
	}
	queue1.mu.Unlock()
	return Interval{}
}

func NewIntervalTask(a float64, b float64, threadCount int, sum *float64) Runnable {
	return &Interval{a,b, threadCount, 100, sum }
}

func (task *Interval) Run()  {
	run_monte_carlo_approx(task.A, task.B, task.threadCount, task.breakUpCount, task.sum)
}

func run_monte_carlo_approx(a float64, b float64, threadCount int, breakUpCount int, sum * float64){
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var low float64 = -1000000
	var high float64 = 1000000
	var countInsideAbove float64 = 0
	var countTotal float64 = 0
	var countInsideBelow float64 = 0
	maxVal := 100000000 / (threadCount * breakUpCount)
	for i := 0; i < maxVal; i++ {
		randNumX := a + r1.Float64() * (b - a)
		randNumY := low + r1.Float64() * (high - low)
		countTotal +=  1
		trueY := math.Pow(randNumX,3)
		if trueY > randNumY && randNumY >= 0{
			countInsideAbove = countInsideAbove + 1
		} else if randNumY < 0 && trueY < randNumY {
			countInsideBelow += 1
		}
	}
	fractionAbove := countInsideAbove / countTotal
	fractionBelow := countInsideBelow / countTotal
	totalArea := (high - low) * (b - a)
	result := totalArea * (fractionAbove - fractionBelow)
	(*sum) = (*sum) + result
}
