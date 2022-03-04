---
title: 211216key
weight: 1
bookHidden: true
---

### 211216作业

```
type BabyNames struct {
	keys  []int
	names []string
	nums  []int
}

// 初始化
func (b *BabyNames) Init(names []string, nums []int) {
	b.keys = make([]int, len(names))
	for i := 0; i < len(names); i++ {
		b.keys[i] = -1
	}
	b.names = names
	b.nums = nums
}

func trulyMostPopular(names []string, synonyms []string) []string {
	var Names []string
	var Nums []int
	var buf BabyNames
	namesMap := make(map[string]int, len(names))
	// 统计频率
	for k, v := range names {
		temp := strings.Split(v, "(")
		Names = append(Names, temp[0])
		num, _ := strconv.Atoi(temp[1][:len(temp[1])-1])
		Nums = append(Nums, num)
		namesMap[temp[0]] = k
	}
	buf.Init(Names, Nums)
	// 同义词
	for _, v := range synonyms {
		temp := strings.Split(v[1:len(v)-1], ",")
		// 寻找父
		key1, _ := namesMap[temp[0]]
		key2, _ := namesMap[temp[1]]
		buf.Merge(key1, key2)
	}

	var result []string
	for i, v := range buf.keys {
		if v < 0 {
			tmp := buf.names[i] + "(" + strconv.Itoa(buf.nums[i]) + ")"
			result = append(result, tmp)
		}
	}
	return result
}
func (b *BabyNames) Find(node int) int {
	if b.keys[node] < 0 {
		return node
	}
	tparent := b.Find(b.keys[node])
	b.keys[node] = tparent
	return tparent
}

func (b *BabyNames) Merge(first, second int) int {
	fp := b.Find(first)
	sp := b.Find(second)
	result := fp
	if fp == sp {
		return result
	}
	if b.names[fp] < b.names[sp] {
		b.keys[sp] = fp
		b.nums[fp] += b.nums[sp]
	} else {
		b.keys[fp] = sp
		b.nums[sp] += b.nums[fp]
		result = sp
	}
	return result
}
```