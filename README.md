<p align="center"><img src="https://raw.githubusercontent.com/wawandco/tailush/main/logo.png" width="300px" height="300px">
</p>

<p align="center"><a href="https://pkg.go.dev/github.com/wawandco/tailush">
<img src="https://github.com/wawandco/tailush/actions/workflows/tests.yml/badge.svg"/>
</a><a href="https://pkg.go.dev/github.com/wawandco/tailush">
        <img src="https://pkg.go.dev/badge/github.com/wawandco/tailush.svg" alt="Go Reference"/>
    </a>
</p>

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

#### Customizing

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


