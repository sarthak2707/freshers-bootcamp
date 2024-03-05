package day1

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

func (n *Node) InsertLeft(val string) {
	newNode := Node{Val: val, Left: nil, Right: nil}
	n.Left = &newNode
}

func (n *Node) InsertRight(val string) {
	newNode := Node{Val: val, Left: nil, Right: nil}
	n.Right = &newNode
}

func (n *Node) PrintPreOrderTraversal() {
	if n.Val != "" {
		print(n.Val, " ")
	}
	if n.Left != nil {
		n.Left.PrintPreOrderTraversal()
	}
	if n.Right != nil {
		n.Right.PrintPreOrderTraversal()
	}
}

func (n *Node) PrintPostOrderTraversal() {
	if n.Left != nil {
		n.Left.PrintPostOrderTraversal()
	}
	if n.Right != nil {
		n.Right.PrintPostOrderTraversal()
	}
	if n.Val != "" {
		print(n.Val, " ")
	}
}
