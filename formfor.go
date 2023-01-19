package tailush

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// func (f FormFor) buildOptions(field string, opts tags.Options) tags.Options {
// 	opts["tags-field"] = field
// 	fieldName := validators.GenerateKey(field)
// 	if err := f.Errors.Get(fieldName); err != nil {
// 		opts["errors"] = err
// 	}

// 	return opts
// }

// // CheckboxTag adds a checkbox to a form wrapped with a form-control and a label.
// func (f FormFor) CheckboxTag(field string, opts tags.Options) *tags.Tag {
// 	opts["class"] = fmt.Sprintf("%v %v", "block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm", opts["class"])
// 	value := opts["value"]
// 	if value == nil {
// 		value = "true"
// 	}

// 	checked := opts["checked"]
// 	delete(opts, "checked")

// 	opts["type"] = "checkbox"
// 	opts["name"] = field
// 	opts["value"] = value
// 	chb := tags.New("input", opts)

// 	if checked != nil {
// 		chb.Checked = checked.(bool)
// 	}

// 	if opts["tag_only"] == true {
// 		return chb
// 	}

// 	fdiv := tags.New("div", tags.Options{
// 		"class": "h-5 flex items-center text-lg",
// 		"body":  chb,
// 	})

// 	if value == "true" {
// 		fdiv.Append(tags.New("input", tags.Options{
// 			"type":  "hidden",
// 			"value": "false",
// 			"name":  field,
// 		}))
// 	}

// 	ldiv := tags.New("div", tags.Options{
// 		"class": "ml-3 text-sm",
// 	})

// 	label := tags.New("label", tags.Options{
// 		"class": "font-medium text-sixth-700",
// 		"for":   field,
// 		"body":  opts["label"],
// 	})

// 	ldiv.Append(label)
// 	ldiv.Append(tags.New("p", tags.Options{
// 		"class": "text-sixth-500",
// 		"body":  opts["description"],
// 	}))

// 	inner := tags.New("div", tags.Options{
// 		"class": "relative flex items-start",
// 	})

// 	inner.Prepend(fdiv)
// 	inner.Append(ldiv)

// 	container := tags.New("div", tags.Options{
// 		"class": opts["containerClass"],
// 		"body":  inner,
// 	})

// 	return container
// }

// // InputTag builds an input[type=text] by default wrapped with a form-control and a label.
// func (f FormFor) InputTag(field string, opts tags.Options) *tags.Tag {
// 	opts = f.buildOptions(field, opts)
// 	if opts["type"] == "hidden" {
// 		return f.HiddenTag(field, opts)
// 	}

// 	return divWrapper(opts, func(o tags.Options) tags.Body {
// 		return f.formFor.InputTag(field, opts)
// 	})
// }

// // HiddenTag adds an input[type=hidden] to the formFor.
// func (f FormFor) HiddenTag(field string, opts tags.Options) *tags.Tag {
// 	opts = f.buildOptions(field, opts)

// 	return f.formFor.HiddenTag(field, opts)
// }

// // FileTag adds a tailwindcss input[type=file] wrapped with a form-control and a label.
// func (f FormFor) FileTag(field string, opts tags.Options) *tags.Tag {
// 	opts = f.buildOptions(field, opts)

// 	return divWrapper(opts, func(o tags.Options) tags.Body {
// 		return f.formFor.FileTag(field, opts)
// 	})
// }

// // RadioButton adds a tailwindcss input[type=radio] wrapped with a form-control and a label.
// func (f FormFor) RadioButton(field string, opts tags.Options) *tags.Tag {
// 	value := opts["value"]
// 	if value == nil {
// 		value = "true"
// 	}

// 	opts["class"] = fmt.Sprintf("%v %v", "focus:ring-second-900 h-5 w-5 text-second-900 border-sixth-400", opts["class"])
// 	chb := tags.New("input", tags.Options{
// 		"type":  "radio",
// 		"name":  field,
// 		"value": value,
// 		"class": opts["class"],
// 	})

// 	if ch, ok := opts["checked"].(bool); ok {
// 		chb.Checked = ch
// 	}

// 	if opts["tag_only"] == true {
// 		return chb
// 	}

// 	fdiv := tags.New("div", tags.Options{
// 		"class": "flex items-center h-5",
// 		"body":  chb,
// 	})

// 	ldiv := tags.New("div", tags.Options{
// 		"class": "ml-3 text-sm",
// 	})

// 	label := tags.New("label", tags.Options{
// 		"class": "text-sm",
// 		"for":   field,
// 		"body":  opts["label"],
// 	})

// 	ldiv.Append(label)
// 	ldiv.Append(tags.New("p", tags.Options{
// 		"class": "text-sixth-500",
// 		"body":  opts["description"],
// 	}))

// 	inner := tags.New("div", tags.Options{
// 		"class": "relative flex items-start",
// 	})

// 	inner.Prepend(fdiv)
// 	inner.Append(ldiv)

// 	container := tags.New("div", tags.Options{
// 		"class": opts["containerClass"],
// 		"body":  inner,
// 	})

// 	return container
// }

// // RadioButtonTag adds a tailwindcss input[type=radio] wrapped with a form-control and a label.
// func (f FormFor) RadioButtonTag(field string, opts tags.Options) *tags.Tag {
// 	opts = f.buildOptions(field, opts)

// 	return divWrapper(opts, func(o tags.Options) tags.Body {
// 		return f.formFor.RadioButtonTag(field, opts)
// 	})
// }

// // SelectTag adds a tailwindcss select tag wrapped with a form-control and a label.
// func (f FormFor) SelectTag(field string, opts tags.Options) *tags.Tag {
// 	opts = f.buildOptions(field, opts)

// 	return divWrapper(opts, func(o tags.Options) tags.Body {
// 		return f.formFor.SelectTag(field, opts)
// 	})
// }

// // TextArea adds a tailwindcss textarea tag wrapped with a form-control and a label.
// func (f FormFor) TextArea(field string, opts tags.Options) *tags.Tag {
// 	return f.TextAreaTag(field, opts)
// }

// // TextAreaTag adds a tailwindcss textarea tag wrapped with a form-control and a label.
// func (f FormFor) TextAreaTag(field string, opts tags.Options) *tags.Tag {
// 	opts = f.buildOptions(field, opts)

// 	return divWrapper(opts, func(o tags.Options) tags.Body {
// 		return f.formFor.TextAreaTag(field, opts)
// 	})
// }

// // SubmitTag returns a tag for input type submit without wrapping.
// func (f FormFor) SubmitTag(value string, opts tags.Options) *tags.Tag {
// 	return f.formFor.SubmitTag(value, opts)
// }

// // NewFormFor builds a form for a passed model.
// func NewFormFor(model interface{}, opts tags.Options) *FormFor {
// 	return &FormFor{
// 		FormFor: form.NewFormFor(model, opts),
// 	}
// }

var arrayFieldRegExp = regexp.MustCompile(`^([A-Za-z0-9]+)\[(\d+)\]$`)

type interfacer interface {
	Interface() interface{}
}

type tagValuer interface {
	TagValue() string
}

// FormFor is the FormFor version for Tailwindcss.
type formFor struct {
	*form

	Model  any
	Errors tags.Errors

	name       string
	dashedName string
	reflection reflect.Value
	hctx       hctx.Context
}

// NewFormFor creates a new Formfor with passed options. It takes
// into account the passed model.
func NewFormFor(model any, opts tags.Options, hctx hctx.Context) *formFor {
	rv := reflect.ValueOf(model)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	name := rv.Type().Name()
	dashedName := flect.Dasherize(name)

	if opts["id"] == nil {
		opts["id"] = fmt.Sprintf("%s-form", dashedName)
	}

	errors := loadErrors(opts)
	delete(opts, "errors")

	return &formFor{
		form: NewForm(opts, hctx),

		Model:  model,
		Errors: errors,

		name:       name,
		dashedName: dashedName,
		reflection: rv,
		hctx:       hctx,
	}
}

func loadErrors(opts tags.Options) tags.Errors {
	errors := validate.NewErrors()

	if opts["errors"] == nil {
		return errors
	}

	switch t := opts["errors"].(type) {
	default:
		fmt.Printf("Unexpected errors type %T, please\n", t) // %T prints whatever type t has
	case map[string][]string:
		errors = &validate.Errors{
			Errors: opts["errors"].(map[string][]string),
			Lock:   new(sync.RWMutex),
		}
	case tags.Errors:
		return opts["errors"].(tags.Errors)
	}

	return errors
}

// CheckboxTag creates a checkbox for a field on the form Struct
func (f formFor) CheckboxTag(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)
	return f.form.CheckboxTag(opts)
}

// InputTag creates an input for a field on the form Struct
func (f formFor) InputTag(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)
	f.addFormatTag(field, opts)

	return divWrapper(opts, func(opts tags.Options) tags.Body {
		return f.form.InputTag(opts)
	})
}

// HiddenTag adds a wrappter for input type hidden on the form
func (f formFor) HiddenTag(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)
	return f.form.HiddenTag(opts)
}

// FileTag creates a input[type=file] for a field name passed
func (f formFor) FileTag(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)
	f.addFormatTag(field, opts)
	return f.form.FileTag(opts)
}

// DateTimeTag creates a input[type=datetime-local] for a field name passed
func (f formFor) DateTimeTag(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)
	f.addFormatTag(field, opts)
	return f.form.DateTimeTag(opts)
}

// RadioButton creates a radio button for a struct field
func (f formFor) RadioButton(field string, opts tags.Options) *tags.Tag {
	return f.RadioButtonTag(field, opts)
}

// RadioButtonTag creates a radio button for a struct field
func (f formFor) RadioButtonTag(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)
	return f.form.RadioButtonTag(opts)
}

// SelectTag creates a select tag for a specified struct field and loads
// options from the options map
func (f formFor) SelectTag(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)

	return divWrapper(opts, func(opts tags.Options) tags.Body {
		return f.form.SelectTag(opts)
	})
}

// TextArea creates text area for the specified struct field
func (f formFor) TextArea(field string, opts tags.Options) *tags.Tag {
	return f.TextAreaTag(field, opts)
}

// TextAreaTag creates text area for the specified struct field
func (f formFor) TextAreaTag(field string, opts tags.Options) *tags.Tag {
	f.buildOptions(field, opts)

	return f.form.TextArea(opts)
}

// SubmitTag adds a submit button to the form
func (f formFor) SubmitTag(value string, opts tags.Options) *tags.Tag {
	return f.form.SubmitTag(value, opts)
}

func (f formFor) buildOptions(field string, opts tags.Options) {
	opts["tags-field"] = field

	if opts["value"] == nil {
		opts["value"] = f.value(field)
	}

	if opts["name"] == nil {
		opts["name"] = f.findFieldNameFor(field)
	}

	if opts["id"] == nil {
		opts["id"] = fmt.Sprintf("%s-%s", f.dashedName, opts["name"])
	}
}

func (f formFor) addFormatTag(field string, opts tags.Options) {
	if opts["format"] != nil {
		return
	}

	toff := reflect.TypeOf(f.Model)
	if toff.Kind() == reflect.Ptr {
		toff = toff.Elem()
	}

	if toff.Kind() == reflect.Struct {
		fi, found := toff.FieldByName(field)

		if !found {
			return
		}

		if format, ok := fi.Tag.Lookup("format"); ok && format != "" {
			opts["format"] = format
		}
	}
}

func (f formFor) value(field string) interface{} {
	fn := f.reflection.FieldByName(field)

	if !fn.IsValid() {
		dots := strings.Split(field, ".")

		if len(dots) == 1 && !arrayFieldRegExp.Match([]byte(dots[0])) {
			if !strings.HasSuffix(field, "ID") {
				return f.value(field + "ID")
			}
			return ""
		}

		matches := arrayFieldRegExp.FindStringSubmatch(dots[0])
		if len(matches) != 0 {
			dots[0] = matches[1]
		}

		fn = f.reflection.FieldByName(dots[0])

		if fn.IsValid() {
			fn = reflect.Indirect(fn)

			if fn.Kind() == reflect.Slice || fn.Kind() == reflect.Array {
				index, _ := strconv.Atoi(matches[2])
				fn = reflect.Indirect(fn.Index(index))
			}

			if fn.Kind() == reflect.Struct {
				ff := NewFormFor(fn.Interface(), f.Options, f.hctx)
				return ff.value(strings.Join(dots[1:], "."))
			}
		}
	}

	fn = reflect.Indirect(fn)

	i := fn.Interface()
	switch t := i.(type) {
	case uuid.UUID:
		return t.String()
	case tagValuer:
		return t.TagValue()
	case driver.Valuer:
		value, _ := t.Value()

		if value == nil {
			return ""
		}

		return fmt.Sprintf("%v", value)
	case interfacer:
		return fmt.Sprintf("%v", t.Interface())
	}
	return i
}

func (f formFor) findFieldNameFor(field string) string {
	ty := reflect.TypeOf(f.Model)

	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
	}

	rf, ok := ty.FieldByName(field)
	if !ok {
		if rf, ok = ty.FieldByName(field + "ID"); !ok {
			return field
		}
		field = field + "ID"
	}

	formDefined := string(rf.Tag.Get("form"))
	if formDefined != "" && formDefined != "-" {
		return formDefined
	}

	schemaDefined := string(rf.Tag.Get("schema"))
	if schemaDefined != "" && schemaDefined != "-" {
		return schemaDefined
	}

	return field
}
