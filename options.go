package tailush

type FormHelperOption func(*form)

// UseFieldContainerClass is used to override the default classes for the container element.
func UseFieldContainerClass(class string) FormHelperOption {
	return func(f *form) {
		f.fieldContainerClass = class
	}
}

// UseLabelClass is used to override the default classes for the label element.
func UseLabelClass(class string) FormHelperOption {
	return func(f *form) {
		f.labelClass = class
	}
}

// UseInputClass is used to override the default classes for the input element.
func UseInputClass(class string) FormHelperOption {
	return func(f *form) {
		f.inputClass = class
	}
}

// UseCheckboxClass is used to override the default classes for the checkbox element.
func UseCheckboxClass(class string) FormHelperOption {
	return func(f *form) {
		f.checkboxClass = class
	}
}

// UseRadioClass is used to override the default classes for the radio element.
func UseRadioClass(class string) FormHelperOption {
	return func(f *form) {
		f.radioClass = class
	}
}

// UseFileClass is used to override the default classes for the file element.
func UseFileClass(class string) FormHelperOption {
	return func(f *form) {
		f.fileClass = class
	}
}

// UseTextAreaClass is used to override the default classes for the textarea element.
func UseTextAreaClass(class string) FormHelperOption {
	return func(f *form) {
		f.textAreaClass = class
	}
}

// UseSelectClass is used to override the default classes for the select element.
func UseSelectClass(class string) FormHelperOption {
	return func(f *form) {
		f.selectClass = class
	}
}

// UseErrorClass is used to override the default classes for the errors.
func UseErrorClass(class string) FormHelperOption {
	return func(f *form) {
		f.errorClass = class
	}
}

// UseDateInputClass is used to override the default classes for the errors.
func UseDateInputClass(class string) FormHelperOption {
	return func(f *form) {
		f.dateInputClass = class
	}
}
