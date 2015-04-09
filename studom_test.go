package studom

import(
	"testing"
	"fmt"

	"github.com/zl-leaf/studom/dom"
)

func Test_ParseString(t *testing.T) {
	htm := "<div>哈哈<div><p>haha</p></div><div><a>test</a></div></div>"
	root,_ := ParseString(htm)
	cutStuDomTree(root)
	print(root)
}

func print(node *dom.Node)  {
	fmt.Println("我的标签:"+node.Tag)
	fmt.Println("我的值:"+node.Text)

	fmt.Printf("非链接文本长度:%d\n",node.CountLength)
	fmt.Printf("链接数:%d\n",node.LinkCount)
	fmt.Printf("局部相关度:%f\n",node.LocalCorrelativity())
	fmt.Printf("上下文相关度:%f\n",node.ContextualCorrelativity())

	fmt.Println()
	for _,c := range node.Child {
		print(c)
	}
}
