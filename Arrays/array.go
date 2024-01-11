// An array is a numbered sequence of elements of a specific length. In typical Go code, slices are much more common; arrays are useful in some special scenarios
package main

import "fmt"

func main(){
	// Here we create an array a that will hold exactly 5 ints. 
	// The type of elements and length are both part of the array’s type. 
	// By default an array is zero-valued, which for ints means 0s.
	var i [5]int
	fmt.Println("count", i)

	// We can set a value at an index using the ``array[index] = value`` syntax, and get a value with ``array[index]``
	i[2] = 100
    fmt.Println("set:", i)
    fmt.Println("get:", i[4])

	// The builtin len returns the length of an array
	fmt.Println("len:", len(i))

	// Use this syntax to declare and initialize an array in one line
	j := [5]int{1, 2, 3, 4, 5}
    fmt.Println("count:", j)

	// Array types are one-dimensional, but you can compose types to build multi-dimensional data structures
	// Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println
	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}