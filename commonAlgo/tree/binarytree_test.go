package tree

import (
	"fmt"
	"testing"
)

var tree BinarySearchTree

func initTree(tree *BinarySearchTree) {
	tree.Insert(8, "8")
	tree.Insert(4, "4")
	tree.Insert(10, "10")
	tree.Insert(2, "2")
	tree.Insert(6, "6")
	tree.Insert(1, "1")
	tree.Insert(3, "3")
	tree.Insert(5, "5")
	tree.Insert(7, "7")
	tree.Insert(9, "9")
}

func printFunc(i Item) {
	fmt.Printf("\t %v", i)
}

func TestInsert(t *testing.T) {
	initTree(&tree)
	tree.String()
	tree.Insert(11, "11")
	tree.String()

	if tree.Max() != "11" || tree.Min() != "1" {
		t.Errorf("max of tree =  %s, min of tree = %s", tree.Max(), tree.Min())
	}
}

func TestTraverse(t *testing.T) {
	initTree(&tree)
	fmt.Println("+++先序遍历，根节点-->左子树-->右子树+++")
	tree.PreOrderTraverse(printFunc)
	fmt.Println()
	fmt.Println("+++中序遍历，左子树-->根节点-->右子树+++")
	tree.PostOrderTraverse(printFunc)
	fmt.Println()
	fmt.Println("+++后序遍历，左子树-->右子树-->根节点+++")
	tree.InOrderTraverse(printFunc)
	fmt.Println()
}

func TestRemoveNode(t *testing.T) {
	initTree(&tree)
	tree.Remove(1)
	tree.Remove(3)
	tree.String()

	item, ok := tree.Search(10)
	if ok {
		printFunc(item)
	} else {
		t.Error("Seach Error")
	}
}
