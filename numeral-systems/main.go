package main

import "fmt"

func main() {
	fmt.Printf("decimal %d\n", 42)

	fmt.Printf("binary %b\n", 42)

	fmt.Printf("hexidecimal %x\n", 42)

	fmt.Printf("printing ascii")
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
