package day1

type Node struct {
	Left  *Node
	Right *Node
	Data  string
}

func (n *Node) InsertLeftNode(data string) {
	n.Left = &Node{Data: data}
}

func (n *Node) InsertRightNode(data string) {
	n.Right = &Node{Data: data}
}

func (n *Node) PrintPreOrderTraversal() {
	if n == nil {
		return
	}
	print(n.Data, " ")
	n.Left.PrintPreOrderTraversal()
	n.Right.PrintPreOrderTraversal()
}

func (n *Node) PrintPostOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.PrintPostOrderTraversal()
	n.Right.PrintPostOrderTraversal()
	print(n.Data, " ")
}
