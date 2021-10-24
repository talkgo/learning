package weekly

import (
	"sort"
)

func HK(box [][]int) int {
	b := Int2Slice(box)
	if b.Len() == 0 {
		return 0
	}

	sort.Sort(b)

	dp := make([]int, b.Len())
	for i, elem := range b {
		pre := 0
		for j := i - 1; j >= 0; j-- {
			if dp[j] > pre && b[j][0] < elem[0] && b[j][1] < elem[1] && b[j][2] < elem[2] {
				pre = dp[j]
			}
		}
		dp[i] = pre + elem[2]
	}

	max := dp[0]
	for _, val := range dp {
		if val > max {
			max = val
		}
	}
	return max
}

type Int2Slice [][]int

func (b Int2Slice) Len() int           { return len(b) }
func (b Int2Slice) Less(i, j int) bool { return b[i][2] < b[j][2] }
func (b Int2Slice) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
