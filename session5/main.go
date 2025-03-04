package main

import "strings"

func ReverseString(str string) string {
	str2 := []rune(str) // string are immutable
	for i := 0; i < len(str2)/2; i++ {
		str2[i], str2[len(str)-1-i] = str2[len(str)-1-i], str2[i]
	}
	return string(str2)
}
func main() {
	//fmt.Println(ReverseString("hello"))

	c := "hello"
	strings.Split(c, "")
	//c[0] = 'i'
}
