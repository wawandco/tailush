package tailush

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
)

// form is the tailwindcss version of form
type form struct {
	*tags.Tag

	fieldContainerClass string
	labelClass          string
	inputClass          string
	selectClass         string
}

// SetAuthenticityToken allows to set the authenticity token
// in the form, that way CSRF will work ok there by adding
// a hidden input with the passed token.
func (f *form) SetAuthToken(s string) {
	f.Prepend(tags.New("input", tags.Options{
		"value": s,
		"type":  "hidden",
		"name":  csrfTokenKey,
	}))
}

// Label permits to create a label inside a Form
func (f form) Label(value string, opts tags.Options) *tags.Tag {
	opts["body"] = value

	// Set the class for the label
	if cl := opts["class"]; cl == nil {
		opts["class"] = f.labelClass
	} else {
		opts["class"] = fmt.Sprintf("%v %v", f.labelClass, cl)
	}

	return tags.New("label", opts)
}

// InputTag builds a tailwindcss input with passed options.
// It calls the appropriated input builder based on the type
// assigned in the `type` options.
func (f form) InputTag(opts tags.Options) *tags.Tag {
	switch opts["type"] {
	case "checkbox":
		return f.CheckboxTag(opts)
	case "radio":
		return f.RadioButtonTag(opts)
	case "file":
		return f.FileTag(opts)
	case "hidden":
		return f.HiddenTag(opts)
	default:
		if opts["type"] == nil {
			opts["type"] = "text"
		}

		opts["class"] = f.inputClass
		delete(opts, "tag_only")

		return tags.New("input", opts)
	}
}

// CheckboxTag builds a tailwindcss checkbox with passed options
func (f form) CheckboxTag(opts tags.Options) *tags.Tag {
	opts["type"] = "checkbox"

	value := opts["value"]
	delete(opts, "value")

	checked := opts["checked"]
	delete(opts, "checked")
	if checked == nil {
		checked = "true"
	}
	opts["value"] = checked

	unchecked := opts["unchecked"]
	delete(opts, "unchecked")

	hl := opts["hide_label"]
	delete(opts, "hide_label")

	if opts["tag_only"] == true {
		delete(opts, "label")
		ct := f.InputTag(opts)
		ct.Checked = template.HTMLEscaper(value) == template.HTMLEscaper(checked)
		return ct
	}

	tag := tags.New("label", tags.Options{})
	ct := f.InputTag(opts)
	ct.Checked = template.HTMLEscaper(value) == template.HTMLEscaper(checked)
	tag.Append(ct)

	if opts["name"] != nil && unchecked != nil {
		tag.Append(tags.New("input", tags.Options{
			"type":  "hidden",
			"name":  opts["name"],
			"value": unchecked,
		}))
	}

	if opts["label"] != nil && hl == nil {
		label := fmt.Sprint(opts["label"])
		delete(opts, "label")
		tag.Append(" " + label)
	}
	return tag
}

// FileTag builds an input[type=file] from the options passed, i
// itt also sets the form enctype to be "multipart/form-data"
func (f form) FileTag(opts tags.Options) *tags.Tag {
	f.Options["enctype"] = "multipart/form-data"
	opts["type"] = "file"

	return tags.New("input", opts)
}

// RadioButton builds a tailwindcss input[type=radio] with passed options
func (f form) RadioButtonTag(opts tags.Options) *tags.Tag {
	opts["type"] = "radio"

	var label string
	if opts["label"] != nil {
		label = fmt.Sprint(opts["label"])
		delete(opts, "label")
	}

	var ID string
	if opts["id"] != nil {
		ID = fmt.Sprint(opts["id"])
	}

	value := opts["value"]
	checked := opts["checked"]
	delete(opts, "checked")

	if opts["tag_only"] == true {
		ct := f.InputTag(opts)
		ct.Checked = template.HTMLEscaper(value) == template.HTMLEscaper(checked)
		return ct
	}

	ct := f.InputTag(opts)
	ct.Checked = template.HTMLEscaper(value) == template.HTMLEscaper(checked)
	labelOptions := tags.Options{
		"body": strings.Join([]string{ct.String(), label}, " "),
	}

	// If the ID is provided, give it to the label's for attribute
	if ID != "" {
		labelOptions["for"] = ID
	}

	tag := tags.New("label", labelOptions)
	return tag
}

// SelectTag constructs a new `<select>` tag from a form.
func (f form) SelectTag(opts tags.Options) *tags.Tag {
	opts["class"] = f.selectClass

	return NewSelectTag(opts)
}

// TextArea creates a textarea for a form with passed options
func (f form) TextArea(opts tags.Options) *tags.Tag {
	return f.TextAreaTag(opts)
}

// TextAreaTag creates a textarea for a form with passed options
func (f form) TextAreaTag(opts tags.Options) *tags.Tag {
	if opts["value"] != nil {
		opts["encoded_body"] = opts["value"]

		delete(opts, "value")
	}

	delete(opts, "tag_only")

	return tags.New("textarea", opts)
}

// DateTimeTag generates a tag with type datetime-local
// and adds default format to be 2006-01-02T03:04
func (f form) DateTimeTag(opts tags.Options) *tags.Tag {
	if opts["type"] == nil {
		opts["type"] = "datetime-local"
	}

	if opts["format"] == nil {
		opts["format"] = "2006-01-02T03:04"
	}

	delete(opts, "tag_only")
	return tags.New("input", opts)
}

// HiddenTag adds a hidden input to the form
func (f form) HiddenTag(opts tags.Options) *tags.Tag {
	delete(opts, "tag_only")
	opts["type"] = "hidden"

	return tags.New("input", opts)
}

// SubmitTag generates an input tag with type "submit"
func (f form) SubmitTag(value string, opts tags.Options) *tags.Tag {
	opts["type"] = "submit"
	opts["value"] = value

	return tags.New("input", opts)
}

// NewForm creates a new form from passed options, it sets POST as the
// default method and also handles other methods as PUT by adding
// a `_method` hidden input.
func NewForm(opts tags.Options, help hctx.Context) *form {
	if opts["method"] == nil {
		opts["method"] = "POST"
	}

	if opts["multipart"] != nil {
		opts["enctype"] = "multipart/form-data"

		delete(opts, "multipart")
	}

	form := &form{
		labelClass:  "block text-sm font-medium text-gray-700",
		selectClass: "mt-1 block w-full rounded-md border-gray-300 py-2 pl-3 pr-10 text-base focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm",
		inputClass:  "border border-gray-300 rounded-md py-1.5 px-3 w-full",

		Tag: tags.New("form", opts),
	}

	m := strings.ToUpper(form.Options["method"].(string))
	if m == http.MethodPost || m == http.MethodGet {
		return form
	}

	form.Options["method"] = "POST"
	form.Prepend(tags.New("input", tags.Options{
		"value": m,
		"type":  "hidden",
		"name":  "_method",
	}))

	tok := help.Value(csrfTokenKey)
	if tok == nil {
		return form
	}

	v, ok := tok.(string)
	if !ok {
		return form
	}

	form.SetAuthToken(v)
	return form
}
