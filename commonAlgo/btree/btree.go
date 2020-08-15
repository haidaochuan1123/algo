package btree

type Node struct {
	children []*Node
}

type BTree struct {
	degree int
	length int
	root   *Node
}
