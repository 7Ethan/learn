package mySort

type MySorter interface {
	Lenn() int
	Less(i, j int) bool
	Swap(i, j int)
}

func MySort(data MySorter) {
	for pass := 1; pass < data.Lenn(); pass++ {
		for i := 0; i < data.Lenn()-pass; i++ {
			if data.Less(pass, i) {
				data.Swap(pass, i)
			}
		}
	}
}

//
//func IsSorted(data MySorter)bool  {
//
//}
//
//type MySlice []int
//
//func (m MySlice)lenn()int  {
//	return len(m)
//}
//
//func (m MySlice)less(i,j int)bool  {
//	return m[i] < m[j]
//}
//
//func (m MySlice)swap(i,j int)  {
//	m[i] , m[j] = m[j] , m[i]
//}
