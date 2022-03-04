---
title: 第 015 期（2022.1.20）
weight: 1
bookHidden: true
---

```
type LRUCache struct {
	Link *LinkList
	Map  map[int]*LinkNode
}

type LinkNode struct {
	Key, Val  int
	Pre, Next *LinkNode
}

type LinkList struct {
	Head, Tail     *LinkNode
	Size, Capacity int
}

func Constructor(capacity int) LRUCache {
	list := &LinkList{Head: &LinkNode{}, Tail: &LinkNode{}, Capacity: capacity}
	list.Head.Next = list.Tail
	list.Tail.Pre = list.Head
	m := map[int]*LinkNode{}
	return LRUCache{Link: list, Map: m}
}

func (this *LRUCache) addToHead(node *LinkNode) {
	Head := this.Link.Head
	node.Pre = Head
	node.Next = Head.Next
	Head.Next.Pre = node
	Head.Next = node
}

func (this *LRUCache) removeNode(node *LinkNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) moveToHead(node *LinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Map[key]; ok {
		this.moveToHead(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Map[key]; ok {
		node.Val = value
		this.moveToHead(node)
		return
	}
	if this.Link.Size == this.Link.Capacity {
		lastNode := this.Link.Tail.Pre
		this.removeNode(lastNode)
		delete(this.Map, lastNode.Key)
		this.Link.Size--
	}
	newNode := &LinkNode{Key: key, Val: value}
	this.addToHead(newNode)
	this.Map[key] = newNode
	this.Link.Size++
}
```