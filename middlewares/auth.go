package middlewares

import (
	"github.com/lunny/tango"
	"github.com/tango-contrib/session"

	"github.com/go-xorm/dbweb/models"
)

var (
	LoginIdKey = "auth_user_id"
)

type auther interface {
	AskAuth() bool
	SetUserId(int64)
	SetSession(*session.Session)
}

type Auther struct {
	id int64
	s  *session.Session
}

func (Auther) AskAuth() bool {
	return true
}

func (a *Auther) SetSession(s *session.Session) {
	a.s = s
}

func (a *Auther) SetUserId(id int64) {
	a.id = id
}

func (a *Auther) SetLogin(id int64) {
	a.SetUserId(id)
	a.s.Set(LoginIdKey, id)
}

func (a *Auther) Logout() {
	a.s.Del(LoginIdKey)
	a.s.Release()
}

func (a *Auther) LoginUserId() int64 {
	return a.id
}

func (a *Auther) IsLogin() bool {
	return a.id > 0
}

func (a *Auther) LoginUser() *models.User {
	user, err := models.GetUserById(a.id)
	if err != nil {
		return nil
	}
	return user
}

type AuthUser struct {
	Auther
}

func (AuthUser) AskAuth() bool {
	return false
}

func Auth(redirct string, sessions *session.Sessions) tango.HandlerFunc {
	return func(ctx *tango.Context) {
		if auther, ok := ctx.Action().(auther); ok {
			s := sessions.Session(ctx.Req(), ctx.ResponseWriter)
			auther.SetSession(s)
			if userId := s.Get(LoginIdKey); userId == nil {
				if auther.AskAuth() {
					ctx.Redirect(redirct)
					return
				}
			} else {
				auther.SetUserId(userId.(int64))
			}
		}
		ctx.Next()
	}
}
