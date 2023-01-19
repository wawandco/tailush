package tailush

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/flect"
	"github.com/gobuffalo/tags/v3"
)

var (
	//csrfTokenKey is the key where the authenticity token
	//is stored in the helper context
	csrfTokenKey = "authenticity_token"

	//defaultFormVar is the default variable name for
	//the form in the push template.
	defaultFormVar = "f"
)

// divWrapper
func divWrapper(opts tags.Options, fn func(opts tags.Options) tags.Body) *tags.Tag {
	divClass := opts["containerClass"]
	delete(opts, "containerClass")

	hasErrors := false
	errors := []string{}
	if divClass == nil {
		divClass = "mb-2 w-full"
	}

	if opts["errors"] != nil && len(opts["errors"].([]string)) > 0 {
		divClass = fmt.Sprintf("%v %v", divClass, "has-error")
		hasErrors = true
		errors = opts["errors"].([]string)

		delete(opts, "errors")
	}

	div := tags.New("div", tags.Options{
		"class": divClass,
	})

	if opts["label"] == nil && opts["tags-field"] != nil {
		if tf, ok := opts["tags-field"].(string); ok {
			tf = strings.Join(strings.Split(tf, "."), " ")
			opts["label"] = flect.Titleize(tf)
		}
	}

	delete(opts, "tags-field")

	useLabel := opts["hide_label"] == nil
	if useLabel {
		div.Prepend(tags.New("label", tags.Options{
			"body":  opts["label"],
			"class": "block text-sm font-medium text-gray-700 mb-1",
		}))

		delete(opts, "hide_label")
	}

	delete(opts, "label")
	delete(opts, "hide_label")

	idiv := tags.New("div", tags.Options{
		"class": divClass,
	})

	// buildOptions(opts, hasErrors)

	if opts["tag_only"] == true {
		return fn(opts).(*tags.Tag)
	}

	idiv.Append(fn(opts))
	div.Append(idiv)

	if !hasErrors {
		return div
	}

	idiv.Append(tags.New("p", tags.Options{
		"class": "mt-1 text-xs text-red-600 mb-2",
		"body":  strings.Join(errors, ". "),
	}))

	return div
}
