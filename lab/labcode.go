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

	fmt.Println()

	fmt.Println("I'm adding 43 + 55 = ", add(43, 55))

	fmt.Println()

	a, b := swap("Rock", "You")
	fmt.Println("a, b := swap(\"Rock\", \"You\")")
	fmt.Println("I'm swapping these out, Rock You ! ", a, b, "!")

	fmt.Println()

	//fmt.Println("Named Return example via func nakedreturnsplit(17), ", nakedreturnsplit(17))
	fmt.Println(nakedreturnsplit(17))

}

func nakedreturnsplit(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
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
