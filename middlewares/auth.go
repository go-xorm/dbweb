package middlewares

import (
	"github.com/lunny/tango"
	"github.com/tango-contrib/session"
)

var (
	// LoginIDKey the login user id stored in Session
	LoginIDKey = "auth_user_id"
)

// Auther all actions implmented this interface will be Auth middleware check
type auther interface {
	AskAuth() bool
	SetLoginUserID(id int64)
}

type Auther struct {
	id int64
}

func (Auther) AskAuth() bool {
	return true
}

func (a *Auther) SetLoginUserID(id int64) {
	a.id = id
}

func (a *Auther) LoginUserID() int64 {
	return a.id
}

func (a *Auther) IsLogin() bool {
	return a.id > 0
}

type AuthUser struct {
	Auther
}

func (AuthUser) AskAuth() bool {
	return false
}

// Auth middleware will check the action needs to check, if yes, then get user id from session to check if exist
func Auth(redirct string) tango.HandlerFunc {
	return func(ctx *tango.Context) {
		if auther, ok := ctx.Action().(auther); ok {
			if sess, ok := ctx.Action().(session.Sessioner); ok {
				// get session key to check is logined in
				if userID := sess.GetSession().Get(LoginIDKey); userID != nil {
					auther.SetLoginUserID(userID.(int64))
					ctx.Next()
					return
				}
			}
			// if ask login
			if auther.AskAuth() {
				ctx.Debug("No session found or no user id found in session")
				ctx.Redirect(redirct)
				return
			}
		}
		ctx.Next()
	}
}
