package values

import (
	"strings"
	"testing"

	"github.com/danspts/helmdocs/pkg/types"
)

func TestGenerateMarkdownTableWithHeader_OmitDft(t *testing.T) {
	tests := []struct {
		name    string
		schema  types.Schema
		omitDft bool
		expect  string
	}{
		{
			name: "Empty schema without omitDft",
			schema: types.Schema{
				Type:       "object",
				Properties: map[string]types.Property{},
			},
			omitDft: false,
			expect:  "",
		},
		{
			name: "Single field schema without omitDft",
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
			omitDft: false,
			expect:  "field1: <string>\n",
		},
		{
			name: "Single field schema with omitDft",
			schema: types.Schema{
				Type: "object",
				Properties: map[string]types.Property{
					"field1": {
						Schema:  types.Schema{Type: "string"},
						Title:   "Field 1",
						Default: "default value",
					},
				},
				Required: []string{"field1"},
			},
			omitDft: true,
			expect:  "# field1: default value\n",
		},
		{
			name: "Nested field schema with omitDft",
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
			omitDft: false,
			expect:  "field1:\n  nestedField1: <integer>\n",
		},
		{
			name: "Field with default value without omitDft",
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
			omitDft: false,
			expect:  "field1: default value\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateValues(tt.schema, tt.omitDft)
			if strings.TrimSpace(got) != strings.TrimSpace(tt.expect) {
				t.Errorf("generateMarkdownTableWithHeader() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
