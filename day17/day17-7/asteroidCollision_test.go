package day17_7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func asteroidCollision(asteroids []int) []int {
	//建立栈
	stack := make([]int,0,len(asteroids))
	ans := make([]int,0,len(asteroids))

	//排除数组头中为负数的值
	for {
		if asteroids[0] <0{
			ans = append(ans, asteroids[0])
			asteroids = asteroids[1:]
			if len(asteroids) == 0{
				return ans
			}
		}else {
			stack = append(stack, asteroids[0])
			break
		}
	}

Loop:
	for i:=1 ;i<len(asteroids); i++{
		if asteroids[i] > 0{
			//若向右运行则入栈
			stack = append(stack, asteroids[i])
		}else {
			for{
				//若栈空
				if len(stack) == 0{
					ans = append(ans, asteroids[i])
					continue Loop

				}

				//用栈尾进行比较
				cur := stack[len(stack)-1]
				if cur > -asteroids[i]{
					continue Loop
				}
				if cur == -asteroids[i] {
					stack = stack[:len(stack)-1]
					continue Loop
				}
				stack = stack[:len(stack)-1]

			}
		}
	}
	//添加栈中值
	ans = append(ans, stack...)
	return ans
}

func TestAsteroidCollision(t *testing.T){
	//case 1
	asteroids := []int{10,2,-5}
	res := asteroidCollision(asteroids)
	assert.Equal(t, []int{10}, res, "Error in case 1")

	//case 2
	asteroids = []int{-2, -2, -2, -2}
	res = asteroidCollision(asteroids)
	assert.Equal(t, []int{-2, -2, -2, -2}, res, "Error in case 2")
}