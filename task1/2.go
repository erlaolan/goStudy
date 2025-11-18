package main

import "fmt"

func main() {
	str1 := "123321"
	runStr := []rune(str1)
	runeLen := len(runStr)
	for i := 0; i < runeLen/2; i++ {
		if runStr[i] != runStr[runeLen-i-1] {
			fmt.Println("不是回文数")
			return
		}
	}
	fmt.Println("是回文数")

}
