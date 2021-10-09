package main

import (
	"runtime"

	"example.com/appointy"
)

func main() {
	//Running the HTTP Request Handlers
	appointy.RunServer()
}

func init() {
	// Setting number of processors to number of physical CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())
}
