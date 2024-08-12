package values

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/danspts/helmdocs/pkg/parse"
	"github.com/danspts/helmdocs/pkg/types"
)

func GenerateValues(filename string, omitDft bool) {
	// Read the JSON schema file
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the JSON schema into a struct
	var schema types.Schema
	if err := json.Unmarshal(file, &schema); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	markdown := generateMarkdownTableWithHeader(schema, omitDft)

	// Write the markdown content to a values.yaml file
	if err := os.WriteFile("values.yaml", []byte(markdown), 0644); err != nil {
		fmt.Println("Error writing values.yaml", err)
		return
	}

	fmt.Println("values.yaml generated successfully.")
}

func generateMarkdownTableWithHeader(schema types.Schema, omitDft bool) string {
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
