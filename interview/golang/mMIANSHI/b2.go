package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getgenerator(users map[string]int64) (generator func() string) {
	var sum_num int64
	name_arr := make([]string, len(users))
	offset_arr := make([]int64, len(users))

	var index int
	for u_name, num := range users {
		name_arr[index] = u_name
		offset_arr[index] = sum_num
		sum_num += num
		index += 1
	}

	generator = func() string {
		award_num := rand.Int63n(sum_num)
		return name_arr[bsearch(offset_arr, award_num)]
	}
	return
}

func bsearch(nums []int64, target int64) int {
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
	return -1
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

	rand.Seed(time.Now().Unix())
	award_stat := make(map[string]int64)

	for i := 0; i < 10000; i++ {
		name := getgenerator(users)
		if count, ok := award_stat[name()]; ok {
			award_stat[name()] = count + 1
		} else {
			award_stat[name()] = 1
		}
	}

	for user, count := range award_stat {
		fmt.Printf("User : %s ,count : %d \n", user, count)
	}

}
