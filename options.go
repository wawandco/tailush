package tailush

type FormHelperOption func(*form)

func UseFieldContainerClass(class string) FormHelperOption {
	return func(f *form) {
		f.fieldContainerClass = class
	}
}

func UseLabelClass(class string) FormHelperOption {
	return func(f *form) {
		f.labelClass = class
	}
}

func UseInputClass(class string) FormHelperOption {
	return func(f *form) {
		f.inputClass = class
	}
}
