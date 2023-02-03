package tailush

import (
	"html/template"

	"github.com/gobuffalo/helpers/hctx"
	"github.com/gobuffalo/tags/v3"
)

// FormForFn returns a helper function to build forms. The resulting function can be used
// as replacement for the original form helper.
func FormForFn(formOptions ...FormHelperOption) any {
	return func(model any, opts tags.Options, help hctx.HelperContext) (template.HTML, error) {
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

		// Run the options on the form
		formf := NewFormFor(model, opts, help)
		for _, optionFn := range formOptions {
			optionFn(formf.form)
		}

		help.Set(hn, formf)

		s, err := help.Block()
		if err != nil {
			return "", err
		}

		formf.Append(s)

		return formf.HTML(), nil
	}
}
