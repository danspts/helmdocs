package readme

import (
	"strings"
	"testing"

	"github.com/danspts/helmdocs/pkg/types"
)

func TestGenerateMarkdownTableWithHeader(t *testing.T) {
	tests := []struct {
		name   string
		schema types.Schema
		expect string
	}{
		{
			name: "Empty schema",
			schema: types.Schema{
				Type:       "object",
				Properties: map[string]types.Property{},
			},
			expect: "| Name | Title | Type | Required | Default |\n|------|-------|------|----------|---------|\n",
		},
		{
			name: "Single field schema",
			schema: types.Schema{
				Type: "object",
				Properties: map[string]types.Property{
					"field1": {
						Schema: types.Schema{Type: "string"},
						Title:  "Field 1",
					},
				},
				Required: []string{"field1"},
			},
			expect: "| Name | Title | Type | Required | Default |\n|------|-------|------|----------|---------|\n| `field1` | Field 1 | **string** | true |  |\n",
		},
		{
			name: "Nested field schema",
			schema: types.Schema{
				Type: "object",
				Properties: map[string]types.Property{
					"field1": {
						Schema: types.Schema{
							Type: "object",
							Properties: map[string]types.Property{
								"nestedField1": {
									Schema: types.Schema{Type: "integer"},
									Title:  "Nested Field 1",
								},
							},
						},
						Title: "Field 1",
					},
				},
			},
			expect: "| Name | Title | Type | Required | Default |\n|------|-------|------|----------|---------|\n| `field1` | Field 1 | **object** | false |  |\n| `field1.nestedField1` | Nested Field 1 | **integer** | false |  |\n",
		},
		{
			name: "Field with default value",
			schema: types.Schema{
				Type: "object",
				Properties: map[string]types.Property{
					"field1": {
						Schema:  types.Schema{Type: "string"},
						Title:   "Field 1",
						Default: "default value",
					},
				},
			},
			expect: "| Name | Title | Type | Required | Default |\n|------|-------|------|----------|---------|\n| `field1` | Field 1 | **string** | false | default value |\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateReadme(tt.schema)
			if strings.TrimSpace(got) != strings.TrimSpace(tt.expect) {
				t.Errorf("generateMarkdownTableWithHeader() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
