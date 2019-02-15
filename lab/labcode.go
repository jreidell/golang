// Lab is a collection of Tour of Go examples and learning excercises
// This code corresponds to the Packages, variables, and functions module
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

	timetest()

	randtest()

	sqrttest(7)

	fmt.Println("Have some Pi! ", math.Pi)

	fmt.Println()

	fmt.Println("I'm adding 43 + 55 = ", add(43, 55))

	fmt.Println()

	writeswapresult()

	//fmt.Println("Named Return example via func nakedreturnsplit(17), ", nakedreturnsplit(17))
	fmt.Println(nakedreturnsplit(17))

	fmt.Println()

	cmplxmath()

	defaultvalues()

	typeconvtest()

	funwithconst()

	funwithnumconst()

	bitshifter()
}

// bitshifter is to demonstrate the use of bitwise shift operators
func bitshifter() {
	num := 30
	shift := 1

	result := num >> uint(shift)

	fmt.Println("30 >> 1 = ", result, " or 30 / 2 = 15")
	fmt.Println()

	result = num << uint(shift)

	fmt.Println("1 >> 30 = ", result, " or 30 * 2 = 60")
	fmt.Println()
}

// Pi decalred for funwithconst()
const Pi = 3.14

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// needInt will return a an int, used with funwithnumconst
func needInt(x int) int {
	return x*10 + 1
}

// needFloat will return a float from an int, used with funwithnumconst
func needFloat(x float64) float64 {
	return x * 0.1
}

// funwithnumconst demostrates numeric constant values, depends on needInt and needFloat functions tho
func funwithnumconst() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	fmt.Println()
}

// funwithconst is to demonstrate the use of declaring constants
func funwithconst() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println()
}

// typeconvtest is meant to demonstrate type conversions in Go
func typeconvtest() {

	//types can be inferred by the value on the right
	v := 42                            // change me!
	fmt.Printf("v is of type %T\n", v) // var i = 42

	v1 := 3.142                          // change me!
	fmt.Printf("v1 is of type %T\n", v1) // var i = 42

	v2 := 3.142 + 0.51i                  // change me!
	fmt.Printf("v2 is of type %T\n", v2) // var i = 42

	// var f = float64(i)
	// var u = uint(f)
	// fmt.Println(i, f, u)

	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	// fmt.Println(x, y, z)

	fmt.Println()
}

// defaultvalues demonstrates default values for specific types
func defaultvalues() {
	var i int
	var f float64
	var b bool
	var s string

	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	fmt.Println()
}

// cmplxmath demonstrates complex math problems and printing them using Go formatting
func cmplxmath() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)

	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)

	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Println()
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
	fmt.Println()
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
	fmt.Println()
}

// randtest random number gen demo
func randtest() {
	fmt.Println("Have some random number yo!")

	fmt.Println("The number is now ", rand.Intn(time.Now().Nanosecond()))
	fmt.Println()
}

// timetest deomonstrates getting the current system time
func timetest() {
	fmt.Println("Here we are again!")

	fmt.Println("The time is now ", time.Now())
	fmt.Println()
}
