package day16_8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	m := map[string]int{
		"type": 0,
		"color": 1,
		"name": 2,
	}
	var count int
	idx := m[ruleKey]
	for _, i := range items {
		if i[idx] == ruleValue {
			count++
		}
	}
	return count
}

func TestCountMatches(t *testing.T){
	//case 1
	items := [][]string{{"phone","blue","pixel"},{"computer","silver","lenovo"},{"phone","gold","iphone"}}
	res := countMatches(items, "color", "silver")
	assert.Equal(t, 1, res, "Error in case 1")
}