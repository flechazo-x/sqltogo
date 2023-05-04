package template

import (
	"fmt"
	"work/sqltogo/sql/util"
)

// ModelCustom defines a template for extension
const ModelCustom = `package {{.pkg}}
{{if .withCache}}
import (
	"errors"
	"fmt"
	"strings"
)
`

// ModelGen defines a template for model
var ModelGen = fmt.Sprintf(`%s

package {{.pkg}}
{{.imports}}
{{.vars}}
{{.types}}
{{.delete}}
{{.find}}
{{.findBatch}}
{{.findRow}}
{{.insert}}
{{.update}}
`, util.DoNotEditHead)
