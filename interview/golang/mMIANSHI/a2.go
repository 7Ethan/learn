package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getawarduser(users map[string]int64) (name string) {
	var sum_num int64
	name_arr := make([]string, len(users))

	for u_name, count := range users {
		name_arr = append(name_arr, u_name)
		sum_num += count
	}

	award_num := rand.Int63n(sum_num)
	var offset_num int64
	for _, u_name := range name_arr {
		offset_num += sum_num
		if award_num < offset_num {
			name = u_name
		}
	}
	return
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

	for i := 0; i < 1000; i++ {
		name := getawarduser(users)
		if count, ok := award_stat[name]; ok {
			award_stat[name] = count + 1
		} else {
			award_stat[name] = 1
		}
	}

	for users, count := range award_stat {
		fmt.Printf(" Users is :%s ,count is : %d \n", users, count)
	}
}
