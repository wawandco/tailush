package tailush

import (
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
)

// FormFn returns a helper function to build forms. The resulting function can be used
// as replacement for the original form helper.
func FormFn(formOptions ...FormHelperOption) any {
	return func(opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
		if opts == nil {
			opts = tags.Options{}
		}

		hn := defaultFormVar
		if n, ok := opts["var"]; ok {
			hn = n.(string)
			delete(opts, "var")
		}

		if opts["errors"] == nil && help.Value("errors") != nil {
			opts["errors"] = help.Value("errors")
		}

		form := NewForm(opts, help)

		// Run the options on the form
		for _, optionFn := range formOptions {
			optionFn(form)
		}

		help.Set(hn, form)

		s, err := help.Block()
		if err != nil {
			return "", err
		}

		form.Append(s)

		return form.HTML(), nil
	}
}
