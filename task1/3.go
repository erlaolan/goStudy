package main

import "fmt"

func main() {
	str1 := "([[}])"
	runStr := []rune(str1)
	runeLen := len(runStr)
	if runeLen%2 != 0 {
		fmt.Println("不是有效字符串")
		return
	}
	for i := 0; i < runeLen/2; i++ {
		newStr := string(runStr[i]) + string(runStr[runeLen-i-1])
		fmt.Println(newStr)
		if newStr != "()" && newStr != "[]" && newStr != "{}" {
			fmt.Println("不是有效字符串")
			return
		}
	}
	fmt.Println("是有效字符串")

}
