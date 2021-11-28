---
title: 211118key
weight: 1
bookHidden: true
---

### 211118作业

* 示例代码

```
func ConcatenatedBinary(n int) int {
	length, res := 0, 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			length++
		}
		res = (res<<length | i) % (1e9 + 7)
	}
	return res
}

```
