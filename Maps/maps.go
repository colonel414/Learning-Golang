// Maps are Go’s built-in abstract data type that stores a collection of (key, value) pairs,
// such that each possible key appears at most once in the collection (sometimes called hashes or dicts in other languages)
package main

import (
	"fmt"
	"maps"
)

func main() {
	//To create an empty map, use the builtin make: make(map[key-type]val-type).
	scores := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	scores["Alice"] = 95
	scores["Bob"] = 87

	// Note that maps appear in the form map[k:v k:v] when printed with fmt.Println.
	fmt.Println("map:", scores)

	// Get a value for a key with name[key].
	aliceScore := scores["Alice"]
	fmt.Println("Alice:", aliceScore)

	bobScore := scores["Bob"]
	fmt.Println("Alice:", bobScore)

	// If the key doesn't exist, the zero value of the value type is returned.
	joeScore := scores["Joe"]
	fmt.Println("Joe:", joeScore)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(scores))

	// The builtin delete removes key/value pairs from a map.
	delete(scores, "Bob")
	fmt.Println("map:", scores)

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(scores)
	fmt.Println("map:", scores)

	// The optional second return value when getting a value from a map indicates if the key was present in the map.
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	// Here we didn't need the value itself, so we ignored it with the blank identifier _.
	_, alicescore := scores["Alice"]
	fmt.Println("prs:", alicescore)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// The maps package contains a number of useful utility functions for maps.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
