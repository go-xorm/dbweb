package actions

import (
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
		c.Flash.Set("cfmError", "两次输入密码不一致")
		return
	}

	user := c.Auther.LoginUser()
	if user != nil {
		if models.EncodePassword(oldPass) != user.Password {
			c.Flash.Set("oldError", "原密码不正确")
			return
		}
	} else {
		c.Flash.Set("otherError", "未知错误")
		return
	}

	user.Password = newPass
	err := models.UpdateUser(user)
	if err != nil {
		c.Flash.Set("otherError", err.Error())
		return
	}

	c.Flash.Set("changeSuccess", "密码修改成功")
}
