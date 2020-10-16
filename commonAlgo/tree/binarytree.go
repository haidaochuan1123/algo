package tree

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item 数据类型
type Item generic.Type

// BinaryNode 二叉树结构体
type BinaryNode struct {
	key   int
	value Item
	left  *BinaryNode
	right *BinaryNode
}

// BinarySearchTree 定义二叉树搜索树
type BinarySearchTree struct {
	root *BinaryNode
	lock sync.RWMutex
}

// Insert 向树中插入元素
func (tree *BinarySearchTree) Insert(key int, value Item) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	newNode := &BinaryNode{key, value, nil, nil}
	if tree.root == nil {
		tree.root = newNode
	} else {
		insertNode(tree.root, newNode)
	}
}

// 对比节点key，按顺序插入值
func insertNode(node, newNode *BinaryNode) {
	if newNode.key < node.key {
		// 插入左子树
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		// 插入右子树
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// Search 从树中检索元素
func (tree *BinarySearchTree) Search(key int) (Item, bool) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return search(tree.root, key)
}

func search(node *BinaryNode, key int) (Item, bool) {
	if node == nil {
		return nil, false
	}

	if key < node.key {
		// 搜索左子树
		return search(node.left, key)
	} else if key > node.key {
		// 搜索右子树
		return search(node.right, key)
	} else {
		return node.value, true
	}
}

// Remove 从树中删除节点
func (tree *BinarySearchTree) Remove(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	remove(tree.root, key)
}

func remove(node *BinaryNode, key int) *BinaryNode {
	// 要删除节点不存在
	if node == nil {
		return nil
	}

	// 寻找要删除的节点
	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}

	// 判断节点类型，如节点为叶子节点，直接删除
	if node.left == nil && node.right == nil {
		node = nil
		return node
	}

	// 要删除的节点只有一个自节点，删除其自身
	if node.left == nil {
		node = node.right
		return node
	}

	if node.right == nil {
		node = node.left
		return node
	}

	// 要删除的节点有两个子节点，找到右子树的最左节点，替换为当前节点
	mostLeftNode := node.right
	for {
		if mostLeftNode != nil && mostLeftNode.left != nil {
			mostLeftNode = mostLeftNode.left
		} else {
			break
		}
	}

	// 使用右子树的最左节点替换为当前节点，即删除当前节点
	node.key, node.value = mostLeftNode.key, mostLeftNode.value
	node.right = remove(node.right, key)
	return node
}

// Min 获取二叉树的最左节点
func (tree *BinarySearchTree) Min() Item {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	node := tree.root
	if node == nil {
		return nil
	}

	for {
		if node.left == nil {
			return node.value
		}
		node = node.left
	}
}

// Max 获取二叉树的最右节点
func (tree *BinarySearchTree) Max() Item {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	node := tree.root
	if node == nil {
		return nil
	}

	for {
		if node.right == nil {
			return node.value
		}
		node = node.right
	}
}

// PreOrderTraverse 先序遍历，根节点-->左子树-->右子树
func (tree *BinarySearchTree) PreOrderTraverse(printFunc func(Item)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	preOrderTraverse(tree.root, printFunc)
}

func preOrderTraverse(node *BinaryNode, printFunc func(Item)) {
	if node != nil {
		printFunc(node.value)
		preOrderTraverse(node.left, printFunc)
		preOrderTraverse(node.right, printFunc)
	}
}

// PostOrderTraverse 中序遍历，左子树-->根节点-->右子树
func (tree *BinarySearchTree) PostOrderTraverse(printFunc func(Item)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	postOrderTraverse(tree.root, printFunc)
}

func postOrderTraverse(node *BinaryNode, printFunc func(Item)) {
	if node != nil {
		preOrderTraverse(node.left, printFunc)
		printFunc(node.value)
		preOrderTraverse(node.right, printFunc)
	}
}

// InOrderTraverse 后序遍历，左子树-->右子树-->根节点
func (tree *BinarySearchTree) InOrderTraverse(printFunc func(Item)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverse(tree.root, printFunc)
}

func inOrderTraverse(node *BinaryNode, printFunc func(Item)) {
	if node != nil {
		preOrderTraverse(node.left, printFunc)
		preOrderTraverse(node.right, printFunc)
		printFunc(node.value)
	}
}

// 打印树结构
func (tree *BinarySearchTree) String() {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	// postOrderTraverse(tree.root, printFunc)
	if tree.root == nil {
		fmt.Println("Tree Is Empty")
		return
	}
	stringify(tree.root, 0)
	fmt.Println("------------------")
}

func stringify(node *BinaryNode, level int) {
	if node == nil {
		return
	}

	format := ""
	for i := 0; i < level; i++ {
		format += "\t"
	}

	format += "----[ "
	level++
	stringify(node.left, level)
	fmt.Printf(format+"%d\n", node.key)
	stringify(node.right, level)
}
