package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

// すべてのファイルのOneOf～と*OneOf～をinterface{}に書き換えるスクリプト
func main() {
	fileList, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range fileList {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		if !strings.HasSuffix(name, ".go") {
			continue
		}

		fmt.Println("Processing: " + name)
		rewriteInvalidOneOfTypesToEmptyInterface("./" + name)
	}
}

func rewriteInvalidOneOfTypesToEmptyInterface(filepath string) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filepath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	newF := astutil.Apply(f, func(c *astutil.Cursor) bool {
		node := c.Node()

		if starExprNode, ok := node.(*ast.StarExpr); ok {
			if isOneOfIdent(starExprNode.X) {
				c.Replace(getEmptyInterfaceType())
			}
		} else if isOneOfIdent(node) {
			c.Replace(getEmptyInterfaceType())
		}

		return true
	}, nil)

	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	if err := format.Node(file, fs, newF); err != nil {
		panic(err)
	}
}

func isOneOfIdent(node ast.Node) bool {
	identNode, ok := node.(*ast.Ident)
	if !ok {
		return false
	}

	return strings.HasPrefix(identNode.Name, "OneOf")
}

func getEmptyInterfaceType() ast.Node {
	return &ast.InterfaceType{
		Methods:    &ast.FieldList{},
		Incomplete: false,
	}
}
