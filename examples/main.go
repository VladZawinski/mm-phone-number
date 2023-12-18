package main

import (
	"fmt"

	mmphonenumber "github.com/VladZawinski/mm-phone-number"
)

func main() {
	// test, _ := mmphonenumber.NormalizeInput("+၉၅၉၇၈၄၁၂၃၄၅၆")
	test := mmphonenumber.IsValidMMPhoneNumber("09978412345")
	fmt.Println(test)
}
