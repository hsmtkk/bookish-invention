package main

import (
	"fmt"

	"github.com/hsmtkk/bookish-invention/collatz"
)

func main() {
	for i := 10; i < 100; i++ {
		ns := collatz.CollatzNumbers(i)
		fmt.Println(ns)
	}
}
