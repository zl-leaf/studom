package studom

import(
	"testing"
	"fmt"
	"os"

	"github.com/zl-leaf/studom/dom"
)

func Test_Parse(t *testing.T) {
	fi,err := os.Open("test.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	tree,err := Parse(fi)
	if err != nil {
		fmt.Println(err)
		return
	}
	tree.CutStuDomTree()
	// print(tree.StuDomNode)
	fmt.Println(tree.Text())
}

func print(node *dom.StuDomNode)  {
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
