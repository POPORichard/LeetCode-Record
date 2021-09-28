package day21_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	for _, booking := range bookings {
		nums[booking[0]-1] += booking[2]
		if booking[1] < n {
			nums[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums

}

func TestCorpFlightBookings(t *testing.T){
	//case 1
	bookings := [][]int{{1,2,10},{2,3,20},{2,5,25}}
	res := corpFlightBookings(bookings, 5)
	assert.Equal(t, []int{10,55,45,25,25}, res, "Error in case 1")
}