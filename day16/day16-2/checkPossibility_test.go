package day16_2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func checkPossibility(nums []int) bool {
	wrong := 1
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			wrong--
			if wrong < 1 {
				return false
			}
			if i > 0 && y < nums[i-1] {
				nums[i+1] = x
			}
		}
	}
	return true
}

func TestCheckPossibility(t *testing.T){
	nums := []int{4,2,1}
	res := checkPossibility(nums)
	assert.Equal(t, false, res, "Error in case 1")
}
