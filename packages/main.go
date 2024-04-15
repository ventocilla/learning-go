package main

import (
	"fmt"

	helper "example.com/packages/helpers"
)

func main() {
	var x helper.SomeType
	x.TypeName = "ABC"
	x.TypeInt = 123
	fmt.Println(x.TypeName)
	fmt.Println(x.TypeInt)
}
