---
title: 211202key
weight: 1
bookHidden: true
---

### 211202作业

* 示例代码

```

func rob(nums []int) int {
    prev := 0;
    curr := 0;

    // 每次循环，计算“偷到当前房子为止的最大金额”
    for _,i := range nums {
        // 循环开始时，curr 当前一次金额，prev 表示 前前一次的金额
        // dp[k] = max{ dp[k-1], dp[k-2] + i }
        temp := max(curr, prev + i);
        prev = curr;
        curr = temp;
        // 循环结束时，curr 表示当前金额，prev 表示 前一次金额
    }

    return curr
}

func max(a,b int) int{
    if(a >b){
        return a
    }
    return b
}

```
