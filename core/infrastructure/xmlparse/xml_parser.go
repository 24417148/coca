package xmlparse

import (
	"encoding/xml"
	"io"
	"strings"
)

type ElemType string

const (
	eleTpText ElemType = "text" // 静态文本节点
	eleTpNode ElemType = "XmlNode" // 节点子节点
)

type XmlNode struct {
	Id       string
	Name     string
	Attrs    map[string]xml.Attr
	Elements []element
}

type element struct {
	ElementType ElemType
	Val         interface{}
}

func ParseXml(r io.Reader) *XmlNode {
	parser := xml.NewDecoder(r)
	var root XmlNode

	st := NewStack()
	for {
		token, err := parser.Token()
		if err != nil {
			break
		}
		switch t := token.(type) {
		case xml.StartElement: //tag start
			elmt := xml.StartElement(t)
			name := elmt.Name.Local
			attr := elmt.Attr
			attrMap := make(map[string]xml.Attr)
			for _, val := range attr {
				attrMap[val.Name.Local] = val
			}
			node := XmlNode{
				Name:     name,
				Attrs:    attrMap,
				Elements: make([]element, 0),
			}
			for _, val := range attr {
				if val.Name.Local == "id" {
					node.Id = val.Value
				}
			}
			st.Push(node)

		case xml.EndElement: //tag end
			if st.Len() > 0 {
				//cur XmlNode
				n := st.Pop().(XmlNode)
				if st.Len() > 0 { //if the root XmlNode then append to element
					e := element{
						ElementType: eleTpNode,
						Val:         n,
					}

					pn := st.Pop().(XmlNode)
					els := pn.Elements
					els = append(els, e)
					pn.Elements = els
					st.Push(pn)
				} else { //else root = n
					root = n
				}
			}
		case xml.CharData: //tag content
			if st.Len() > 0 {
				n := st.Pop().(XmlNode)

				bytes := xml.CharData(t)
				content := strings.TrimSpace(string(bytes))
				if content != "" {
					e := element{
						ElementType: eleTpText,
						Val:         content,
					}
					els := n.Elements
					els = append(els, e)
					n.Elements = els
				}

				st.Push(n)
			}

		case xml.Comment:
		case xml.ProcInst:
		case xml.Directive:
		default:
		}
	}

	if st.Len() != 0 {
		panic("Parse xml error, there is tag no close, please check your xml config!")
	}

	return &root
}
