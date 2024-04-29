/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-04-29 11:25:40
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-04-29 18:42:13
 * @FilePath: /dao-generator/pkg/generator/gen_data.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package generator

import (
	"go/ast"
	"go/token"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/chein-huang/dao-generator/pkg/model"
	"github.com/chein-huang/errorc"
	"github.com/iancoleman/strcase"
)

func GetMetaDataFromFile(file *ast.File) (data *model.GenerationMetaData, err error) {
	data = &model.GenerationMetaData{}
	for _, d := range file.Decls {
		switch dec := d.(type) {
		case *ast.GenDecl:
			switch dec.Tok {
			case token.TYPE:
				if isGenDecl(dec) {
					table, err := getTableFromDecl(dec)
					if err != nil {
						return nil, err
					}
					data.Tables = append(data.Tables, table)
				}
			}
		}
	}
	return data, nil
}

func getTableFromDecl(decl *ast.GenDecl) (*model.GenerationTable, error) {
	table := &model.GenerationTable{}
	for _, doc := range decl.Doc.List {
		if strings.Contains(doc.Text, "gen:") {
			ParseGenFlags(doc.Text, table)
			if table.Name == "" {
				return nil, errorc.Newf("table name empty")
			}
		}
	}

	ParseFields(decl.Specs, table)

	return table, nil
}

var reg, _ = regexp.Compile(`\/\/ *gen *: *"(.*)"`)

func ParseGenFlags(text string, table *model.GenerationTable) {
	flagStr := reg.FindStringSubmatch(text)
	if len(flagStr) > 1 {
		items := strings.Split(flagStr[1], ",")
		for _, item := range items {
			splits := strings.Split(item, ":")
			if len(splits) == 1 {
				table.StructName = strings.TrimSpace(splits[0])
				if table.Name == "" {
					strcase.ToSnake(table.StructName)
				}
			} else {
				key := strings.ToLower(strings.TrimSpace(splits[0]))
				value := strings.TrimSpace(splits[1])
				switch key {
				case "package":
					table.Imports = append(table.Imports, value)
					if table.ModelPackage == "" {
						table.ModelPackage = filepath.Base(value)
					}
				case "packagealicename":
					table.ModelPackage = value
				case "tablename":
					table.Name = value
				case "flags":
					flags := strings.Split(strings.ToLower(value), ";")
					for _, flag := range flags {
						switch flag {
						case "ispreload":
							table.IsPreload = true
						}
					}
				default:
				}
			}
		}
	}
}

func ParseFields(specs []ast.Spec, table *model.GenerationTable) {
	if len(specs) != 1 {
		return
	}

	spec := specs[0]
	switch s := spec.(type) {
	case *ast.TypeSpec:
		if table.IsPreload {
			table.StructNameWithPreload = s.Name.Name
		} else {
			table.StructNameWithPreload = table.StructName
		}
		parseFields(s.Type, table)
	}
}

func parseFields(typ ast.Expr, table *model.GenerationTable) {
	switch t := typ.(type) {
	case *ast.StructType:
		for _, field := range t.Fields.List {
			if len(field.Names) == 0 {
				newTable := &model.GenerationTable{}
				parseFields(field.Type, newTable)
				table.Fields = append(table.Fields, newTable.Fields...)
			} else if table.IsPreload {
				flags := &FieldFlags{}
				if field.Doc != nil {
					for _, text := range field.Doc.List {
						parseFieldFlags(text.Text, flags)
					}
				}
				table.Preloads = append(table.Preloads, &model.GenerationPreload{
					Name:    field.Names[0].Name,
					OrderBy: flags.OrderBy,
				})
			} else {
				flags := &FieldFlags{}
				if field.Doc != nil {
					for _, text := range field.Doc.List {
						parseFieldFlags(text.Text, flags)
					}
				}

				table.Fields = append(table.Fields, &model.GenerationField{
					Name:      field.Names[0].Name,
					NameSnake: strcase.ToSnake(field.Names[0].Name),
					Type:      getTypeName(field.Type),
					Order:     flags.Order,
					Range:     flags.Range,
				})
			}
		}
	case *ast.StarExpr:
		parseFields(t.X, table)
	case *ast.Ident:
		parseFields(t.Obj.Decl.(*ast.TypeSpec).Type, table)
	}
	if structType, ok := typ.(*ast.StructType); !ok {
		return
	} else {
		for _, field := range structType.Fields.List {
			if len(field.Names) == 0 {
				parseFields(field.Type, table)
			} else {

			}
		}
	}
}

type FieldFlags struct {
	OrderBy string
	Order   bool
	Range   bool
}

func parseFieldFlags(text string, flags *FieldFlags) {
	flagStr := reg.FindStringSubmatch(text)
	if len(flagStr) > 1 {
		items := strings.Split(flagStr[1], ",")
		for _, item := range items {
			splits := strings.Split(item, ":")
			if len(splits) > 1 {
				key := strings.ToLower(strings.TrimSpace(splits[0]))
				value := strings.TrimSpace(splits[1])
				switch key {
				case "orderby":
					flags.OrderBy = value
				case "flags":
					fs := strings.Split(strings.ToLower(value), ";")
					for _, f := range fs {
						switch f {
						case "order":
							flags.Order = true
						case "range":
							flags.Range = true
						}
					}
				default:
				}
			}
		}
	}
}

func isGenDecl(dec *ast.GenDecl) bool {
	if dec.Doc == nil {
		return false
	}

	exists := false
	for _, doc := range dec.Doc.List {
		if strings.Contains(doc.Text, "gen:") {
			exists = true
		}
	}
	return exists
}

func getTypeName(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.StarExpr:
		return getTypeName(t.X)
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return getTypeName(t.X) + "." + t.Sel.Name
	default:
		return ""
	}
}
