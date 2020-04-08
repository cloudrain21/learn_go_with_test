package array_slice

func ArraySum(n [5]int) int {
	sum := 0
	for _, number := range n {
		sum += number
	}
	return sum
}

func SliceSum(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

func Sum(n []int) int {
	sum := 0
	for _, v := range n {
		sum += v
	}
	return sum
}

//func SumAll(n ...[]int) []int {
//	sliceNum := len(n)
//	sums := make([]int, sliceNum)
//	for i, v := range n {
//		sums[i] = Sum(v)
//	}
//	return sums
//}

// refactor
func SumAll(n ...[]int) []int {
	var sum []int
	for _, v := range n {
		sum = append(sum, Sum(v))
	}
	return sum
}

func SumAllTails(n ...[]int) []int {
	var sumTail []int
	s := 0
	for _, v := range n {
		if len(v) > 0 {
			s = Sum(v[1:])
		}
		sumTail = append(sumTail, s)
	}
	return sumTail
}