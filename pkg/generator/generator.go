package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/chein-huang/dao-generator/pkg/model"
	"github.com/chein-huang/errorc"
	"github.com/pkg/errors"
)

func Generate(input, output string, ormType model.ORMType) error {
	inputFile, err := os.Open(input)
	if err != nil {
		return errorc.AddField(err, "input", input)
	}
	defer inputFile.Close()

	ext := filepath.Ext(input)
	inputName := strings.TrimSuffix(filepath.Base(input), ext)
	outputFileName := filepath.Join(output, inputName+"_gorm.go")
	_, err = os.Stat(outputFileName)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return errorc.AddField(err, "output", outputFileName)
	} else if err == nil {
		return errorc.Newf("output exists: %v", outputFileName)
	}

	outputFile, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return errorc.AddField(err, "output", outputFileName)
	}
	defer outputFile.Close()

	return GenerateWithData(inputFile, outputFile, ormType)
}

func GenerateWithData(input io.Reader, output io.Writer, ormType model.ORMType) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", input, parser.ParseComments)
	if err != nil {
		fmt.Printf("err = %s", err)
	}

	// for _, d := range f.Decls {
	// 	// TODO: 读取结构体定义
	// 	// TODO: 遍历字段
	// 	// TODO:
	// 	// switch dec := d.(type) {
	// 	// case ast.StructType:
	// 	// }
	// }
	ast.Print(fset, f)
	return nil
}

func getDeclsFromFile(file *ast.File) (objects []ast.GenDecl) {
	for _, d := range file.Decls {
		switch dec := d.(type) {
		case *ast.GenDecl:
			switch dec.Tok {
			case token.TYPE:

			}
		}
	}
	return nil
}
