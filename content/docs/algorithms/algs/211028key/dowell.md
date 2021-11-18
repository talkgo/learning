---
title: 1028key
weight: 1
bookHidden: true
---

### 201028作业

* 示例代码

```
func cherryPickup(grid [][]int) int {
 N := len(grid)
 M := (N << 1) - 1
 dp := make([][]int, N)
 for k, _ := range dp {
  dp[k] = make([]int, N)
 }
 dp[0][0] = grid[0][0]

 for n := 1; n < M; n++ {
  for i := N - 1; i >= 0; i-- {
   for p := N - 1; p >= 0; p-- {
    j, q := n-i, n-p

    if j < 0 || j >= N || q < 0 || q >= N || grid[i][j] < 0 || grid[p][q] < 0 {
     dp[i][p] = -1
     continue
    }

    if i > 0 {
     dp[i][p] = max(dp[i][p], dp[i-1][p])
    }

    if p > 0 {
     dp[i][p] = max(dp[i][p], dp[i][p-1])
    }
    if i > 0 && p > 0 {
     dp[i][p] = max(dp[i][p], dp[i-1][p-1])
    }
    if dp[i][p] >= 0 {

     if i != p {
      dp[i][p] = dp[i][p] + grid[i][j] + grid[p][q]
     } else {
      dp[i][p] = dp[i][p] + grid[i][j]
     }
    }

   }
  }
 }

 return max(dp[N-1][N-1], 0)
}

func max(a, b int) int {
 if a > b {
  return a
 }
 return b

}

```