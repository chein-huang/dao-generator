/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-04-29 11:25:40
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-05-06 14:11:28
 * @FilePath: /dao-generator/pkg/generator/gen_data.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package generator

import (
	"bufio"
	"go/ast"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/chein-huang/dao-generator/pkg/model"
	"github.com/chein-huang/errorc"
	"github.com/iancoleman/strcase"
)

var (
	genReg, _     = regexp.Compile(`\/\/ *gen *: *"(.*)"`)
	packageReg, _ = regexp.Compile(`package *(.*)`)
	gormReg, _    = regexp.Compile(`.* *gorm *: *"(.*)"`)
)

func GetTablesFromFile(path string, file *ast.File) (tables []*model.GenerationTable, err error) {
	packagePath, err := GetPackagePath(filepath.Dir(path))
	if err != nil {
		return nil, err
	}

	packageName, err := GetModuleName(path)
	if err != nil {
		return nil, err
	}

	tables = []*model.GenerationTable{}
	for _, d := range file.Decls {
		switch dec := d.(type) {
		case *ast.GenDecl:
			switch dec.Tok {
			case token.TYPE:
				if isGenDecl(dec) {
					table := &model.GenerationTable{
						ModelPackage: packageName,
						Imports:      []string{packagePath},
					}
					err := processTableFromDecl(dec, table)
					if err != nil {
						return nil, err
					}

					tables = append(tables, table)
				}
			}
		}
	}
	return tables, nil
}

func processTableFromDecl(decl *ast.GenDecl, table *model.GenerationTable) error {
	for _, doc := range decl.Doc.List {
		if strings.Contains(doc.Text, "gen:") {
			ParseGenFlags(doc.Text, table)
			if table.Name == "" {
				return errorc.Newf("table name empty")
			}
		}
	}

	ParseFields(decl.Specs, table)

	return nil
}

func ParseGenFlags(text string, table *model.GenerationTable) {
	flagStr := genReg.FindStringSubmatch(text)
	if len(flagStr) > 1 {
		items := strings.Split(flagStr[1], ",")
		for _, item := range items {
			splits := strings.Split(item, ":")
			if len(splits) == 1 {
				table.StructName = strings.TrimSpace(splits[0])
				if table.Name == "" {
					table.Name = strcase.ToSnake(table.StructName)
				}
				if table.StructNameSmallCamel == "" {
					table.StructNameSmallCamel = strcase.ToLowerCamel(table.StructName)
				}
				if table.NameWithSpace == "" {
					table.NameWithSpace = strings.ReplaceAll(strcase.ToSnake(table.StructName), "_", " ")
				}
			} else if len(splits) > 1 {
				key := strings.ToLower(strings.TrimSpace(splits[0]))
				value := strings.TrimSpace(splits[1])
				switch key {
				case "packagealicename":
					table.ModelPackage = value
				case "namesnake":
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
				newTable := &model.GenerationTable{
					ModelPackage: table.ModelPackage,
				}
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

				tags := &FieldTags{}
				if field.Tag != nil {
					parseFieldTags(field.Tag.Value, tags)
				}

				if tags.Column == "" {
					tags.Column = strcase.ToSnake(field.Names[0].Name)
				}

				f := &model.GenerationField{
					Name:      field.Names[0].Name,
					NameSnake: tags.Column,
					Type:      getTypeName(field.Type),
					Order:     flags.Order,
					Range:     flags.Range,
					In:        flags.In,
				}
				if !isBasicTypeStr(f.Type) && !strings.Contains(f.Type, ".") {
					f.Type = table.ModelPackage + "." + f.Type
				}

				table.Fields = append(table.Fields, f)
			}
		}
	case *ast.StarExpr:
		parseFields(t.X, table)
	case *ast.Ident:
		parseFields(t.Obj.Decl.(*ast.TypeSpec).Type, table)
	}
}

type FieldFlags struct {
	OrderBy string
	Order   bool
	Range   bool
	In      bool
}

func parseFieldFlags(text string, flags *FieldFlags) {
	flagStr := genReg.FindStringSubmatch(text)
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
						case "in":
							flags.In = true
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

func GetModuleName(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", errorc.Wrap(err)
	}
	defer f.Close()

	packageName := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "package") {
			res := packageReg.FindStringSubmatch(text)
			if len(res) > 1 {
				packageName = res[1]
				break
			}
		}
	}
	return packageName, nil
}

func GetPackagePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", errorc.Wrap(err)
	}

	modPath, err := findGoModPath(absPath)
	if err != nil {
		return "", err
	}

	if modPath == "" {
		return "", nil
	}

	relPath, err := filepath.Rel(filepath.Dir(modPath), absPath)
	if err != nil {
		return "", errorc.Wrap(err)
	}

	f, err := os.Open(modPath)
	if err != nil {
		return "", errorc.Wrap(err)
	}
	defer f.Close()

	module := ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "module") {
			res := moduleReg.FindStringSubmatch(text)
			if len(res) > 1 {
				module = res[1]
				break
			}
		}
	}

	return filepath.Join(module, relPath), err
}

var moduleReg, _ = regexp.Compile(`module *(.*)`)

func findGoModPath(absPath string) (string, error) {
	p := absPath

	for p != "." && p != "/" {
		entries, err := os.ReadDir(p)
		if err != nil {
			return "", errorc.Wrap(err)
		}

		for _, entry := range entries {
			if !entry.IsDir() && entry.Name() == "go.mod" {
				return filepath.Join(p, entry.Name()), nil
			}
		}
		p = filepath.Dir(p)
	}
	return "", nil
}

func isBasicTypeStr(typName string) bool {
	for strings.HasPrefix(typName, "*") {
		typName = strings.TrimPrefix(typName, "*")
	}
	switch typName {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"complex64", "complex128",
		"string",
		"bool":
		return true
	default:
		if strings.HasPrefix(typName, "map") {
			return true
		} else if strings.HasPrefix(typName, "[]") {
			return isBasicTypeStr(typName)
		}
		return false
	}
}

type FieldTags struct {
	Column string
}

func parseFieldTags(tag string, tags *FieldTags) {
	tagStr := gormReg.FindStringSubmatch(tag)
	if len(tagStr) > 1 {
		items := strings.Split(tagStr[1], ";")
		for _, item := range items {
			splits := strings.Split(item, ":")
			if len(splits) > 1 {
				value := strings.TrimSpace(splits[1])
				switch strings.ToLower(splits[0]) {
				case "column":
					tags.Column = value
				}
			}
		}
	}
}
