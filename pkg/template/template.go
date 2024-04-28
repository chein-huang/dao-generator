package template

import (
	_ "embed"
	"text/template"
)

//go:embed tmpl/crud_gorm.tmpl
var crudGormTmplStr string

//go:embed tmpl/crud_opts_gorm.tmpl
var crudOptsGormTmplStr string

//go:embed tmpl/dao.tmpl
var daoTmplStr string

//go:embed tmpl/errors.tmpl
var errorsTmplStr string

//go:embed tmpl/transaction_gorm.tmpl
var transactionGormTmplStr string

var (
	CrudGormTmpl        = template.New("crudGormTmpl")
	CrudOptsGormTmpl    = template.New("crudOptsGormTmpl")
	DaoTmpl             = template.New("daoTmpl")
	ErrorsTmpl          = template.New("errorsTmpl")
	TransactionGormTmpl = template.New("transactionGormTmpl")
)

func init() {
	_, err := CrudGormTmpl.Parse(crudGormTmplStr)
	if err != nil {
		panic(err)
	}

	_, err = CrudOptsGormTmpl.Parse(crudOptsGormTmplStr)
	if err != nil {
		panic(err)
	}

	_, err = DaoTmpl.Parse(daoTmplStr)
	if err != nil {
		panic(err)
	}

	_, err = ErrorsTmpl.Parse(errorsTmplStr)
	if err != nil {
		panic(err)
	}

	_, err = TransactionGormTmpl.Parse(transactionGormTmplStr)
	if err != nil {
		panic(err)
	}
}
