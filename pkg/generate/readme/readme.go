package readme

import (
	"fmt"
	"strings"

	"github.com/danspts/helmdocs/pkg/parse"
	"github.com/danspts/helmdocs/pkg/types"
)

func GenerateReadme(schema types.Schema) string {
	var sb strings.Builder

	// Initialize the table headers
	sb.WriteString("| Name | Title | Type | Required | Default |\n")
	sb.WriteString("|------|-------|------|----------|---------|\n")

	tree := parse.ConvertTableToFieldTree(types.Property{Schema: schema}, "")
	flatList := tree.Flatten()
	for _, field := range flatList {
		typeDetails := field.TypeDetails
		if len(typeDetails) > 0 {
			typeDetails = " " + typeDetails
		}
		if field.ParentFullName == "" {
			sb.WriteString(fmt.Sprintf("| `%s` | %s | **%s**%s | %v | %s |\n", field.Name, field.Title, field.Type, field.TypeDetails, field.Required, field.DefaultValue))

		} else {
			sb.WriteString(fmt.Sprintf("| `%s.%s` | %s | **%s**%s | %v | %s |\n", field.ParentFullName, field.Name, field.Title, field.Type, field.TypeDetails, field.Required, field.DefaultValue))
		}
	}
	return sb.String()

}
