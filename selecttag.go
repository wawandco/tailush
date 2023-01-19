package tailush

import (
	"html/template"
	"reflect"

	"github.com/gobuffalo/tags/v3"
)

// selectTag describes a HTML <select> tag meta data.
type selectTag struct {
	*tags.Tag
	SelectedValue      interface{}
	selectedValueCache map[interface{}]struct{}
	SelectOptions      selectOptions
}

func (s selectTag) String() string {
	return s.Tag.String()
}

// HTML gives the HTML template representation for the select tag.
func (s selectTag) HTML() template.HTML {
	return template.HTML(s.String())
}

// NewSelectTag constructs a new `<select>` tag.
func NewSelectTag(opts tags.Options) *tags.Tag {
	so := parseSelectOptions(opts)
	selected := opts["value"]
	delete(opts, "value")

	// Transform selected value(s) into an empty map with values as keys
	// (faster lookup than slice / array)
	selectedMap := make(map[interface{}]struct{})

	multiple, ok := opts["multiple"].(bool)
	if multiple && ok {
		// Set nil to use the empty attribute notation
		opts["multiple"] = nil

		rv := reflect.ValueOf(selected)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}
		if rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice {
			for i := 0; i < rv.Len(); i++ {
				x := rv.Index(i).Interface()
				if s, ok := x.(Selectable); ok {
					// Use Selectable value as the selected value
					x = s.SelectValue()
				}
				selectedMap[template.HTMLEscaper(x)] = struct{}{}
			}
		} else {
			// Set unique value as a map key
			selectedMap[template.HTMLEscaper(selected)] = struct{}{}
		}
	} else {
		if s, ok := selected.(Selectable); ok {
			selected = s.SelectValue()
		}
		// Set unique value as a map key
		selectedMap[template.HTMLEscaper(selected)] = struct{}{}
	}

	delete(opts, "tag_only")

	st := &selectTag{
		Tag:                tags.New("select", opts),
		SelectOptions:      so,
		SelectedValue:      selected,
		selectedValueCache: selectedMap,
	}

	for _, x := range st.SelectOptions {
		if _, ok := st.selectedValueCache[template.HTMLEscaper(x.Value)]; ok {
			x.Selected = true
		}
		st.Append(x.String())
	}

	return st.Tag
}

func parseSelectOptions(opts tags.Options) selectOptions {
	if opts["options"] == nil {
		return selectOptions{}
	}

	allowBlank := opts["allow_blank"]
	delete(opts, "allow_blank")

	sopts := opts["options"]
	delete(opts, "options")

	placeHolder := opts["placeholder"]
	delete(opts, "placeholder")

	so := selectOptions{}
	if ph, ok := placeHolder.(string); ok {
		so = append(so, selectOption{
			Value:       "",
			Label:       ph,
			Placeholder: true,
		})
	}

	if aw, ok := allowBlank.(bool); ok && aw {
		so = append(so, selectOption{
			Value: "",
			Label: "",
		})
	}

	if x, ok := sopts.(selectOptions); ok {
		x = append(so, x...) // prepend placerholder or blank SelectOption if present
		return x
	}

	rv := reflect.ValueOf(sopts)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		selectableType := reflect.TypeOf((*Selectable)(nil)).Elem()

		for i := 0; i < rv.Len(); i++ {
			x := rv.Index(i).Interface()

			if rv.Index(i).Type().Implements(selectableType) {
				isSelected := false
				selectableMultipleType := reflect.TypeOf((*SelectableMultiple)(nil)).Elem()
				if rv.Index(i).Type().Implements(selectableMultipleType) {
					isSelected = x.(SelectableMultiple).IsSelected()
				}
				so = append(so, selectOption{Value: x.(Selectable).SelectValue(), Label: x.(Selectable).SelectLabel(), Selected: isSelected})
				continue
			}

			if m, ok := rv.Index(i).Interface().(map[string]interface{}); ok {
				for k, v := range m {
					so = append(so, selectOption{Value: v, Label: k})
				}

				continue
			}

			so = append(so, selectOption{Value: x, Label: x})
		}
	case reflect.Map:
		keys := rv.MapKeys()
		for i := 0; i < len(keys); i++ {
			k := keys[i]
			so = append(so, selectOption{
				Value: rv.MapIndex(k).Interface(),
				Label: k.Interface(),
			})
		}
	}
	return so
}

// Selectable allows any struct to become an option in the select tag.
type Selectable interface {
	SelectValue() interface{}
	SelectLabel() string
}

// Selectables is the plural for selectable
type Selectables []Selectable

// SelectableMultiple allows any struct to add Selected option in the select tag.
type SelectableMultiple interface {
	IsSelected() bool
}

// SelectableMultiples is the plural for SelectableMultiple
type SelectableMultiples []SelectableMultiple
