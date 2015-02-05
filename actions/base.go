package actions

import (
	"github.com/Unknwon/i18n"
	"github.com/lunny/tango"

	"github.com/tango-contrib/renders"
)

type Base struct {
	tango.Compress
}

type RenderBase struct {
	Base
	renders.Renderer
	tango.Req
}

func (r *RenderBase) Render(tmpl string, t ...renders.T) error {
	if len(t) > 0 {
		return r.Renderer.Render(tmpl, t[0].Merge(renders.T{
			"Lang": func() string {
				al := r.Header.Get("Accept-Language")
				if len(al) > 4 {
					al = al[:5] // Only compare first 5 letters.
					if i18n.IsExist(al) {
						return al
					}
				}
				return "en-US"
			}(),
		}))
	}
	return r.Renderer.Render(tmpl)
}
