package pkg

import (
	"errors"
	"fmt"
)

const (
	//Id = 3333
	ID = 3333
)

func Test() error {

	//fmt.Printf("Golangci-lint test(%d)", Id)
	fmt.Printf("Golangci-lint test(%d)", ID)

	return errors.New("Test Return")
}
