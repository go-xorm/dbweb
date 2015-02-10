package actions

import (
	"github.com/Unknwon/i18n"
	"github.com/tango-contrib/flash"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/xsrf"

	"github.com/go-xorm/dbweb/middlewares"
	"github.com/go-xorm/dbweb/models"
)

type ChgPass struct {
	RenderBase

	xsrf.Checker
	middlewares.Auther
	flash.Flash
}

func (c *ChgPass) Get() error {
	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("chgpass.html", renders.T{
		"XsrfFormHtml": c.XsrfFormHtml(),
		"engines":      engines,
		"flash":        c.Flash.Data(),
		"IsLogin":      c.IsLogin(),
	})
}

func (c *ChgPass) Post() {
	oldPass := c.FormValue("old_pass")
	newPass := c.FormValue("new_pass")
	cfmPass := c.FormValue("cfm_pass")

	defer c.Flash.Redirect("/chgpass")

	if newPass != cfmPass {
		c.Flash.Set("cfmError", i18n.Tr(c.CurLang(), "password_not_eq"))
		return
	}

	user := c.Auther.LoginUser()
	if user != nil {
		if models.EncodePassword(oldPass) != user.Password {
			c.Flash.Set("oldError", i18n.Tr(c.CurLang(), "ori_password_not_correct"))
			return
		}
	} else {
		c.Flash.Set("otherError", i18n.Tr(c.CurLang(), "unknown_error"))
		return
	}

	user.Password = newPass
	err := models.UpdateUser(user)
	if err != nil {
		c.Flash.Set("otherError", err.Error())
		return
	}

	c.Flash.Set("changeSuccess", i18n.Tr(c.CurLang(), "password_change_success"))
}
