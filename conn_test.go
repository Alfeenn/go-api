package main

import (
	"fmt"
	"strconv"
	"testing"
)

func TestW(t *testing.T) {
	var x string
	x = ""
	y, _ := strconv.Atoi(x)
	fmt.Println(y)
}
