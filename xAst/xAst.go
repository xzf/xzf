package xAst

import (
	"go/token"
	"go/parser"
	"go/ast"
	"strings"
)

/*
	get all const in source code
	all value will be convert to string
*/

type StringConst struct {
	Name  string
	Value string
}

func GetConstList(sourceCode string) (result []StringConst, err error) {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, "", sourceCode, 0)
	if err != nil {
		return
	}
	ast.Inspect(astFile, func(astNode ast.Node) bool {
		switch one := astNode.(type) {
		case *ast.GenDecl:
			if one.Tok.String() != "const" {
				return true
			}
			var oneGroup []StringConst
			for _, one := range one.Specs {
				oneX := one.(*ast.ValueSpec)
				for index, oneName := range oneX.Names {
					val := oneX.Values[index].(*ast.BasicLit)
					oneGroup = append(oneGroup, StringConst{
						Name:  oneName.Name,
						Value: strings.Replace(val.Value, `"`, ``, -1),
					})
				}
			}
			result = append(result, oneGroup...)
		}
		return true
	})
	return
}