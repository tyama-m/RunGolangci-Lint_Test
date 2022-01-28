package main

import (
	"RunGolangci-Lint_Test/platform"
	"errors"
	"fmt"
)

const (
	Id = 3333
	//ID = 3333
)

func main() {

	fmt.Printf("Golangci-lint test(%d)", Id)
	//fmt.Printf("Golangci-lint test(%d)", ID)
	fmt.Printf("%s", platform.GetPlatform())

	//test()
	_ = test()
}

func test() error {

	return errors.New("test return")
}
