package main

import (
	"fmt"

	"github.com/mazen160/go-random"
)

func main() {
	charset := "abcdef0123456789"
	length := 64
	data, err := random.Random(length, charset, true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)
}
