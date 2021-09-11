package day4_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String()
}

func TestDayOfTheWeek(t *testing.T){
	//case 1
	res := dayOfTheWeek(31, 8, 2019)
	assert.Equal(t, "Saturday", res, "Error in case 1")

}
