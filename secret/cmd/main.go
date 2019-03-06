package main

import (
	"fmt"
	"secret"
)

func main() {
	v := secret.Memory("my-fake-key")
	err := v.Set("demo-key", "somecrazyvalue")

	if err != nil {
		panic(err)
	}

	plain, err := v.Get("demo-key")
	if err != nil {
		panic(err)
	}

	fmt.Println(plain)
}
