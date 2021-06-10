package dz3Lesson

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
func Sort(dir string) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, dir, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}
	ast.SortImports(fset, file)
}
