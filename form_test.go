package tailush_test

import (
	"net/http"
	"strings"
	"tailush"
	"testing"

	"github.com/gobuffalo/plush/v4"
	"github.com/gobuffalo/tags/v3"
)

func newHctx() plush.HelperContext {
	return plush.HelperContext{Context: plush.NewContext()}
}

func TestNew(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		if f == nil {
			t.Fatalf("f should not be nil")
		}

		h := f.HTML()
		if !strings.Contains(string(h), "form") {
			t.Fatalf("f.HTML() should contain 'form'")
		}
	})

	t.Run("method not specified defaults to POST", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		if f.Options["method"] != "POST" {
			t.Fatalf("f.Options[\"method\"] should be 'POST'")
		}
	})

	t.Run("get should be respected", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{"method": http.MethodGet}, newHctx())
		if f.Options["method"] != http.MethodGet {
			t.Fatalf("f.Options[\"method\"] should be 'GET'")
		}

		h := f.HTML()
		if strings.Contains(string(h), "_method") {
			t.Fatalf("f.HTML() should contain not '_method' field")
		}
	})

	t.Run("not get/post should be translated to '_method'", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{"method": http.MethodPut}, newHctx())
		if f.Options["method"] != "POST" {
			t.Fatalf("f.Options[\"method\"] should be 'POST'")
		}

		h := f.HTML()
		if !strings.Contains(string(h), `<input name="_method" type="hidden" value="PUT" />`) {
			t.Fatalf("f.HTML() should contain '_method' field")
		}
	})

	t.Run("not get/post should be translated to '_method'", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{"method": http.MethodPut}, newHctx())
		if f.Options["method"] != "POST" {
			t.Fatalf("f.Options[\"method\"] should be 'POST'")
		}

		h := f.HTML()
		if !strings.Contains(string(h), `<input name="_method" type="hidden" value="PUT" />`) {
			t.Fatalf("f.HTML() should contain '_method' field")
		}
	})

	t.Run("contains multipart", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{"multipart": "true"}, newHctx())
		if f.Options["enctype"] != "multipart/form-data" {
			t.Fatalf(`f.Options["enctype"] should be 'multipart/form-data'`)
		}
	})
}

func TestLabel(t *testing.T) {
	t.Run("form label", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		l := f.Label("The Label", tags.Options{})
		if !strings.Contains(string(l.HTML()), `<label class="block text-sm font-medium text-gray-700">The Label</label>`) {
			t.Fatalf("Label should contain 'The Label'")
		}
	})

	t.Run("label with empty value", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		l := f.Label("", tags.Options{})
		if !strings.Contains(string(l.HTML()), `<label class="block text-sm font-medium text-gray-700"></label>`) {
			t.Fatalf("form should contain empty label")
		}
	})

	t.Run("using other tag attributes", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		l := f.Label("something", tags.Options{"id": "label-id"})
		if !strings.Contains(string(l.HTML()), `<label class="block text-sm font-medium text-gray-700" id="label-id">something</label>`) {
			t.Fatalf("form should contain empty label")
		}
	})

	t.Run("passing class to the label", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		l := f.Label("something", tags.Options{"id": "label-id", "class": "text-red-500"})
		if !strings.Contains(string(l.HTML()), `<label class="block text-sm font-medium text-gray-700 text-red-500" id="label-id">something</label>`) {
			t.Fatalf("form should contain mixed classes")
		}
	})
}

func TestFileTag(t *testing.T) {
	t.Run("file tag with no options", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		f.Append(f.FileTag(tags.Options{"name": "file"}))

		if !strings.Contains(string(f.HTML()), `<input name="file" type="file" />`) {
			t.Fatalf("form should contain file input")
		}

		if !strings.Contains(string(f.HTML()), `enctype="multipart/form-data"`) {
			t.Fatalf("enctype should be multipart/form-data")
		}

	})
}

func TestTextArea(t *testing.T) {
	t.Run("text area with no options", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		f.Append(f.TextAreaTag(tags.Options{"name": "text"}))

		if !strings.Contains(string(f.HTML()), `<textarea name="text"></textarea>`) {
			t.Fatalf("form should contain the textarea")
		}

	})

	t.Run("text area with value", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		f.Append(f.TextAreaTag(tags.Options{"name": "text", "value": "some text"}))

		if !strings.Contains(string(f.HTML()), `<textarea name="text">some text</textarea>`) {
			t.Fatalf("textarea should contain `some text`")
		}
	})

	t.Run("text area with encoded content", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		f.Append(f.TextAreaTag(tags.Options{"name": "text", "value": "<script>alert</script>"}))

		if !strings.Contains(string(f.HTML()), `<textarea name="text">&lt;script&gt;alert&lt;/script&gt;</textarea>`) {
			t.Fatalf("textarea should contain the encoded content")
		}
	})

	t.Run("text area with tag_only", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		f.Append(f.TextAreaTag(tags.Options{"name": "text", "tag_only": "true"}))

		if !strings.Contains(string(f.HTML()), `<textarea name="text"></textarea>`) {
			t.Fatalf("textarea should not contain tag_only")
		}
	})

	t.Run("text area", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		f.Append(f.TextArea(tags.Options{"name": "text", "tag_only": "true"}))

		if !strings.Contains(string(f.HTML()), `<textarea name="text"></textarea>`) {
			t.Fatalf("textarea should not contain tag_only")
		}
	})
}

func TestHiddenTag(t *testing.T) {
	t.Run("hidden field with tag_only", func(t *testing.T) {
		f := tailush.NewForm(tags.Options{}, newHctx())
		f.Append(f.HiddenTag(tags.Options{"name": "id", "tag_only": "true", "value": "1"}))

		if !strings.Contains(string(f.HTML()), `<input name="id" type="hidden" value="1" />`) {
			t.Fatalf("textarea should not contain tag_only")
		}
	})
}
