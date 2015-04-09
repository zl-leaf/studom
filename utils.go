package studom

import(
    "github.com/zl-leaf/studom/dom"
)

const(
	Cm = 0
	Lcm = 0.02
	Ccm = 0.001
)

/**
 * 剪枝
 */
func cutStuDomTree(node *dom.Node) {
    for _,child := range node.Child {
        cutStuDomTree(child)
    }

    if node.Tag == "root" {
        return
    }

    parent := node.Parent
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
