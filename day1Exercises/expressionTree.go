package day1Exercises

import "fmt"

type Node struct {
	val                   string
	leftChild, rightChild *Node
}

func (node *Node) preOrder() string {
	if node == nil {
		return ""
	}
	result := node.val
	result += node.leftChild.preOrder()
	result += node.rightChild.preOrder()
	return result
}

func (node *Node) postOrder() string {
	if node == nil {
		return ""
	}
	result := node.leftChild.postOrder()
	result += node.rightChild.postOrder()
	result += node.val

	return result
}

func ExpressionTree() {
	var node2, node4, node5 Node
	node2.val, node4.val, node5.val = "a", "b", "c"
	node3 := Node{"-", &node4, &node5}
	node1 := Node{"+", &node2, &node3}

	preOrder := node1.preOrder()
	fmt.Println("pre order: ", preOrder)
	postOrder := node1.postOrder()
	fmt.Println("post order: ", postOrder)

}
