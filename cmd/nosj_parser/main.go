package main

import (
	"fmt"
	deserializer "nosj_parser/internal/deserialize"
	"os"
)

func main() {
	os.Exit(run())
}

func run() int {
	result, err := deserializer.Deserialize("(<abcd:t e s t s>)")

	if err != nil {
		fmt.Println(err)
		return 66
	}

	fmt.Println(result)
	return 0
}
