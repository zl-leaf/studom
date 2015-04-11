package dom

type StuDomNode struct {
	LinkCount int
	CountLength int
	Text string
	Tag string
	Child []*StuDomNode
	Parent *StuDomNode
}

/**
 * 获取分快父元素
 */
func (node *StuDomNode) BlockParent() (blockParent *StuDomNode) {
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

func (node *StuDomNode) Add(element *StuDomNode) {
	node.Child = append(node.Child, element)
}

func (node *StuDomNode) Remove(element *StuDomNode) {
	tmpNode := element
	for {
		if tmpNode.Parent == node {
			break
		}
		tmpNode = tmpNode.Parent
	}
	tempChild := make([]*StuDomNode, 0)
	for _,child := range node.Child {
		if child != tmpNode {
			tempChild = append(tempChild, child)
		}
	}
	node.Child = tempChild
}
