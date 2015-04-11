package dom

var blockHtmlDom = []string{"title" , "body", "div", "p", "ul", "ol", "h1", "ul", "h2", "h3", "h4", "h5"}
var uselessHtmlDom = []string{"img", "br", "hr"}

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
 * 判断是否属于无用标签
 */
func IsUselessDom(tag string) bool {
    for _, d := range uselessHtmlDom {
        if d == tag {
            return true
        }
    }

    return false
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
