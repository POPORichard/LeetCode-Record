package day15_6

import(
	"github.com/stretchr/testify/assert"
	"testing"
)

func minOperations(logs []string) int {
	result := 0
	for i := range logs {
		op := logs[i][:2]
		switch op {
		case "..":
			if result > 0 {
				result--
			}
		case "./":
		default:
			result++
		}
	}
	return result
}

func TestMinOperations(t *testing.T){
	//case 1
	logs := []string{"d1/","d2/","../","d21/","./"}
	res := minOperations(logs)
	assert.Equal(t, 2, res, "Error in case 1")
}