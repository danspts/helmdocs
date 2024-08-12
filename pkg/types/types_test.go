package types

import (
	"reflect"
	"testing"
)

func TestField_Flatten(t *testing.T) {
	tests := []struct {
		name   string
		field  *Field
		expect []*FlatField
	}{
		{
			name:   "Empty field",
			field:  &Field{},
			expect: []*FlatField{},
		},
		{
			name: "Single field with no children",
			field: &Field{
				Children: []*Field{{Name: "root"}}},
			expect: []*FlatField{
				{Field: &Field{Name: "root"}, depth: 0},
			},
		},
		{
			name: "Field with children",
			field: &Field{
				Children: []*Field{{
					Name: "root",
					Children: []*Field{
						{Name: "child1"},
						{Name: "child2"},
					},
				}}},
			expect: []*FlatField{
				{Field: &Field{Name: "root",
					Children: []*Field{
						{Name: "child1"},
						{Name: "child2"},
					},
				}, depth: 0},
				{Field: &Field{Name: "child1"}, depth: 1},
				{Field: &Field{Name: "child2"}, depth: 1},
			},
		},
		{
			name: "Field with nested children",
			field: &Field{
				Children: []*Field{{
					Name: "root",
					Children: []*Field{
						{
							Name: "child1",
							Children: []*Field{
								{Name: "grandchild1"},
							},
						},
					},
				}}},
			expect: []*FlatField{
				{Field: &Field{Name: "root", Children: []*Field{
					{
						Name: "child1",
						Children: []*Field{
							{Name: "grandchild1"},
						},
					},
				}}, depth: 0},
				{Field: &Field{Name: "child1", Children: []*Field{
					{Name: "grandchild1"},
				}}, depth: 1},
				{Field: &Field{Name: "grandchild1"}, depth: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.Flatten()
			if !reflect.DeepEqual(got, tt.expect) {
				t.Errorf("Flatten() = %v, expect %v", got, tt.expect)
			}
		})
	}
}

func TestFlatField_Depth(t *testing.T) {
	tests := []struct {
		name   string
		field  *FlatField
		expect int
	}{
		{
			name:   "Depth zero",
			field:  &FlatField{Field: &Field{Name: "root"}, depth: 0},
			expect: 0,
		},
		{
			name:   "Depth one",
			field:  &FlatField{Field: &Field{Name: "child1"}, depth: 1},
			expect: 1,
		},
		{
			name:   "Depth two",
			field:  &FlatField{Field: &Field{Name: "grandchild1"}, depth: 2},
			expect: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.field.Depth()
			if got != tt.expect {
				t.Errorf("Depth() = %v, expect %v", got, tt.expect)
			}
		})
	}
}
