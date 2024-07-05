package jsonutils

import (
	"fmt"
	"strings"
)

type jsonElement uint8

const (
	jsonElementObject jsonElement = iota
	jsonElementArray
)

type jsonPathNode struct {
	key  string
	elem jsonElement
}

func getKey(nodes []jsonPathNode) string {
	var sb strings.Builder
	sb.WriteString("$")
	for _, node := range nodes {
		if node.elem == jsonElementArray {
			sb.WriteString(fmt.Sprintf("[%s]", node.key))
		} else {
			if sb.Len() > 0 {
				sb.WriteString(".")
			}
			sb.WriteString(node.key)
		}
	}
	return sb.String()
}
