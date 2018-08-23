// Package main and function main is the entry point.
package main

// Will not compile if you import a package but don't use it.
import (
	"fmt"
	"os"
)

func main() {
	// len returns the size of a string, or the number of values in a
	// dictionary, or the number of elements in an array
	if len(os.Args) != 2 {
		fmt.Println("Need at least 1 argument")
		os.Exit(1)
	}

	// we prefix the function name with the package
	fmt.Println("It's over", os.Args[1])
}
