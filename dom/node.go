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

/**
 * 获取stu-dom树的内容
 */
func (node *StuDomNode) AllText() string {
    if node.Tag == "text" {
        return node.Text
    }
    content := ""
    for _,child := range node.Child {
        line := child.AllText()
        if IsBlockDom(child.Tag) {
            line += "\n"
        }
        content += line
    }
    return content
}

/**
 * 局部相关度
 */
func (node *StuDomNode) LocalCorrelativity() float32 {
    if node.CountLength == 0 {
        return float32(1)
    }
    return float32(node.LinkCount)/float32(node.CountLength)
}

/**
 * 上下文相关度
 */
func (node *StuDomNode) ContextualCorrelativity() float32 {
    parent := node.BlockParent()

    if parent != nil && parent.Tag != "root" {
        if parent.CountLength == 0 {
            return float32(1)
        }
        return float32(node.LinkCount)/float32(parent.CountLength)
    }
    return -1
}
