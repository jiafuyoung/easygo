package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 将一维的数组转成二维数组，每层是上层的两倍
func Build2dArray(arr []int) [][]int {

	twoDarray := make([][]int, 0)
	count := len(arr)
	if count == 0 {
		return twoDarray
	}
	left := 0
	level := 1
	for len(arr[left:]) > 0 {
		right := 2 ^ level - 1
		if right > count {
			right = count
		}
		t_arr := arr[left:right]
		twoDarray = append(twoDarray, t_arr)
		left = right + 1
		level++
	}
	return twoDarray
}
