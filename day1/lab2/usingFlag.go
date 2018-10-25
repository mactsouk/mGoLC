//Jeffry M

// Create a Go program that uses flag with the following command line
// parameters:
// - A boolean named subdirs. - An integer named count.
// - A string named name.
// - A float named limit.

package main

import (
	"flag"
	"fmt"
)

func main() {
	subdirs := flag.Bool("s", false, "subdirectory")
	count := flag.Int("c", 0, "int")
	name := flag.String("n", "test", "string")
	limit := flag.Float64("l", 12312.4, "float64")

	flag.Parse()

	fmt.Println("-s:", *subdirs)
	fmt.Println("-c:", *count)
	fmt.Println("-n:", *name)
	fmt.Println("-l:", *limit)

	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
