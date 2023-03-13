package main

import (
	"fmt"
	"math/rand"
	"math"
	"time"
	"os"
	"strconv"
)

func main() {
	start := time.Now()
	a,b, err:= handle_input_x3()
	if err == "yes" {
		usage()
		fmt.Println("Wrong number of args need three. For a, b")
		return 
	}
	approx := run_monte_carlo_approx(a, b)

	fmt.Println("Approximate integral on interval a b on x^3")
	fmt.Println("DISCLAIMER: Approximation less accurate the bigger the numbers get.")
	fmt.Println("Approx: ", approx)
	fmt.Println("Time it took sequentially: ", time.Since(start).Seconds())
}

func run_monte_carlo_approx(a float64, b float64) float64{
	
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var low float64 = -1000000
	var high float64 = 1000000
	var countInsideAbove float64 = 0
	var countTotal float64 = 0
	var countInsideBelow float64 = 0

	for i := 0; i < 100000000; i++ {
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
	fmt.Println(fractionAbove, fractionBelow)
	result := totalArea * (fractionAbove - fractionBelow)
	return result
}

func handle_input_x3() (float64, float64, string) {
	args := os.Args[1:]
	err1 := "yes"
	if len(args) != 2 {
		fmt.Println("Error 1")
		return 0, 0, err1
	}
	var a float64
	var b float64
	a, _ = strconv.ParseFloat(args[0], 64)
	b, _ = strconv.ParseFloat(args[1], 64)
	return a, b, "no"
}

func usage() {
	fmt.Println("This program returns an estimate for an integral on the graph y = x^3.  The user must enter the bounds (a,b) respectively as an input.")
}
