package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}
	var Count int = 0
	var Sum, Min, Max int64
	Sum = 0
	Min = nums[1] - nums[0]
	Max = nums[1] - nums[0]

	for i:= 0; i < len(nums) - 1; i++ {
		Count++
		diff := nums[i+1] - nums[i]
		Sum += diff
		if diff < Min {
			Min = diff
		}
		if diff > Max {
			Max = diff
		}
	}
	return Stats{Count: Count, Sum: Sum, Min: Min, Max: Max}
}

