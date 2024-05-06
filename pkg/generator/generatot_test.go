/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-04-28 10:46:05
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-05-06 11:16:54
 * @FilePath: /dao-generator/pkg/generator/generatot_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package generator_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/chein-huang/dao-generator/pkg/generator"
	"github.com/chein-huang/dao-generator/pkg/model"
	"github.com/stretchr/testify/require"
)

var workspaceDir string

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get current file path")
	}

	workspaceDir = filepath.Join(filepath.Dir(filename), "../..")
}

func TestGenerateWithData(t *testing.T) {
	ass := require.New(t)

	inputDir := filepath.Join(workspaceDir, "pkg/generator/test_file")
	outputDir := filepath.Join(workspaceDir, "pkg/generator/output")
	err := generator.Generate(inputDir, outputDir, model.ORMTypeGorm)
	ass.NoError(err)
}

func TestGenerationMetaData(t *testing.T) {
	ass := require.New(t)

	data := model.GenerationMetaData{
		Tables: []*model.GenerationTable{
			{
				Name:          "approval_info",
				NameWithSpace: "approval info",
				Imports: []string{
					"gitlab.bj.sensetime.com/elementary/graviton/graviton-data-compliance-service/pkg/model",
				},
				StructName:            "ApprovalInfo",
				StructNameSmallCamel:  "approvalInfo",
				ModelPackage:          "model",
				StructNameWithPreload: "ApprovalInfoWithPreload",
				Fields: []*model.GenerationField{
					{
						Name:      "ID",
						NameSnake: "id",
						Type:      "string",
						Order:     false,
						Range:     false,
					},
					{
						Name:      "State",
						NameSnake: "state",
						Type:      "model.ApprovalState",
						Order:     false,
						Range:     false,
					},
					{
						Name:      "CreatedAt",
						NameSnake: "created_at",
						Type:      "time.Time",
						Order:     true,
						Range:     false,
					},
					{
						Name:      "UpdatedAt",
						NameSnake: "updated_at",
						Type:      "time.Time",
						Order:     true,
						Range:     true,
					},
				},
				Preloads: []*model.GenerationPreload{
					{
						Name:    "AuthInfo",
						OrderBy: "",
					},
					{
						Name:    "ApprovalRecords",
						OrderBy: "created_at DESC",
					},
					{
						Name:    "SampleFiles",
						OrderBy: "",
					},
				},
			},
		},
	}

	dir := filepath.Join(workspaceDir, "pkg/generator/output")
	err := generator.GenByTemplate(&data, dir, true)
	ass.NoError(err)
}

func TestParseGenFlags(t *testing.T) {
	texts := []string{
		`// gen:"approval_info"`,
		`// gen :"approval_info"`,
		`//  gen:"approval_info"`,
		`// gen:" approval_info"`,
		`// gen:"approval_info "`,
	}
	ass := require.New(t)

	for _, text := range texts {
		table := &model.GenerationTable{}
		generator.ParseGenFlags(text, table)
		ass.Equal("approval_info", table.Name, "text: %q", text)
	}
}

func TestGetPackagePath(t *testing.T) {
	ass := require.New(t)
	path, err := generator.GetPackagePath(filepath.Join(workspaceDir, "pkg/generator/test_file"))
	ass.NoError(err)
	ass.Equal("github.com/chein-huang/dao-generator/pkg/generator/test_file", path)
}
