package vecty

import (
	"fmt"
	"testing"
)

var _ = func() bool {
	isTest = true
	return true
}()

// TODO(slimsag): TestCore; Core.Context
// TODO(slimsag): TestComponent; Component.Render; Component.Context
// TODO(slimsag): TestUnmounter; Unmounter.Unmount
// TODO(slimsag): TestComponentOrHTML
// TODO(slimsag): TestRestorer; Restorer.Restore
// TODO(slimsag): TestHTML; HTML.Restore
// TODO(slimsag): TestTag
// TODO(slimsag): TestText
// TODO(slimsag): TestRerender

// TestRenderBody_ExpectsBody tests that RenderBody always expects a "body" tag
// and panics otherwise.
func TestRenderBody_ExpectsBody(t *testing.T) {
	cases := []struct {
		name      string
		render    *HTML
		wantPanic string
	}{
		{
			name:      "text",
			render:    Text("Hello world!"),
			wantPanic: "vecty: RenderBody expected Component.Render to return a body tag, found \"\"", // TODO(slimsag): bug
		},
		{
			name:      "div",
			render:    Tag("div"),
			wantPanic: "vecty: RenderBody expected Component.Render to return a body tag, found \"div\"",
		},
		{
			name:      "body",
			render:    Tag("body"),
			wantPanic: "runtime error: invalid memory address or nil pointer dereference", // TODO(slimsag): relies on js
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var gotPanic string
			func() {
				defer func() {
					r := recover()
					if r != nil {
						gotPanic = fmt.Sprint(r)
					}
				}()
				RenderBody(&componentFunc{render: func() *HTML {
					return c.render
				}})
			}()
			if c.wantPanic != gotPanic {
				t.Fatalf("want panic %q got panic %q", c.wantPanic, gotPanic)
			}
		})
	}
}

// TODO(slimsag): TestRenderBody_Standard
// TODO(slimsag): TestSetTitle
// TODO(slimsag): TestAddStylesheet

type componentFunc struct {
	Core
	render func() *HTML
}

func (c *componentFunc) Render() *HTML { return c.render() }
