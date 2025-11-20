package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(plusOne(arr))
}

func plusOne(digits []int) []int {
	str := ""

	for i := 0; i < len(digits); i++ {
		str += strconv.Itoa(digits[i])
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	str1 := strconv.Itoa(num + 1)
	strArr := strings.Split(str1, "")
	intArr := make([]int, 0, len(strArr))
	for _, s := range strArr {
		num, _ := strconv.Atoi(s)
		intArr = append(intArr, num)
	}
	return intArr
}
