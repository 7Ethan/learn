package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	str := `{{"userName":"张三","Age":19},{"userName":"张三","Age":19}}`
	var m []map[string]interface{}
	json.Unmarshal([]byte(str), &m)
	fmt.Println(m)
	//for k,v := range m{
	//	switch v.(type) {
	//	case float64:
	//		fmt.Println(k, " 是 int 类型,值为:", v)
	//	case string:
	//		fmt.Println(k, " 是 string 类型,值为:", v)
	//	default:
	//		fmt.Println(k, "无法误用别的类型")
	//	}
	//}
}
