---
title: 10.14题
weight: 1
---

### 10.14题

#### 题目描述:

```

 给定一个只包含数字的字符串，用以表示一个 IP 地址，
 返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
 有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，
 且不能含有前导 0），整数之间用 '.' 分隔。
 例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，
 但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。
 请使用回溯算法解题。

```

#### 示例

示例1

```
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]

```
示例2

```
输入：s = "0000"
输出：["0.0.0.0"]

```

示例3

```
输入：s = "1111"
输出：["1.1.1.1"]

```

示例4

```
输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

```

* 提示

```

0 <= s.length <= 3000
s 仅由数字组成的字符串

```

#### 解题思路

* 回溯算法(伪代码)

<<<<<<< HEAD
```

=======
>>>>>>> 6911083c1f597ba2916712b1e67d7c41e49ae96e
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return
    
    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择

<<<<<<< HEAD
```
=======
>>>>>>> 6911083c1f597ba2916712b1e67d7c41e49ae96e

* 非标准示例

```

func RestoreIpAddresses(s string) []string {
	ips := []string{}
	if len(s) < 4 || len(s) > 12 {
		return ips
	}

	var dfs func(ipAddr []string, str string)
	dfs = func(ipAddr []string, str string) {
		// 结束的状态
		lenAddr := len(ipAddr)
		if lenAddr == 4 && len(str) == 0 {
			ips = append(ips, strings.Join(ipAddr, "."))
			return
		}
		// 循环内容
		for i := 0; i < 3; i++ {
			if len(str)-(i+1) >= 3-lenAddr && len(str)-(i+1) <= (3-lenAddr)*3 {
				// 多位数字前面有0的情况
				if i >= 1 {
					tempnum, _ := strconv.Atoi(str[:1])
					if tempnum == 0 {
						break
					}
				}
                // 对数字的判断,大于255的情况
				if i == 2 {
					tempnum, _ := strconv.Atoi(str[0 : i+1])
					if tempnum > 255 {
						break
					}
				}
				tempAddr := append(ipAddr, str[0:i+1])
				tempstr := str[i+1 : len(str)]
				dfs(tempAddr, tempstr)
			}
		}
	}
	dfs(ips, s)
	return ips
}

```
