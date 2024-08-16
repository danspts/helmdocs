package parse

import (
	"reflect"
	"testing"

	"github.com/danspts/helmdocs/pkg/types"
)

func TestConvertTableToFieldTree(t *testing.T) {
	tests := []struct {
		name   string
		schema types.Property
		prefix string
		expect *types.Field
	}{
		{
			name: "Empty schema",
			schema: types.Property{
				Schema: types.Schema{
					Type:       "object",
					Properties: map[string]types.Property{},
				},
			},
			prefix: "",
			expect: &types.Field{
				Type:     "object",
				Children: nil,
			},
		},
		{
			name: "Schema with a single property",
			schema: types.Property{
				Schema: types.Schema{
					Type: "object",
					Properties: map[string]types.Property{
						"field1": {
							Schema: types.Schema{Type: "string"},
							Title:  "Field 1",
						},
					},
				},
			},
			prefix: "",
			expect: &types.Field{
				Type: "object",
				Children: []*types.Field{
					{
						Name:           "field1",
						Type:           "string",
						Title:          "Field 1",
						ParentFullName: "",
						Children:       nil,
					},
				},
			},
		},
		{
			name: "Schema with a nested property",
			schema: types.Property{
				Schema: types.Schema{
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
			},
			prefix: "",
			expect: &types.Field{
				Type: "object",
				Children: []*types.Field{
					{
						Name:           "field1",
						Type:           "object",
						Title:          "Field 1",
						ParentFullName: "",
						Children: []*types.Field{
							{
								Name:           "nestedField1",
								Type:           "integer",
								Title:          "Nested Field 1",
								ParentFullName: "field1",
								Children:       nil,
							},
						},
					},
				},
			},
		},
		{
			name: "Schema with type restrictions and enum",
			schema: types.Property{
				Schema: types.Schema{
					Type: "object",
					Properties: map[string]types.Property{
						"field1": {
							Schema: types.Schema{
								Type: "integer",
							},
							Enum:  []string{"1", "2", "3"},
							Title: "Field 1",
						},
					},
				},
			},
			prefix: "",
			expect: &types.Field{
				Type: "object",
				Children: []*types.Field{
					{
						Name:           "field1",
						Type:           "integer",
						Title:          "Field 1",
						TypeDetails:    ` *[enum]* <details>"1", "2", "3"</details>`,
						ParentFullName: "",
						Children:       nil,
					},
				},
			},
		},
		{
			name: "Schema with a pattern",
			schema: types.Property{
				Pattern: "^[a-zA-Z0-9]+$",
				Schema: types.Schema{
					Type: "string",
				},
				Title: "Field 1",
			},
			prefix: "",
			expect: &types.Field{
				Type:        "string",
				Title:       "Field 1",
				TypeDetails: `(pattern: *"^[a-zA-Z0-9]+$"*)`,
				Children:    nil,
			},
		},
		{
			name: "Schema with a pattern",
			schema: types.Property{
				SliderMin:  func(i int) *int { return &i }(1),
				SliderMax:  func(i int) *int { return &i }(10),
				SliderUnit: "x",
				Schema: types.Schema{
					Type: "string",
				},
				Title: "Field 1",
			},
			prefix: "",
			expect: &types.Field{
				Type:        "string",
				Title:       "Field 1",
				TypeDetails: `(slider "1*x*\" ≤ "x*x*" ≤ "10*x*")`,
				Children:    nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertTableToFieldTree(tt.schema, tt.prefix)
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("ConvertTableToFieldTree() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
