package actions

import (
	"github.com/tango-contrib/flash"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/xsrf"

	"github.com/go-xorm/dbweb/middlewares"
	"github.com/go-xorm/dbweb/models"
)

type Login struct {
	RenderBase

	middlewares.AuthUser
	xsrf.Checker
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
	name := c.Request.FormValue("user")
	password := c.Request.FormValue("password")

	user, err := models.GetUserByName(name)
	if err != nil {
		c.Flash.Set("user", name)
		c.Flash.Set("AuthError", "账号或密码错误")
		c.Redirect("/login")
		return
	}

	if user.Password != models.EncodePassword(password) {
		c.Flash.Set("user", name)
		c.Flash.Set("AuthError", "账号或密码错误")
		c.Redirect("/login")
	} else {
		c.SetLogin(user.Id)
		c.Redirect("/")
	}
}
