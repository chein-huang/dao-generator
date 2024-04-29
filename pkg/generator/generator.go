package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"

	"github.com/chein-huang/dao-generator/pkg/model"
	"github.com/chein-huang/errorc"
)

func Generate(input, output string, ormType model.ORMType) error {
	inputFile, err := os.Open(input)
	if err != nil {
		return errorc.AddField(err, "input", input)
	}
	defer inputFile.Close()

	return GenerateWithData(inputFile, output, ormType)
}

func GenerateWithData(input io.Reader, output string, ormType model.ORMType) error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", input, parser.ParseComments)
	if err != nil {
		fmt.Printf("err = %s", err)
	}

	data, err := GetMetaDataFromFile(f)
	if err != nil {
		return err
	}

	err = GenByTemplate(data, output, true)
	if err != nil {
		return err
	}
	ast.Print(fset, f)
	return nil
}
