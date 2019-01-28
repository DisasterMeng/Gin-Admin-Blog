package main

import (
	"fmt"
	"testing"
)

func findEnd(start int, arr []int) int {
	height := arr[start]
	newArr := arr[start+1:]
	max, index := 0, 0
	for i, a := range newArr {
		if a >= height {
			return start + i + 1
		} else {
			if a > max {
				max = a
				index = i
			}
		}
	}
	if index != 0 {
		return start + index + 1
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getTotal(start, end int, arr []int) int {
	m := min(arr[start], arr[end])
	total := m * (end - start - 1)
	for i := start + 1; i < end; i++ {
		total -= arr[i]
	}
	return total
}

func trap(height []int) int {
	start, end, result := 0, 0, 0
	for start != len(height) {
		if height[start] < 1 {
			start += 1
			continue
		}
		end = findEnd(start, height)
		if end-start == 1 {
			start += 1
			continue
		}

		if end == 0 {
			start += 1
		} else {
			total := getTotal(start, end, height)
			result += total

			start = end
		}
	}
	return result
}

func TestWater(t *testing.T) {
	//arr := []int{4,2,3}
	arr := []int{5, 5, 1, 7, 1, 1, 5, 2, 7, 6}
	//arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	//arr := []int{1}
	result := trap(arr)
	fmt.Println(result)
}
