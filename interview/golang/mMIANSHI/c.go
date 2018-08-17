package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getGenerator(users map[string]int64) (generator func() string) {
	var sum_num int64
	name_arr := make([]string, len(users))
	offset_arr := make([]int64, len(users))

	var index int
	for Uname, count := range users {
		offset_arr[index] = sum_num
		name_arr[index] = Uname
		sum_num += count
		index += 1
	}

	generator = func() string {
		award_num := rand.Int63n(sum_num)
		return name_arr[b_search(offset_arr, award_num)]
	}
	return
}

func b_search(nums []int64, target int64) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			if mid+1 == len(nums) {
				return mid
			}
			if nums[mid+1] > target {
				return mid
			}
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1 //返回错误状态
}

func main() {
	var users map[string]int64 = map[string]int64{
		"a": 10,
		"b": 5,
		"c": 15,
		"d": 20,
		"e": 10,
		"f": 30,
	}

	rand.Seed(time.Now().UnixNano())
	award_stat := make(map[string]int64)

	for i := 0; i < 10000; i++ {
		name := getGenerator(users)
		if count, ok := award_stat[name()]; ok {
			award_stat[name()] = count + 1
		} else {
			award_stat[name()] = 1
		}
	}

	for user, count := range award_stat {
		fmt.Printf("User : %s ,award count : %d \n", user, count)
	}
}
