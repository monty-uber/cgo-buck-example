package main

import "C"

//export GoFunction
func GoFunction(a, b int) int {
	return a + b
}