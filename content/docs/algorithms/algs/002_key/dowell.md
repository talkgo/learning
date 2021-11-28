---
title: 211021key
weight: 1
bookHidden: true
---

* 示例代码

```
func pileBox(box [][]int) int {
    sort.Slice(box, func(i, j int) bool {
		return box[i][0] < box[j][0]
	})
	// 初始化多少堆
	sum := make([]int, len(box))
	for i := range sum {
		sum[i] = box[i][2]
	}
	for i := 0; i < len(box); i++ {
		temp := 0
		for j := 0; j < i; j++ {
			// 将后面的箱子往上面堆
			if sum[j] > temp && box[i][0] > box[j][0] && box[i][1] > box[j][1] && box[i][2] > box[j][2] {
				temp = sum[j]
			}
		}
		// 如果能对上去，高度++
		sum[i] = sum[i] + temp
	}
	// 对高度内容排序,获取最大的内容
	sort.Slice(sum, func(i, j int) bool {
		return sum[i] > sum[j]
	})
	return sum[0]
}
```