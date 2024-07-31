package parse

import (
	"fmt"
	"slices"
	"strings"

	"github.com/danspts/helmdocs/pkg/types"
)

func extractDefaultFromHidden(hidden types.Hidden) string {
	// Check if the hidden attribute has a condition and a value
	if !hidden.Condition && hidden.Value != "" {
		// Return the default value based on the hidden attribute's value
		return hidden.Value
	}
	return ""
}

func ConvertTableToFieldTree(schema types.Property, prefix string) *types.Field {
	// Add type restriction details if present
	typeDetails := "**" + schema.Type + "**"
	if schema.SliderMin != nil || schema.SliderMax != nil || schema.SliderUnit != "" {
		unit := "*" + schema.SliderUnit + "*"
		typeDetails += ` (slider `
		if schema.SliderMin != nil {
			typeDetails += fmt.Sprintf(`"%d%s\" ≤ `, *schema.SliderMin, unit)
		}
		typeDetails += fmt.Sprintf(`"x%s" `, unit)
		if schema.SliderMin != nil {
			typeDetails += fmt.Sprintf(`≤ "%d%s"`, *schema.SliderMax, unit)
		}
		typeDetails += ")"
	}

	if len(schema.Enum) > 0 {
		typeDetails += fmt.Sprintf(` *[enum]* <details>"%s"</details>`, strings.Join(schema.Enum, `", "`))
	}

	if schema.Pattern != "" {
		typeDetails += fmt.Sprintf(` (pattern: *"%s"*)`, schema.Pattern)
	}

	f := &types.Field{
		DefaultValue: extractDefaultFromHidden(schema.Hidden),
		TypeDetails:  typeDetails,
		Title:        schema.Title,
	}

	// Traverse the properties and generate the table rows
	for key, prop := range schema.Properties {
		fullName := key
		if prefix != "" {
			fullName = prefix + "." + key
		}
		fc := ConvertTableToFieldTree(prop, fullName)
		fc.Name = key
		fc.ParentFullName = prefix
		if slices.Contains(schema.Required, key) {
			fc.Required = true
		}
		if fc.Required {
			f.Required = true
		}
		f.Children = append(f.Children, fc)
	}

	return f
}
