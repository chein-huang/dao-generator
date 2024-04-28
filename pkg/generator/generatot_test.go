/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-04-28 10:46:05
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-04-28 18:39:14
 * @FilePath: /dao-generator/pkg/generator/generatot_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package generator_test

import (
	"bytes"
	"embed"
	"testing"

	"github.com/chein-huang/dao-generator/pkg/generator"
	"github.com/chein-huang/dao-generator/pkg/model"
	"github.com/stretchr/testify/require"
)

//go:embed test_file/approval.go
var approval embed.FS

func TestGenerate(t *testing.T) {
	ass := require.New(t)
	f, err := approval.Open("test_file/approval.go")
	ass.NoError(err)
	defer f.Close()

	buf := bytes.Buffer{}
	err = generator.GenerateWithData(f, &buf, model.ORMTypeGorm)
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

	dir := "/Users/huangcheng/go/src/github.com/chein.huang/dao-generator/pkg/generator/output"
	err := generator.GenByTemplate(&data, dir, true)
	ass.NoError(err)
}