package main

import "fmt"

func foo() func(string) {
	return func(z string) {
		fmt.Println("returning a func " + z)
	}
}
