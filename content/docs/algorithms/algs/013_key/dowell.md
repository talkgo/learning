---
title: 第 013 期（2022.1.6）
weight: 1
bookHidden: true
---

```
// 节点
type BstNode struct {
	Key   int
	Value int
	left  []*BstNode
	right []*BstNode
}

type BstTree struct {
	node []*BstNode
}

func newBstNode(k, v int) *BstNode {
	return &BstNode{
		Key:   k,
		Value: v,
	}
}

// 添加
func (b *BstTree) InsertBstTree(node *BstNode) {
	b.node = append(b.node, node)
}

func (b *BstTree) InsertBstNode(nodeI, nodeJ *BstNode) {
	for i := 0; i < len(b.node); i++ {
		if b.node[i].Key == nodeI.Key {
			if b.node[i].Key < nodeJ.Key && b.node[i].Value > 2*nodeJ.Value {
				b.node[i].right = append(b.node[i].right, nodeJ)
			} else {
				b.node[i].left = append(b.node[i].left, nodeJ)
			}
		}
	}
}

// 统计
func (b *BstTree) SumBstTree() (sum int) {
	for i := 0; i < len(b.node); i++ {
		sum = sum + len(b.node[i].right)

	}
	return sum
}

func reversePairs(nums []int) int {
	var bt BstTree
	sum := len(nums)

	for i := 0; i < sum; i++ {
		bst := newBstNode(i, nums[i])
		bt.InsertBstTree(bst)
	}
	for i := 0; i < sum; i++ {
		for j := 0; j < sum; j++ {
			if i == j {
				continue
			}
			bsti := newBstNode(i, nums[i])
			bstj := newBstNode(j, nums[j])
			bt.InsertBstNode(bsti, bstj)
		}
	}
	return bt.SumBstTree()
}
```