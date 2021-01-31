package main

import (
	"fmt"
	"strconv"
)

func FixNumbers(number int) string {
	if number < 10 {
		return fmt.Sprintf("0%v", number)
	}

	return strconv.Itoa(number)
}
