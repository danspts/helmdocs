package types

import "encoding/json"

type Schema struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required,omitempty"`
}

type Property struct {
	Schema
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	Form        bool        `json:"form,omitempty"`
	Render      string      `json:"render,omitempty"`
	Enum        []string    `json:"enum,omitempty"`
	Pattern     string      `json:"pattern,omitempty"`
	SliderMin   *int        `json:"sliderMin,omitempty"`
	SliderMax   *int        `json:"sliderMax,omitempty"`
	SliderUnit  string      `json:"sliderUnit,omitempty"`
	Default     interface{} `json:"default,omitempty"`
	Hidden      Hidden      `json:"hidden,omitempty"`
}

type Hidden struct {
	Condition bool `json:"condition"`
	Value     any  `json:"value"`
}

func (h *Hidden) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		type Alias Hidden
		aux := &struct{ *Alias }{Alias: (*Alias)(h)}
		return json.Unmarshal(data, &aux)
	}
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	h.Condition = false
	h.Value = str
	return nil
}

type (
	Field struct {
		Name           string
		ParentFullName string
		Title          string
		Type           string
		TypeDetails    string
		Required       bool
		DefaultValue   string
		Children       []*Field
	}
	FlatField struct {
		*Field
		depth int
	}
)

func (f *Field) flatten(depth int) []*FlatField {
	ff := &FlatField{Field: f, depth: depth}
	flat := []*FlatField{ff}
	for _, c := range f.Children {
		flat = append(flat, c.flatten(depth+1)...)
	}
	return flat
}

func (f *Field) Flatten() []*FlatField {
	flat := []*FlatField{}
	for _, c := range f.Children {
		flat = append(flat, c.flatten(0)...)
	}
	return flat
}

func (f *FlatField) Depth() int {
	return f.depth
}
