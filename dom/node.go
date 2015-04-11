package dom

type Node struct {
	LinkCount int
	CountLength int
	Text string
	Tag string
	Child []*Node
	Parent *Node
}

/**
 * 获取分快父元素
 */
func (node *Node) BlockParent() (blockParent *Node) {
	parent := node.Parent
	for {
		if parent==nil || parent.Tag == "root" || IsBlockDom(parent.Tag) {
			blockParent = parent
			return
		}
		parent = parent.Parent
	}
	return
}

func (node *Node) Add(element *Node) {
	node.Child = append(node.Child, element)
}

func (node *Node) Remove(element *Node) {
	tmpNode := element
	for {
		if tmpNode.Parent == node {
			break
		}
		tmpNode = tmpNode.Parent
	}
	tempChild := make([]*Node, 0)
	for _,child := range node.Child {
		if child != tmpNode {
			tempChild = append(tempChild, child)
		}
	}
	node.Child = tempChild
}
