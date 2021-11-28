---
title: 211125key
weight: 1
bookHidden: true
---

### 211125作业

* 示例代码

```
func WaysToChange(n int) int {
	// 假设有n+1个可能性
	dp := make([]int, n+1)
	// 零钱的集合
	coins := []int{1, 5, 10, 25}
	//刚好可以用一个硬币凑成的情况，是一种情况
	dp[0] = 1
	/**
	* dp方程：dp[i] += dp[i - coin];
	 */
	for _, coin := range coins {
		for i := coin; i <= n; i++ {
			dp[i] = (dp[i] + dp[i-coin]) % 1000000007
		}
	}
	return dp[n]
}

```
