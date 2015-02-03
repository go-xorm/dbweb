package actions

import (
	"github.com/go-xorm/dbweb/middlewares"
	"github.com/lunny/tango"
	"github.com/tango-contrib/flash"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/xsrf"
)

type Login struct {
	middlewares.AuthUser
	renders.Renderer
	xsrf.Checker
	tango.Req
	flash.Flash
}

func (c *Login) Get() error {
	if c.IsLogin() {
		c.Redirect("/")
		return nil
	}

	return c.Render("login.html", renders.T{
		"XsrfFormHtml": c.XsrfFormHtml(),
		"flash":        c.Flash.Data(),
	})
}

func (c *Login) Post() {
	c.Request.ParseForm()
	user := c.Request.FormValue("user")
	password := c.Request.FormValue("password")

	if user != "admin" || password != "admin" {
		c.Flash.Set("AuthError", "账号或密码错误")
		c.Redirect("/login")
	} else {
		c.SetLogin(1)
		c.Redirect("/")
	}
}
