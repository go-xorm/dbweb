package actions

import (
	"github.com/Unknwon/i18n"
	"github.com/go-xorm/dbweb/middlewares"
	"github.com/go-xorm/dbweb/models"
	"github.com/tango-contrib/binding"
	"github.com/tango-contrib/flash"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/xsrf"
)

type Addb struct {
	RenderBase

	binding.Binder
	xsrf.Checker
	middlewares.Auther
	flash.Flash
}

func (c *Addb) Get() error {
	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("add.html", renders.T{
		"dbs":          SupportDBs,
		"flash":        c.Flash.Data(),
		"engines":      engines,
		"XsrfFormHtml": c.XsrfFormHtml(),
		"IsLogin":      c.IsLogin(),
		"isAdd":        true,
	})
}

func (c *Addb) Post() {
	var engine models.Engine
	if err := c.MapForm(&engine); err != nil {
		c.Flash.Set("ErrAdd", i18n.Tr(c.CurLang(), "err_param"))
		c.Flash.Redirect("/addb")
		return
	}

	if err := models.AddEngine(&engine); err != nil {
		c.Flash.Set("ErrAdd", i18n.Tr(c.CurLang(), "err_add_failed"))
		c.Flash.Redirect("/addb")
		return
	}

	c.Flash.Redirect("/")
}
