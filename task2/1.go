package main

import (
	"fmt"
)

func main() {
	num := 10
	fmt.Println(add10(&num))

}

func add10(target *int) int {
	return *target + 10
}
