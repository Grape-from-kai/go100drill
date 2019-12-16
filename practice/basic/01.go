package main

import (
	"fmt"
	"strconv"
)

func main() {
	const (
		lhs = 12345
		rhs = 23456
	)

	var output = strconv.Itoa(lhs) + "+" + strconv.Itoa(rhs) + "=" + strconv.Itoa(lhs + rhs)
	fmt.Println(output)
	
}
