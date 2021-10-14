---
title: 第 014 期（2022.1.13）
weight: 1
bookHidden: true
---

```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 新树
func NewTreeNode(v int) *TreeNode {
	return &TreeNode{
		Val: v,
	}
}

// 创建树
func (node *TreeNode) CreateTreeNode(v int) *TreeNode {
	if node == nil {
		node = &TreeNode{
			Val: v,
		}
	} else if node.Val > v {
		node.Left = node.Left.CreateTreeNode(v)
	} else if node.Val < v {
		node.Right = node.Right.CreateTreeNode(v)
	}
	return node
}

// 先序：考察到一个节点后，即刻输出该节点的值，并继续遍历其左右子树。(根左右)
func (node *TreeNode) FrontTreeNode() {
	if node == nil {
		// fmt.Println("nil")
		return
	}
	node.Left.FrontTreeNode()
	fmt.Println(node.Val)
	node.Right.FrontTreeNode()
}

// 中序：考察到一个节点后，将其暂存，遍历完左子树后，再输出该节点的值，然后遍历右子树。(左根右)
func (node *TreeNode) MidTreeNode() {
	if node == nil {
		// fmt.Println("nil")
		return
	}
	fmt.Println(node.Val)
	node.Left.FrontTreeNode()
	node.Right.FrontTreeNode()

}

// 后序：考察到一个节点后，将其暂存，遍历完左右子树后，再输出该节点的值。(左右根)
func (node *TreeNode) BackTreeNode() {
	if node == nil {
		// fmt.Println("nil")
		return
	}
	node.Right.FrontTreeNode()
	fmt.Println(node.Val)
	node.Left.FrontTreeNode()
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val >= root.Val {
		return inorderSuccessor(root.Right, p)
	} else {
		left := inorderSuccessor(root.Left, p)
		if left == nil {
			return root
		}
		return left
	}
}
```