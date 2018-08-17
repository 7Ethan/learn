package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardUser(users map[string]int64) (name string) {
	var sum_num int64
	var name_arr = make([]string, len(users))
	for name, num := range users {
		sum_num += num
		name_arr = append(name_arr, name)
	}

	award_num := rand.Int63n(sum_num)
	var offset_num int64
	for _, u_name := range name_arr {
		offset_num += users[u_name]
		if award_num < offset_num {
			name = u_name
			return
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
		name := awardUser(users)
		if count, ok := award_stat[name]; ok {
			award_stat[name] = count + 1
		} else {
			award_stat[name] = 1
		}
	}

	for user, num := range award_stat {
		fmt.Printf("User : %s ,award counr : %d \n", user, num)
	}
}
