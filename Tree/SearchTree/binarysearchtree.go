package searchtree

import "fmt"

// type constraints for number tree
type Number interface {
	~int | ~float32 | ~float64
}

// struct for Node
type Node[T Number] struct {
	key   T
	left  *Node[T]
	right *Node[T]
}

// struct for tree
type BST[T Number] struct {
	root *Node[T]
}

// init for object creation
func New[T Number]() *BST[T] {
	return &BST[T]{}
}

// create new node oject
func (b *BST[T]) newNode(value T) *Node[T] {
	return &Node[T]{key: value}
}

// insert new item in the tree
func (b *BST[T]) Insert(val T) {
	if b.root == nil {
		b.root = b.newNode(val)
		return
	}
	currNode := b.root
	prevNode := b.root
	for currNode != nil {
		if currNode.key > val {
			prevNode = currNode
			currNode = currNode.left
		} else if currNode.key < val {
			prevNode = currNode
			currNode = currNode.right
		} else {
			// duplicate not allowed, hence simply return
			return
		}
	}
	if prevNode.key > val {
		prevNode.left = b.newNode(val)
	} else {
		prevNode.right = b.newNode(val)
	}
}

// preorder helper method
func (b *BST[T]) preOrderHelper(root *Node[T]) {
	if root != nil {
		fmt.Printf("%v ", root.key)
		b.preOrderHelper(root.left)
		b.preOrderHelper(root.right)
	}
}

// preorder tree traversal root->left->right
func (b *BST[T]) PreOrder() {
	b.preOrderHelper(b.root)
	fmt.Println()
}

// inorder helper method
func (b *BST[T]) inOrderHelper(root *Node[T]) {
	if root != nil {
		b.inOrderHelper(root.left)
		fmt.Printf("%v ", root.key)
		b.inOrderHelper(root.right)
	}
}

// inorder tree traversal left->root->right
func (b *BST[T]) InOrder() {
	b.inOrderHelper(b.root)
	fmt.Println()
}

// postorder helper method
func (b *BST[T]) postOrderHelper(root *Node[T]) {
	if root != nil {
		b.postOrderHelper(root.left)
		b.postOrderHelper(root.right)
		fmt.Printf("%v ", root.key)
	}
}

// postorder tree traversal left->right->root
func (b *BST[T]) PostOrder() {
	b.postOrderHelper(b.root)
	fmt.Println()
}

// search value in the tree and return boolean as result
func (b *BST[T]) Search(value T) bool {
	if b.root == nil {
		return false
	}
	currNode := b.root
	for currNode != nil {
		if currNode.key > value {
			currNode = currNode.left
		} else if currNode.key < value {
			currNode = currNode.right
		} else {
			return true
		}
	}
	return false
}
