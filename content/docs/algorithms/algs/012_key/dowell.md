---
title: 211230key
weight: 1
bookHidden: true
---

### 211230作业

```

type trieNode struct {
	char     string               // Unicode 字符
	isEnding bool                 // 是否是文件结尾
	children map[string]*trieNode // 该节点的子节点字典
}

type Trie struct {
	root *trieNode // 根节点指针
}

func removeSubfolders(folder []string) []string {
	myTree := NewTrie()
	sort.SliceStable(folder, func(i, j int) bool {
		return len(folder[i]) < len(folder[j])
	})
	// 添加字典
	for _, v := range folder {
		myTree.AddTree(v)
	}
	for k := 0; k < len(folder); k++ {
		if myTree.Find(folder[k]) == false {
			folder = append(folder[:k], folder[k+1:]...)
			k--
		}
	}
	return folder
}

// 初始化 Trie 树
func NewTrie() *Trie {
	// 初始化根节点
	trieNode := NewTrieNode("/")
	return &Trie{trieNode}
}

// 初始化 Trie 树节点
func NewTrieNode(char string) *trieNode {
	return &trieNode{
		char:     char,
		isEnding: false,
		children: make(map[string]*trieNode),
	}
}

// 往 Trie 树中插入一个单词
func (t *Trie) AddTree(folder string) {
	node := t.root // 获取根节点
	folder = folder[1:len(folder)]
	folders := strings.Split(folder, "/")
	for _, dir := range folders { // 以 Unicode 字符遍历该单词
		value, ok := node.children[dir] // 获取 code 编码对应子节点
		if !ok {
			// 不存在则初始化该节点
			value = NewTrieNode(string(dir))
			// 然后将其添加到子节点字典
			node.children[dir] = value
		}
		// 当前节点指针指向当前子节点
		node = value
	}
	node.isEnding = true // 一个单词遍历完所有字符后将结尾字符打上标记
}

// 在 Trie 树中查找一个单词
func (t *Trie) Find(folder string) bool {
	node := t.root
	folder = folder[1:len(folder)]
	folders := strings.Split(folder, "/")
	// 是否是根目录
	num := len(folder) //当前目录总共几层
	sum := 1           // 循环到了第几层
	for _, code := range folders {
		// 是否
		value, _ := node.children[code] // 获取对应子节点
		if sum == num && len(node.children) > 1 {
			value.isEnding = true
			return true
		}
		if node.isEnding == true {
			return false
		}
		node = value
		sum++
	}
	return true // 找到对应单词
}

```