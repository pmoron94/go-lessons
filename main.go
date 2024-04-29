// All files starts with package
package main

import (
	/*
		Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
		The format 'verbs' are derived from C's but are simpler.

		By convention, the package name is the same as the last element of the import path.
	*/
	"fmt"
	"math/rand"

	"github.com/pabloMoron/go-lessons/02-imports/constants"
)

func main() { // If cloud environment => main() receives params
	fmt.Println("Random:", rand.Intn(10)) // Add newline
	constants.PrintlnInts()
}

// in console type go run ./main.go
// in console type go build ./main.go to compile executable file