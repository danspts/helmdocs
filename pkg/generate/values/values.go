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
		comment := ""
		if field.TypeDetails != "" {
			comment = " # " + field.TypeDetails
		}
		if len(field.Children) > 0 {
			sb.WriteString(fmt.Sprintf("%s%s:%s\n", indent, field.Name, comment))
		} else if len(field.DefaultValue) > 0 {
			sb.WriteString(fmt.Sprintf("%s%s%s: %s%s\n", indent, commitDefaultToken, field.Name, field.DefaultValue, comment))
		} else {
			sb.WriteString(fmt.Sprintf("%s%s: <%s>%s\n", indent, field.Name, field.Type, comment))
		}

	}
	return sb.String()
}
