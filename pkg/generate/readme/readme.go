package readme

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/danspts/helmdocs/pkg/parse"
	"github.com/danspts/helmdocs/pkg/types"
)

func GenerateReadme(filename string) {
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

	// Traverse the schema and generate the markdown content
	markdown := generateMarkdownTableWithHeader(schema)

	// Write the markdown content to a README file
	if err := os.WriteFile("README.md", []byte(markdown), 0644); err != nil {
		fmt.Println("Error writing README.md:", err)
		return
	}

	fmt.Println("README.md generated successfully.")
}

func generateMarkdownTableWithHeader(schema types.Schema) string {
	var sb strings.Builder

	// Initialize the table headers
	sb.WriteString("| Name | Title | Type | Required | Default |\n")
	sb.WriteString("|------|-------|------|----------|---------|\n")

	tree := parse.ConvertTableToFieldTree(types.Property{Schema: schema}, "")
	flatList := tree.Flatten()
	for _, field := range flatList {
		if field.ParentFullName == "" {
			sb.WriteString(fmt.Sprintf("| `%s` | %s | %s | %v | %s |\n", field.Name, field.Title, field.TypeDetails, field.Required, field.DefaultValue))

		} else {
			sb.WriteString(fmt.Sprintf("| `%s.%s` | %s | %s | %v | %s |\n", field.ParentFullName, field.Name, field.Title, field.TypeDetails, field.Required, field.DefaultValue))
		}
	}
	return sb.String()

}
