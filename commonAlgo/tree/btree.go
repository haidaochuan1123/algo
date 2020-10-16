package tree

// BtreeNode BTree节点
type BtreeNode struct {
	children []*BtreeNode
}

// BTree btree
type BTree struct {
	degree int
	length int
	root   *BtreeNode
}
