package main

import (
	"fmt"
	"unicode"
)

func main() {

	for _, r := range "中gé，；； ~^" {
		fmt.Println(unicode.IsLetter(r))
	}
}
