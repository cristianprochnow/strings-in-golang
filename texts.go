package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := "Hello, World!"

	fmt.Println(text) // Hello, World!

	/* fmt.Println(text[0:3]) // Hel
	fmt.Println(text[2:17]) // Error... */

	fmt.Println(substring(text, 0, 3))  // Hel
	fmt.Println(substring(text, 2, 17)) // llo, World!

	fmt.Println(justNumbers())
}

func substring(text string, start int, end int) string {
	length := len(text)

	if start > length {
		start = length
	}
	if end > length {
		end = length
	}

	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}

	return text[start:end]
}

func justNumbers() string {
	str := "ABC1234DEF5678GHI"
	re := regexp.MustCompile("[0-9]+")
	nums := re.FindAllString(str, -1)

	return strings.Join(nums, "")
}
