package readme

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Schema struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required,omitempty"`
}

type Property struct {
	Type        string              `json:"type"`
	Title       string              `json:"title,omitempty"`
	Description string              `json:"description,omitempty"`
	Form        bool                `json:"form,omitempty"`
	Render      string              `json:"render,omitempty"`
	Enum        []string            `json:"enum,omitempty"`
	Pattern     string              `json:"pattern,omitempty"`
	SliderMin   *int                 `json:"sliderMin,omitempty"`
	SliderMax   *int                 `json:"sliderMax,omitempty"`
	SliderUnit  string              `json:"sliderUnit,omitempty"`
	Default     interface{}         `json:"default,omitempty"`
	Properties  map[string]Property `json:"properties,omitempty"`
	Hidden      Hidden              `json:"hidden,omitempty"`
	Required    []string            `json:"required,omitempty"`
}

type Hidden struct {
	Condition bool   `json:"condition"`
	Value     string `json:"value"`
}

func (h *Hidden) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		// Data is an object
		type Alias Hidden
		aux := &struct{ *Alias }{Alias: (*Alias)(h)}
		return json.Unmarshal(data, &aux)
	}
	// Data is a string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	h.Condition = false
	h.Value = str
	return nil
}

func GenerateReadme(filename string) {
	// Read the JSON schema file
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse the JSON schema into a struct
	var schema Schema
	if err := json.Unmarshal(file, &schema); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Traverse the schema and generate the markdown content
	markdown := generateMarkdownTableWithHeader(schema, "", make(map[string]bool))

	// Write the markdown content to a README file
	if err := os.WriteFile("README.md", []byte(markdown), 0644); err != nil {
		fmt.Println("Error writing README.md:", err)
		return
	}

	fmt.Println("README.md generated successfully.")
}

func generateMarkdownTableWithHeader(schema Schema, prefix string, globalRequired map[string]bool) string {
	var sb strings.Builder

	// Initialize the table headers
	sb.WriteString("| Name | Title | Type | Required | Default |\n")
	sb.WriteString("|------|-------|------|----------|---------|\n")

	return generateMarkdownTable(schema, "", make(map[string]bool), &sb).String()

}

func extractDefaultFromHidden(hidden Hidden) string {
	// Check if the hidden attribute has a condition and a value
	if !hidden.Condition && hidden.Value != "" {
		// Return the default value based on the hidden attribute's value
		return hidden.Value
	}
	return ""
}

func generateMarkdownTable(schema Schema, prefix string, globalRequired map[string]bool, sb *strings.Builder) *strings.Builder {

	// Mark required fields
	requiredFields := make(map[string]bool)
	for _, field := range schema.Required {
		requiredFields[field] = true
	}
	for k, v := range globalRequired {
		requiredFields[k] = v
	}

	// Traverse the properties and generate the table rows
	for key, prop := range schema.Properties {
		// Create the full name of the property (e.g., kafka.enabled)
		fullName := key
		if prefix != "" {
			fullName = prefix + "." + key
		}

		// Add type restriction details if present
		typeDetails :=  "**" + prop.Type + "**"
		if prop.SliderMin != nil || prop.SliderMax != nil|| prop.SliderUnit != "" {
			unit := "*" + prop.SliderUnit + "*"
			typeDetails += ` (slider `
			if prop.SliderMin != nil{
				typeDetails += fmt.Sprintf(`"%d%s\" ≤ `, *prop.SliderMin, unit,)
			}
			typeDetails += fmt.Sprintf(`"x%s" `, unit) 
			if prop.SliderMin != nil{
				typeDetails += fmt.Sprintf(`≤ "%d%s"`, *prop.SliderMax, unit)
			}	
			typeDetails += ")"	
		}

		if len(prop.Enum) > 0 {
			typeDetails += fmt.Sprintf(` *[enum]* <details>"%s"</details>`, strings.Join(prop.Enum, `", "`))
		}

		if prop.Pattern != "" {
			typeDetails += fmt.Sprintf(` (pattern: *"%s"*)`, prop.Pattern)
		}

		// Add a row for the property
		required := "No"
		if requiredFields[key] {
			required = "Yes"
		}
		defaultValue := extractDefaultFromHidden(prop.Hidden)

		sb.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s | %s |\n", fullName, prop.Title, typeDetails, required, defaultValue))

		// Recursively process nested properties
		if prop.Type == "object" && len(prop.Properties) > 0 {
			generateMarkdownTable(Schema{Properties: prop.Properties, Required: prop.Required}, fullName, requiredFields, sb)
		}
	}

	return sb
}
