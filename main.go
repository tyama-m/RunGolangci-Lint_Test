package main

import (
	"fmt"

	"golangci-lint.com/m/pkg"
)

const (
	Id = 3333
	//ID = 3333
)

func main() {

	fmt.Printf("Golangci-lint test(%d)", Id)
	//fmt.Printf("Golangci-lint test(%d)", ID)

	//test()
	_ = pkg.Test()
}

// func test() error {

// 	return errors.New("test return")
// }
