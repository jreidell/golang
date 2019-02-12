// Lab is a collection of Tour of Go examples and learning excercises
package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {

	python, c, java := true, false, "no!"
	k := 3
	var i, j int = 1, 2

	fmt.Println("Various Var Declaration Example: ", k, python, c, java, j, i)

	fmt.Println()

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

	writeswapresult()

	fmt.Println()

	//fmt.Println("Named Return example via func nakedreturnsplit(17), ", nakedreturnsplit(17))
	fmt.Println(nakedreturnsplit(17))

	fmt.Println()

	cmplxmath()

	fmt.Println()

	defaultvalues()

	fmt.Println()

	typeconvtest()
}

// typeconvtest is meant to demonstrate type conversions in Go
func typeconvtest() {

	// var i = 42
	// var f = float64(i)
	// var u = uint(f)
	// fmt.Println(i, f, u)

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

// defaultvalues demonstrates default values for specific types
func defaultvalues() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// cmplxmath demonstrates complex math problems and printing them using Go formatting
func cmplxmath() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)

	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)

	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// nakedreturnsplit demonstrates a method using a naked return. return vars are named in the signature and not explicitly returned
func nakedreturnsplit(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// writeswapresult write out the swap method results
func writeswapresult() {
	a, b := swap("Rock", "You")

	fmt.Println("a, b := swap(\"Rock\", \"You\")")

	fmt.Println("I'm swapping these out, Rock You ! ", a, b, "!")
}

// swap the values that are input to the method, return them in reverse order
func swap(x, y string) (string, string) {
	return y, x
}

// add demonstrates addition
func add(x, y int) int {
	return x + y
}

// sqrttest math demo
func sqrttest(f float64) {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(f))
}

// randtest random number gen demo
func randtest() {
	fmt.Println("Have some random number yo!")

	fmt.Println("The number is now ", rand.Intn(time.Now().Nanosecond()))
}

// timetest deomonstrates getting the current system time
func timetest() {
	fmt.Println("Here we are again!")

	fmt.Println("The time is now ", time.Now())
}
