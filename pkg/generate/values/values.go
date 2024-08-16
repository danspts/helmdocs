package values

import (
	"fmt"
	"strings"

	"github.com/danspts/helmdocs/pkg/parse"
	"github.com/danspts/helmdocs/pkg/types"
)

func GenerateValues(schema types.Schema, omitDft bool) string {
	var sb strings.Builder
	commitDefaultToken := ""
	if omitDft {
		commitDefaultToken = "# "
	}

	tree := parse.ConvertTableToFieldTree(types.Property{Schema: schema}, "")
	flatList := tree.Flatten()
	for i, field := range flatList {
		if field.Depth() == 0 && i != 0 {
			sb.WriteString("\n")
		}
		indent := strings.Repeat("  ", field.Depth())
		if len(field.Children) > 0 {
			sb.WriteString(fmt.Sprintf("%s%s:\n", indent, field.Name))
		} else if len(field.DefaultValue) > 0 {
			sb.WriteString(fmt.Sprintf("%s%s%s: %s\n", indent, commitDefaultToken, field.Name, field.DefaultValue))
		} else {
			sb.WriteString(fmt.Sprintf("%s%s: <%s>\n", indent, field.Name, field.Type))
		}

	}
	return sb.String()
}
