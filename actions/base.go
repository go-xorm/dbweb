package actions

import (
	"strings"

	"github.com/Unknwon/i18n"
	"github.com/lunny/tango"

	"github.com/tango-contrib/renders"
)

func formatLang(l string) string {
	if len(l) != 5 || l[2] != '-' {
		return "en-US"
	}
	return strings.ToLower(l[:2]) + "-" + strings.ToUpper(l[3:])
}

type Base struct {
	//tango.Compress
}

type RenderBase struct {
	Base
	renders.Renderer
	tango.Ctx
}

func (r *RenderBase) CurLang() string {
	al := r.Req().Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5] // Only compare first 5 letters.
		if i18n.IsExist(formatLang(al)) {
			return formatLang(al)
		}
	}
	return "en-US"
}

func (r *RenderBase) Render(tmpl string, t ...renders.T) error {
	if len(t) > 0 {
		return r.Renderer.Render(tmpl, t[0].Merge(renders.T{
			"Lang": r.CurLang(),
		}))
	}
	return r.Renderer.Render(tmpl, renders.T{
		"Lang": r.CurLang(),
	})
}
