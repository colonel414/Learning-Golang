package main

import (
	"fmt"
	"math"
)

// const declares a constant value
const s string = "Colonel"

func main() {
	fmt.Println(s)

	// A const statement can appear anywhere a var statement can.
	const t = 800000000

	// Constant expressions perform arithmetic with arbitrary precision.
	const u = 3e20 / t
	fmt.Println(u)

	// A numeric constant has no type until it’s given one, such as by an explicit conversion
	fmt.Println(int64(u))

	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(t))
}