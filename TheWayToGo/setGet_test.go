package TheWayToGo

import (
	"fmt"
	_ "fmt"
	"testing"
)

var p Person = Person{"Sergey", "Jay"}

func TestPerson_String(t *testing.T) {

}

func BenchmarkPerson_String(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//fmt.Println(p.String())
		fmt.Sprintf("%v", p.String())
	}
}
