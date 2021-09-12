package day5_10

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	l := "2006-01-02"
	t1, _ := time.Parse(l, date1)
	t2, _ := time.Parse(l, date2)
	return int(math.Abs(t1.Sub(t2).Hours()))/24
}

func TestDaysBetweenDates(t *testing.T)  {
	//case 2
	res := daysBetweenDates("2020-01-15","2019-12-31")
	assert.Equal(t, 15, res, "Error in case 1")

}