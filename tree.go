package tracex

import "encoding/json"

type Node struct {
	*TraceX
	Children []*Node `json:"children,omitempty"`
}

func NewTree(root *Node, nodes []*Node) *Node {
	return makeTree(nodes, root)
}
func (n *Node) Marshal() []byte {
	b, _ := json.MarshalIndent(n, "", "  ")
	return b
}
func makeTree(nodes []*Node, root *Node) *Node {
	childs, _ := haveChildren(nodes, root)
	if childs != nil {
		root.Children = append(root.Children, childs[0:]...)
		for _, v := range childs {
			_, has := haveChildren(nodes, v)
			if has {
				makeTree(nodes, v)
			}
		}
	}
	return root
}

func haveChildren(Trees []*Node, node *Node) (childs []*Node, yes bool) {
	for _, v := range Trees {
		if v.TracePid == node.TraceId {
			childs = append(childs, v)
		}
	}
	if childs != nil {
		yes = true
	}
	return
}
