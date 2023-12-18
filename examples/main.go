package main

import (
	"fmt"

	mmphonenumber "github.com/VladZawinski/mm-phone-number"
)

func main() {
	test, _ := mmphonenumber.SanitizeInput("+959962460148")
	fmt.Println(test)
}
