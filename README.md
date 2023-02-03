# Tailush

Tailush is a library with form plush helpers for TailwindCSS forms and inputs. Its a drop in replacement for the `form` and `form_for` helpers that ship with Buffalo and are oriented towards Bootstrap. The goal of this library is to accelerate the development of web applications by make it simpler to develop forms with TailwindCSS in plush.

# Installation
    
```bash
go get github.com/wawandco/tailush@latest
```

### Usage

To install Tailush you should replace the `form` and `form_for` helpers in your `actions/helpers.go` file.

```go
// In your helpers.go
import (
    ...
    "tailush"
    ...
)

... 
    // Custom Tailwindcss form helpers
    "formFor": tailush.FormForFn(),
    "form":    tailush.FormFn(),
...

```

### Customizing

Customizing the form helpers is done by passing different instances of `tailush.Option` to the `FormFn` and `FormForFn` functions. The following example shows a set of options that can be passed to the helpers.

```go
... 
    // Custom Tailwindcss form helpers
    "formFor": tailush.FormForFn(
        tailush.UseFieldContainerClass("mb-4"),
        tailush.UseLabelClass("block text-gray-700 text-sm font-bold mb-2"),
        tailush.UseInputClass("shadow border rounded-md w-full p-2 text-gray-700 focus:outline-none focus:shadow-outline"),
        tailush.UseErrorClass("text-red-500 text-xs italic"),
    ),
    ...
...
```

#### Customizing options

The following options are available for the form helpers:

| Option ⚙️                           | Component ⚒️                 | Preview 🔎 |
|------------------------------------|-----------------------------|------------|
| `tailush.UseLabelClass()`          | `<%= f.Label() %>`          |            |
| `tailush.UseInputClass()`          | `<%= f.InputTag() %>`       |            |
| `tailush.UseCheckboxClass()`       | `<%= f.CheckboxTag() %>`    |            |
| `tailush.UseRadioClass()`          | `<%= f.RadioButtonTag() %>` |            |
| `tailush.UseFileClass()`           | `<%= f.FileTag() %>`        |            |
| `tailush.UseTextAreaClass()`       | `<%= f.TextAreaTag() %>`    |            |
| `tailush.UseSelectClass()`         | `<%= f.SelectTag() %>`      |            |
| `tailush.UseDateInputClass()`      | `<%= f.DateTimeTag() %>`    |            |
| `tailush.UseFieldContainerClass()` | Not implemented yet         |            |
| `tailush.UseErrorClass()`          | Not implemented yet         |            |

