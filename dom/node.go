package dom

type Node struct {
	LinkCount int
	CountLength int
	Text string
	Tag string
	Child []*Node
	Parent *Node
}

func (node *Node) Add(element *Node) {
	node.Child = append(node.Child, element)
}

func (node *Node) Remove(element *Node) {
	tempChild := make([]*Node, 0)
	for _,child := range node.Child {
		if child != element {
			tempChild = append(tempChild, child)
		} else {
			node.LinkCount -= element.LinkCount
			node.CountLength -= element.CountLength
		}
	}
	node.Child = tempChild
}
