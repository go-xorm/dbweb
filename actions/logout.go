package actions

import (
	"github.com/go-xorm/dbweb/middlewares"
	"github.com/lunny/tango"
)

type Logout struct {
	RenderBase
	middlewares.AuthUser
	tango.Ctx
}

func (l *Logout) Get() {
	if l.IsLogin() {
		l.Logout()
	}
	l.Redirect("/")
}
