package evaluator

import (
	"GoRilla/ast"
	"GoRilla/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}