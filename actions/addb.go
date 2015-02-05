package actions

import (
	"errors"

	"github.com/go-xorm/dbweb/middlewares"
	"github.com/go-xorm/dbweb/models"
	"github.com/tango-contrib/binding"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/xsrf"
)

type Addb struct {
	RenderBase

	binding.Binder
	xsrf.Checker
	middlewares.Auther
}

func (c *Addb) Get() error {
	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("add.html", renders.T{
		"engines":      engines,
		"XsrfFormHtml": c.XsrfFormHtml(),
		"IsLogin":      c.IsLogin(),
	})
}

func (c *Addb) Post() error {
	var engine models.Engine
	if err := c.MapForm(&engine); err != nil {
		return errors.New("")
	}

	if err := models.AddEngine(&engine); err != nil {
		return err
	}

	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("addsuccess.html", renders.T{
		"engines": engines,
		"IsLogin": c.IsLogin(),
	})
}
