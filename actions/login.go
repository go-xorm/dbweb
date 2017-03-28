package actions

import (
	"github.com/Unknwon/i18n"
	"github.com/tango-contrib/captcha"
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
	captcha.Captcha
}

func (c *Login) Get() error {
	if c.IsLogin() {
		c.Redirect("/")
		return nil
	}

	return c.Render("login.html", renders.T{
		"XsrfFormHtml": c.XsrfFormHtml(),
		"flash":        c.Flash.Data(),
		"captcha":      c.CreateHtml(),
	})
}

func (c *Login) Post() {
	c.Req().ParseForm()
	name := c.Req().FormValue("user")
	password := c.Req().FormValue("password")

	if !c.Captcha.Verify() {
		c.Flash.Set("user", name)
		c.Flash.Set("AuthError", i18n.Tr(c.CurLang(), "captcha_error"))
		c.Redirect("/login")
		return
	}

	user, err := models.GetUserByName(name)
	if err != nil {
		c.Flash.Set("user", name)
		c.Flash.Set("AuthError", i18n.Tr(c.CurLang(), "pasword_error"))
		c.Redirect("/login")
		return
	}

	if user.Password != models.EncodePassword(password) {
		c.Flash.Set("user", name)
		c.Flash.Set("AuthError", i18n.Tr(c.CurLang(), "pasword_error"))
		c.Redirect("/login")
	} else {
		c.SetLogin(user.Id)
		c.Redirect("/")
	}
}
