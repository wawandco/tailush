package tailush_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gobuffalo/tags/v3"
	"github.com/wawandco/tailush"
)

func TestFormForInputTag(t *testing.T) {
	t.Run("type hidden", func(t *testing.T) {
		var model struct{}
		f := tailush.NewFormFor(model, tags.Options{}, newHctx())

		it := f.InputTag("", tags.Options{
			"type": "hidden",
		})

		if strings.Contains(string(it.HTML()), `<label class="block text-sm font-medium text-gray-700 mb-1">Name</label>`) {
			t.Fatalf("form shouldn't contain label")
		}

		if !strings.Contains(string(it.HTML()), `<input id="-" name="" tags-field="" type="hidden" value="" />`) {
			t.Fatalf("form should contain hidden input")
		}
	})
}
