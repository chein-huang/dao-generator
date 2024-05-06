package model

const (
	OpEq = "eq"
	OpIn = "in"
)

type GenerationMetaData struct {
	Tables []*GenerationTable
}

type GenerationTable struct {
	Name                  string
	NameWithSpace         string
	IsPreload             bool
	StructName            string
	StructNameSmallCamel  string
	Imports               []string
	ModelPackage          string
	StructNameWithPreload string
	Fields                []*GenerationField
	Preloads              []*GenerationPreload
}

type GenerationField struct {
	Name      string
	NameSnake string
	Type      string
	Order     bool
	Range     bool
	In        bool
}

type GenerationPreload struct {
	Name    string
	OrderBy string
}
