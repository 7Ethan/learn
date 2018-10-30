package language

import (
	"fmt"
	"math/rand"
	"time"
)

func GetAwardUserName(users map[string]int64) (name string) {
	var sum_num int64
	name_arr := make([]string, len(users))
	for u_name, num := range users {
		sum_num += num
		name_arr = append(name_arr, u_name)
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
		"b": 6,
		"c": 3,
		"d": 12,
		"e": 20,
		"f": 1,
	}

	rand.Seed(time.Now().Unix())
	award_stat := make(map[string]int64)
	for i := 0; i < 1000; i += 1 {
		name := GetAwardUserName(users)
		if count, ok := award_stat[name]; ok {
			award_stat[name] = count + 1
		} else {
			award_stat[name] = 1
		}
	}

	for name, count := range award_stat {
		fmt.Printf("user: %s, award count: %d\n", name, count)
	}

	return
}
