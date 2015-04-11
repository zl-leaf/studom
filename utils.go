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
 * @param node 需要剪枝的结点 complete 是否把无用结点（br、img等）去掉
 */
func CutStuDomTree(node *dom.Node, complete bool) {
    for _,child := range node.Child {
        CutStuDomTree(child, complete)
    }

    if node.Tag=="root" || node.Tag=="h1" || !dom.IsBlockDom(node.Tag) {
        // root h1 h2 h3 h4 h5 text标签不会被删除
        return
    }

    parent := node.BlockParent()
    if parent == nil || parent.Tag == "root" {
        return
    }
    if complete && dom.IsUselessDom(node.Tag) {
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
