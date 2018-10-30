package language

import (
	"fmt"
	"math/rand"
	"time"
)

func GetAwardGenerator(users map[string]int64) (generator func() string) {
	var sum_num int64
	name_arr := make([]string, len(users))
	for u_name, num := range users {
		sum_num += num
		name_arr = append(name_arr, u_name)
	}

	generator = func() string {
		award_num := rand.Int63n(sum_num)

		var offset_num int64
		for _, u_name := range name_arr {
			offset_num += users[u_name]
			if award_num < offset_num {
				return u_name
			}
		}
		// 缺省返回，正常情况下，不会运行到此处
		return name_arr[0]
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
	for i := 0; i < 10000; i += 1 {
		name := GetAwardGenerator(users)
		if count, ok := award_stat[name()]; ok {
			award_stat[name()] = count + 1
		} else {
			award_stat[name()] = 1
		}
	}

	for name, count := range award_stat {
		fmt.Printf("user: %s, award count: %d\n", name, count)
	}

	return
}
