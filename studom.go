package studom
import(
	"io"
	"strings"

	"github.com/opesun/goquery"
	"github.com/opesun/goquery/exp/html"

	"github.com/zl-leaf/studom/dom"
)

func Parse(rd io.Reader) (root *dom.Node, err error) {
	nodes,err := goquery.Parse(rd)
	if err != nil {
		return
	}
	root = parseStuDomTree(nodes)
	return
}

func ParseString(htm string) (root *dom.Node, err error) {
	nodes,err := goquery.ParseString(htm)
	if err != nil {
		return
	}
	root = parseStuDomTree(nodes)
	return
}

/**
 * 解析成stu-dom树
 */
func parseStuDomTree(ns goquery.Nodes) (root *dom.Node) {
	divs := ns.Find("body")
	divRoot := divs[0]

	root = &dom.Node{}
	root.Tag = "root"

	t := ns.Find("title")[0].Node
	stuTitleNode := parseStuDomTitle(t)
	root.Add(stuTitleNode)
	stuTitleNode.Parent = root

	// 获取body信息
	body := parseStuDomNode(divRoot.Node)
	root.Add(body)
	body.Parent = root

	return
}

func parseStuDomTitle(n *html.Node) (stuTitleNode *dom.Node)  {
	stuTitleNode = &dom.Node{}
	stuTitleNode.Tag = "title"

	for _,child := range n.Child {
		if child.Type == html.TextNode {
			stuNode := &dom.Node{}
			stuNode.Tag = "text"
			stuNode.Text = child.Data
			stuNode.Parent = stuTitleNode
			stuTitleNode.Add(stuNode)
		}
	}

	return
}

/**
 * 解析结点
 */

func parseStuDomNode(n *html.Node) (stuNode *dom.Node) {
	if n.Type==html.ElementNode && (n.Data=="script" || n.Data=="style" || n.Data=="noscript") {
		return
	}
	if n.Type == html.TextNode {
		// 文字
		text := strings.TrimSpace(n.Data)
		// 取出空格和中间的换行
		text = strings.Replace(text, " ", "", -1)
		text = strings.Replace(text, "\n", "", -1)
		if text == "" {
			return
		}

		stuNode = &dom.Node{}
		stuNode.Tag = "text"
		stuNode.Text = text

		stuNode.CountLength = len([]rune(text))
		stuNode.LinkCount = 0
	} else if n.Type == html.ElementNode {
		// HTML标签
		stuNode = &dom.Node{}

		isLink := false
		stuNode.Tag = n.Data
		if n.Data == "a" || n.Data == "button" || n.Data == "input" {
			stuNode.LinkCount = 1
			isLink = true
		}
		for _,child := range n.Child {
			childStuNode := parseStuDomNode(child)
			if childStuNode != nil {
				if !isLink {
					stuNode.CountLength += childStuNode.CountLength
				}
				stuNode.LinkCount += childStuNode.LinkCount
				stuNode.Add(childStuNode)
				childStuNode.Parent = stuNode
			}
		}

		if stuNode.Tag == "h1" {
			// h1标题的特殊处理
			stuNode.LinkCount = 0
		}
	}
	return
}
