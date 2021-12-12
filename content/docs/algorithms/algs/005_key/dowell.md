<!--
 * @Author: dowell87
 * @Date: 2021-12-06 09:20:41
 * @Descripttion: 
 * @LastEditTime: 2021-12-12 22:53:11
-->
---
title: 211111key
weight: 1
bookHidden: true
---

### 211111作业

* 示例代码

```
func robot(command string, obstacles [][]int, x int, y int) bool {
    newX, newY := 0, 0
	j := len(command)
	i := 0
	for {
		z := i % j
		if command[z:z+1] == "R" {
			newX++
		} else {
			newY++
		}
		// 到达目的地返回元素
		if newX == x && newY == y {
			return true
		}
		for w := 0; w < len(obstacles); w++ {
			// 匹配是否碰到障碍物
			if obstacles[w][0] == newX && obstacles[w][1] == newY {
				return false
			}
			// 已经超越障碍物从障碍物中删除
			if obstacles[w][0] < newX && obstacles[w][1] < newY {
				obstacles = append(obstacles[:w], obstacles[w+1:]...)
				w--
			}
		}
		// 超过终点终止运行
		if x < newX || y < newY {
			return false
		}
		i++
	}
}
```
