package main

import (
	"fmt"
)

/*Implement the shiftASCII(s string, step int) string function,which takes an ASCII string s as input and returns a new string
where each character in the input string is shifted forward by the step number.

If, after the shift, the character code is out of ASCII, then the number must be taken modulo 256.*/

func main() {
	s := "abc"
	fmt.Println(shiftASCII(s, 513))
}

func shiftASCII(s string, step int) string {
	if step == 0 {
		return s
	}

	if step > 256 || step < -256 {
		step %= 256
	}

	bytes := []byte(s)

	for i := 0; i < len(bytes); i++ {
		bytes[i] += byte(step)
	}

	return string(bytes)
}
