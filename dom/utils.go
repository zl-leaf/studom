package dom

var blockHtmlDom = []string{"title" , "body", "div", "p", "ul", "ol", "h1", "ul", "h2", "h3", "h4", "h5"}
var uselessHtmlDom = []string{"img", "br", "hr"}

const(
	Cm = 0
	Lcm = 0.02
	Ccm = 0.001
)

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
 * 剪枝
 */
func (node *StuDomNode) CutStuDomTree() {
    for _,child := range node.Child {
        child.CutStuDomTree()
    }

    if node.Tag=="root" || node.Tag=="h1" || !IsBlockDom(node.Tag) {
        // root h1 h2 h3 h4 h5 text标签不会被删除
        return
    }

    parent := node.BlockParent()
    if parent == nil || parent.Tag == "root" {
        return
    }
    if IsUselessDom(node.Tag) {
        parent.Remove(node)
    }

    if node.CountLength <= Cm {
        // 空结点
        parent.Remove(node)
        return
    }

    if node.LocalCorrelativity()>=Lcm && node.ContextualCorrelativity()>=Ccm {
        // 局部不相关
        parent.Remove(node)
        return
    }

}
