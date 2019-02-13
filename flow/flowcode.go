// Flow is a collection of Tour of Go examples and learning excercises
// This code corresponds to the "Flow control statements: for, if, else, switch and defer" module
package main

import (
	"fmt"
	"math"
)

func main() {

	forloopexample()

	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println()

	// order of operations will allow the calls to pow(...) to complete
	// prior to the execution of the fmt.Println call below
	fmt.Println(
		// 3 squared is less than 10 so the result is returned
		pow(3, 2, 10),
		// 3 cubed is greater than 20 so 20 is returned here
		pow(3, 3, 20),
	)
	fmt.Println()

}

// infiniteloop is simply an infinite loop that demonstrates creating a for loop without a
// conditional which creates an infinite loop and requires Ctrl+C to exit the application
func infiniteloop() {
	for {

	}
}

// forloopexample is an example of a basic for loop
func forloopexample() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println()

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	fmt.Println()
}

// sqrt function is an example of an if statement
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// pow function demonstrates an if statement that includes a
// short statement to execute before testing the condition
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// this Printf will execute before the
		// Println in the main function/method
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
