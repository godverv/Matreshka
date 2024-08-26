package patterns

import (
	_ "embed"
	"text/template"
)

//go:embed data_sources.go.pattern
var dataSourcePattern string

var TmplDataSource *template.Template

func init() {
	TmplDataSource = template.New("data_sources")

	_, err := TmplDataSource.Parse(dataSourcePattern)
	if err != nil {
		panic(err)
	}
}
