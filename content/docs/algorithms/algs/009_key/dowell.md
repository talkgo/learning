<!--
 * @Author: dowell87
 * @Date: 2021-12-12 22:48:51
 * @Descripttion: 
 * @LastEditTime: 2021-12-12 22:53:33
-->
---
title: 211209key
weight: 1
bookHidden: true
---
### 211209作业

```

type NestedInteger struct {
}

func (n *NestedInteger) SetInteger(value int) {}

func (n *NestedInteger) Add(elem NestedInteger) {}

func Deserialize(s string) *NestedInteger {
	// 只生成单字节
	num, numType := 0, 0
	var list *NestedInteger
	stack := []*NestedInteger{}
	// 单值
	if s[0] != '[' {
		n := NestedInteger{}
		num, _ = strconv.Atoi(s)
		n.SetInteger(num)
		return &n
	}
	// 多值
	for _, v := range s {
		switch v {
		case '[':
			// 插入一层
			if list != nil {
				stack = append(stack, list)
			}
			list = &NestedInteger{}
		case ']':
			// 结束
			if numType > 0 {
				elem := NestedInteger{}
				elem.SetInteger(num)
				list.Add(elem)
				num, numType = 0, 0
			}
			// 获取长度
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				parent.Add(*list)
				list = parent
				stack = stack[:len(stack)-1]
			}
		case ',':
			// 一个值的结束
			if list == nil {
				list = &NestedInteger{}
			}
			if numType > 0 {
				temp := NestedInteger{}
				temp.SetInteger(num)
				list.Add(temp)
				num, numType = 0, 0
			}
		case '-':
			numType = -1
		default:
			if numType == 0 {
				numType = 1
			}
			if num < 0 {
				num = num*10 - int(v-'0')
			} else {
				num = num*10 + int(v-'0')
			}
			if numType < 0 {
				num = -num
				numType = 1
			}
		}
	}
	return list
}

```