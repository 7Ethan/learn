package Language

import "fmt"

func try()  {
	const temp
	fmt.Println(temp)
}


func main() {
	//outer:
		for i := 0; i < 3 ;i++  {
			for j := 0; j < 3 ;j++  {
				print(i," ",j," ")
				break
				//break outer
			}
			println()
		}

		try()
}
