package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := []string{"moon", "monkey", "moonligit"}

	prifix := str1[0]
	for _, k := range str1 {
		for strings.Index(k, prifix) != 0 {
			if len(prifix) == 0 {
				fmt.Println("没有公共前缀")
				return
			}
			prifix = prifix[:len(prifix)-1]
		}

	}
	fmt.Println("公共前缀是", prifix)
}
