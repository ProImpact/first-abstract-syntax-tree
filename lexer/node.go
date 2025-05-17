package lexer

import (
	"bytes"
	"fmt"
	"log"
)

type Node struct {
	data   string
	childs []*Node
}

func newNode(data string) *Node {
	return &Node{
		data: data,
	}
}

func (n *Node) addNode(newNode *Node) {
	if n.childs == nil {
		n.childs = make([]*Node, 0)
	}
	n.childs = append(n.childs, newNode)
}

func (n *Node) String() string {
	buff := bytes.Buffer{}
	_, err := buff.WriteString(fmt.Sprintln(n.data))
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range n.childs {
		_, err = buff.WriteString(v.String())
		if err != nil {
			log.Fatal(err)
		}
	}
	return buff.String()
}
