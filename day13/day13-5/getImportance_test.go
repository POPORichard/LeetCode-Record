package day13_5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}
//广度搜索
func getImportance(employees []*Employee, id int) int{
	//建立hash
	mp := map[int]*Employee{}
	for _, employee := range employees {
		mp[employee.Id] = employee
	}
	//建立队列
	queue := []int{id}
	total := 0

	for len(queue) > 0 {
		//出队列
		employee := mp[queue[0]]
		queue = queue[1:]

		total += employee.Importance

		//广度搜索下几个加入队列的值
		for _, subId := range employee.Subordinates {
			queue = append(queue, subId)
		}
	}
	return total
}

func TestGetImportance(t *testing.T){
	//case 1
	employees1 := Employee{
		Id:           1,
		Importance:   5,
		Subordinates: []int{2,3},
	}
	employees2 := Employee{
		Id:           2,
		Importance:   3,
		Subordinates: nil,
	}
	employees3 := Employee{
		Id:           3,
		Importance:   3,
		Subordinates: nil,
	}

	res := getImportance([]*Employee{&employees1, &employees2, &employees3},1)
	assert.Equal(t, 11, res, "Error in case 1")
}


