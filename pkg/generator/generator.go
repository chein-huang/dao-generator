package generator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/chein-huang/dao-generator/pkg/model"
	"github.com/chein-huang/errorc"
)

func Generate(input, output string, ormType model.ORMType) error {
	entries, err := os.ReadDir(input)
	if err != nil {
		return errorc.Wrap(err)
	}

	data := model.GenerationMetaData{}
	for _, entry := range entries {
		if !entry.IsDir() {
			err = func() error {
				fileName := filepath.Join(input, entry.Name())
				inputFile, err := os.Open(fileName)
				if err != nil {
					return errorc.AddField(err, "input", input)
				}
				defer inputFile.Close()

				fset := token.NewFileSet()
				f, err := parser.ParseFile(fset, "", inputFile, parser.ParseComments)
				if err != nil {
					fmt.Printf("err = %s", err)
				}

				tables, err := GetTablesFromFile(fileName, f)
				if err != nil {
					return err
				}
				data.Tables = append(data.Tables, tables...)

				ast.Print(fset, f)
				return nil
			}()
			if err != nil {
				return err
			}
		}
	}

	err = GenByTemplate(&data, output, true)
	if err != nil {
		return err
	}
	return nil
}
