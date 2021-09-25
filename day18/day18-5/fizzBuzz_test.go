package day18_5

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func fizzBuzz(n int) []string {
	out := []string{}
	if n==0 {
		return out
	}

	for i := 1;i <= n;i++{
		if i%3 == 0 && i%5 == 0 {
			out = append(out,"FizzBuzz")
		} else if i%3==0{
			out = append(out,"Fizz")
		} else if i%5==0{
			out = append(out,"Buzz")
		} else {
			out = append(out,strconv.Itoa(i))
		}
	}

	return out
}

func TestFizzBuzz(t *testing.T){
	//case 1
	res := fizzBuzz(15)
	assert.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz"}, res, "Error in case 1")

}
