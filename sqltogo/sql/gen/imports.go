package gen

import (
	"work/sqltogo/sql/template"
	"work/sqltogo/sql/util"
	"work/sqltogo/sql/util/pathx"
)

func genImports(table Table, timeImport bool) (string, error) {
	text, err := pathx.LoadTemplate(category, importsTemplateFile, template.Imports)
	if err != nil {
		return "", err
	}

	output, err := util.With("import").Parse(text).Execute(map[string]interface{}{
		"time":       timeImport,
		"containsPQ": table.ContainsPQ,
		"data":       table,
	})
	if err != nil {
		return "", err
	}
	return output.String(), nil
}
