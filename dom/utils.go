package dom

var blockHtmlDom = []string{"div", "p", "h1", "ul", "h2", "h3", "h4", "h5"}

/**
 * 局部相关度
 */
func (node *Node) LocalCorrelativity() float32 {
    return float32(node.LinkCount)/float32(node.CountLength)
}

/**
 * 上下文相关度
 */
func (node *Node) ContextualCorrelativity() float32 {
    parent := node.Parent
    if parent != nil {
        return float32(node.LinkCount)/float32(parent.CountLength)
    }
    return -1
}

/**
 * 判断是否属于分快标签
 */
func IsBlockDom(tag string) bool {
    for _, d := range blockHtmlDom {
        if d == tag {
            return true
        }
    }

    return false
}

/**
 * 获取stu-dom树的内容
 */
func (node *Node) Content() string {
    if node.Tag == "text" {
        return node.Text
    }
    content := ""
    for _,child := range node.Child {
        line := child.Content()
        if IsBlockDom(child.Tag) {
            line += "\r\n"
        }
        content += line
    }
    return content
}
