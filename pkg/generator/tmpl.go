/*
 * @Author: huangcheng1 huangcheng1@sensetime.com
 * @Date: 2024-04-28 18:39:19
 * @LastEditors: huangcheng1 huangcheng1@sensetime.com
 * @LastEditTime: 2024-04-30 17:28:32
 * @FilePath: /dao-generator/pkg/generator/tmpl.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package generator

import (
	"os"
	"path/filepath"

	"github.com/chein-huang/dao-generator/pkg/model"
	"github.com/chein-huang/dao-generator/pkg/template"
	"github.com/chein-huang/dao-generator/pkg/utils"
	"github.com/chein-huang/errorc"
	"github.com/pkg/errors"
)

func GenByTemplate(data *model.GenerationMetaData, output string, overwrite bool) error {
	stat, err := os.Stat(output)
	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(output, os.ModePerm)
		if err != nil {
			return errorc.AddField(err, "path", output)
		}
	} else if err != nil {
		return errorc.AddField(err, "path", output)
	} else if !stat.IsDir() {
		return errorc.Newf("%s is not dir", output)
	}

	if err = genGlobalFile(data, output, overwrite); err != nil {
		return err
	}

	for _, table := range data.Tables {
		if err = genTableFile(table, output, overwrite); err != nil {
			return err
		}
	}
	return nil
}

func genGlobalFile(data *model.GenerationMetaData, output string, overwrite bool) error {
	daoFileName := filepath.Join(output, "dao_gorm.go")

	daoF, err := utils.OpenFile(daoFileName, os.O_CREATE|os.O_WRONLY, os.ModePerm, overwrite)
	if err != nil {
		return errorc.AddField(err, "file name", daoFileName)
	}
	defer daoF.Close()

	err = template.DaoGormTmpl.Execute(daoF, data)
	if err != nil {
		return errorc.AddField(err, "data", data)
	}

	transactionFileName := filepath.Join(output, "transaction_gorm.go")
	transactionF, err := utils.OpenFile(transactionFileName, os.O_CREATE|os.O_WRONLY, os.ModePerm, overwrite)
	if err != nil {
		return errorc.AddField(err, "file name", transactionFileName)
	}
	defer transactionF.Close()

	err = template.TransactionGormTmpl.Execute(transactionF, data)
	if err != nil {
		return errorc.AddField(err, "data", data)
	}

	errorsFileName := filepath.Join(output, "errors_gorm.go")
	errorsF, err := utils.OpenFile(errorsFileName, os.O_CREATE|os.O_WRONLY, os.ModePerm, overwrite)
	if err != nil {
		return errorc.AddField(err, "file name", errorsFileName)
	}
	defer errorsF.Close()

	err = template.ErrorsTmpl.Execute(errorsF, data)
	if err != nil {
		return errorc.AddField(err, "data", data)
	}
	return nil
}

func genTableFile(table *model.GenerationTable, output string, overwrite bool) error {
	crudFileName := filepath.Join(output, table.Name+"_crud_gorm.go")

	crudF, err := utils.OpenFile(crudFileName, os.O_CREATE|os.O_WRONLY, os.ModePerm, overwrite)
	if err != nil {
		return errorc.AddField(err, "file name", crudFileName)
	}
	defer crudF.Close()

	err = template.CrudGormTmpl.Execute(crudF, table)
	if err != nil {
		return errorc.AddField(err, "table", table)
	}
	return nil
}
