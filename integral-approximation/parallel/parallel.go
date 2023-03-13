package main

import (
	"fmt"
	"integral-approximation/concurrent"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	args := os.Args[1:]
	if len(args) != 3 {
		usage()
		fmt.Println("Wrong number of args need three. For  threadcount and a, b")
		return
	}
	var wg sync.WaitGroup
	threadCount, _ := strconv.Atoi(args[0])
	a, _ := strconv.ParseFloat(args[1], 64)
	b, _ := strconv.ParseFloat(args[2], 64)
	fmt.Println("Integral Bounds (" + args[1] + ", " + args[2] + ")")
	fmt.Println("Approximate integral on interval a b on x^3")
	fmt.Println("DISCLAIMER: Approximation less accurate the bigger the numbers get.")
	executor := concurrent.NewWorkBalancingExecutor(threadCount, 20, &wg) //Switch out for either implementation
	var futures []concurrent.Future
	intervals := threadCount
	interval_range := (b-a) / float64(intervals)
	var start float64
	start = a
	var finish float64
	finish = a + interval_range
	sum := 0.0
	for i := 0; i < threadCount ; i++ {
		wg.Add(1)
		task := concurrent.NewIntervalTask(start, finish, threadCount, &sum)
		start = start + interval_range
		finish = finish + interval_range
		futures = append(futures, executor.Submit(task))
	} 
	for _, future := range futures {
		go future.Get()
	}
	wg.Wait()
	go executor.Shutdown()
	fmt.Println(sum)
	fmt.Println("Approx: ", sum)
	fmt.Println("Time it took parallel: ", time.Since(startTime).Seconds())
}

func usage() {
	fmt.Println("This program returns an estimate for an integral on the graph y = x^3.  The user must enter the bounds (a,b) respectively as an input.")
}
