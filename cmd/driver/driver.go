package main

import (
	"flag"
	"fmt"
	"internal/driver"
)

// main function for driver
func main() {
	N := flag.Int("N", 1, " The number N of map tasks")
	M := flag.Int("M", 1, " The number M of reduce tasks")
	flag.Parse()
	if *M <= 0 || *N <= 0 {
		fmt.Println("N and M should be greater than 0")
		return
	}
	driver.Main(*N, *M)
}
