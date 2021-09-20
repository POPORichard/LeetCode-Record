package day13_2

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func sortSentence(s string) string {
	a := strings.Split(s, " ")
	sort.Slice(a, func(i, j int) bool { return a[i][len(a[i])-1] < a[j][len(a[j])-1] })
	for i, s := range a {
		a[i] = s[:len(s)-1]
	}
	return strings.Join(a, " ")
}

func TestSortSentence(t *testing.T){
	res := sortSentence("is2 sentence4 This1 a3")
	assert.Equal(t, "This is a sentence", res, "Error in case 1")

}
