package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	timetest()

	fmt.Println()

	randtest()

	fmt.Println()

	sqrttest(7)

	fmt.Println()

	fmt.Println("Have some Pi! ", math.Pi)
}

func sqrttest(f float64) {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(f))
}

func randtest() {
	fmt.Println("Have some random number yo!")

	fmt.Println("The number is now ", rand.Intn(time.Now().Nanosecond()))
}

func timetest() {
	fmt.Println("Here we are again!")

	fmt.Println("The time is now ", time.Now())
}
