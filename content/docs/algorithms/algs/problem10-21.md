---
title: 10.22题
weight: 2
---

### 10.22题

#### 题目描述:

```

堆箱子。给你一堆n个箱子，箱子宽 `wi`、深 `di`、高 `hi`。

箱子不能翻转，将箱子堆起来时，下面箱子的宽度、高度和深度必须大于上面的箱子。

实现一种方法，搭出最高的一堆箱子。箱堆的高度为每个箱子高度的总和。

输入使用数组[wi, di, hi]表示每个箱子。

```

#### 示例

示例1

```
 输入：box = [[1, 1, 1], [2, 2, 2], [3, 3, 3]]
 输出：6

```
示例2

```
输入：box = [[1, 1, 1], [2, 3, 4], [2, 6, 7], [3, 4, 5]]
输出：10

```


* 提示

```
箱子的数目不大于3000个。
```

#### 解题思路

做过 [俄罗斯套娃信封问题](https://leetcode-cn.com/problems/russian-doll-envelopes/) 或 [马戏团人塔](https://leetcode-cn.com/problems/circus-tower-lcci/) 的同学应该对这道题不陌生，这题只是在愿题的基础上加了一个维度。

我们考虑用动态规划求解， 我们先对box数组按照一个维度`宽（wi）`进行排序，得到新的box序列，那么最终答案的序列就一定是box序列的某个子序列（这点很重要）。那么之后我们只需要找到这个总高度最大的子序列。
设dp[i]表示以第i个箱子为结尾的上升子序列的最大总高度。则可以知道：

{{< katex  class="text-center"  >}}
dp[i]=box[i][2]+max(dp[j]) 

0<=j<i
{{< /katex >}}

时间复杂度：
{{< katex  class="text-center"  >}}
O(N^2)
{{< /katex >}}


```golang
func pileBox(box [][]int) int {
	//先对box排序
	sort.Slice(box, func(i, j int) bool {
		return box[i][0] < box[j][0]
	})
	dp := make([]int, len(box))
	//初始化每一项dp[i]为当前box的高度
	for i := 0; i < len(box); i++ {
		dp[i] = box[i][2]
	}
	max := dp[0]
	for i := 0; i < len(box); i++ {
		temp := -1
		for j := 0; j < i; j++ {
            //遍历j<i 的box逐个判断当前box[i]是否能放在box[j]上
			if dp[j] > temp&&box[i][0]>box[j][0]&&box[i][1]>box[j][1]&&box[i][2]>box[j][2] {
				temp = dp[j]
			}
		}
		if temp >= 0 {
			dp[i] = dp[i] +temp
			if dp[i] > max {
				max = dp[i]
			}
		}else {
            if dp[i] > max {
				max = dp[i]
			}
        }
	}
	return max
}
```
