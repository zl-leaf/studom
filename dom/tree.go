package dom

type StuDomTree struct {
    *StuDomNode
}

/**
 * 剪枝
 */
func (tree *StuDomTree) CutStuDomTree()  {
    tree.StuDomNode.CutStuDomTree()
}

/**
 * 获取树的title内容
 */
func (tree *StuDomTree) Title() string {
    for _,stuNode := range tree.Child {
        if stuNode.Tag == "title" {
            return stuNode.AllText()
        }
    }
    return ""
}

/**
 * 获取树的body内容
 */
func (tree *StuDomTree) Body() string {
    for _,stuNode := range tree.Child {
        if stuNode.Tag == "body" {
            return stuNode.AllText()
        }
    }
    return ""
}

func (tree *StuDomTree) Text() string {
    return tree.StuDomNode.AllText()
}
