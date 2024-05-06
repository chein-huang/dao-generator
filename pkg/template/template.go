package template

import (
	_ "embed"
	"text/template"
)

//go:embed tmpl/crud_gorm.tmpl
var crudGormTmplStr string

//go:embed tmpl/dao_gorm.tmpl
var daoGormTmplStr string

//go:embed tmpl/errors_gorm.tmpl
var errorsGormTmplStr string

//go:embed tmpl/transaction_gorm.tmpl
var transactionGormTmplStr string

var (
	CrudGormTmpl        = template.New("crudGormTmpl")
	DaoGormTmpl         = template.New("daoGormTmpl")
	ErrorsTmpl          = template.New("errorsTmpl")
	TransactionGormTmpl = template.New("transactionGormTmpl")
)

func init() {
	_, err := CrudGormTmpl.Parse(crudGormTmplStr)
	if err != nil {
		panic(err)
	}

	_, err = DaoGormTmpl.Parse(daoGormTmplStr)
	if err != nil {
		panic(err)
	}

	_, err = ErrorsTmpl.Parse(errorsGormTmplStr)
	if err != nil {
		panic(err)
	}

	_, err = TransactionGormTmpl.Parse(transactionGormTmplStr)
	if err != nil {
		panic(err)
	}
}
