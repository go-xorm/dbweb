package main

import (
	"html/template"
	"reflect"
	"time"

	"github.com/lunny/tango"
	"github.com/tango-contrib/binding"
	"github.com/tango-contrib/debug"
	"github.com/tango-contrib/flash"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/session"

	"github.com/go-xorm/dbweb/actions"
	"github.com/go-xorm/dbweb/middlewares"
)

var (
	sessionTimeout = time.Minute * 20
)

func isNil(a interface{}) bool {
	if a == nil {
		return true
	}
	aa := reflect.ValueOf(a)
	return !aa.IsValid() || (aa.Type().Kind() == reflect.Ptr && aa.IsNil())
}

func InitTango() *tango.Tango {
	t := tango.New()
	t.Use(debug.Debug(debug.Options{
		HideResponseBody: true,
		IgnorePrefix:     "/static",
	}))
	t.Use(tango.ClassicHandlers...)
	t.Use(binding.Bind())
	t.Use(tango.Static(tango.StaticOptions{
		RootPath: "./static",
		Prefix:   "static",
	}))
	t.Use(renders.New(renders.Options{
		Reload: t.Mode == tango.Dev,
		Funcs: template.FuncMap{
			"isempty": func(s string) bool {
				return len(s) == 0
			},
			"add": func(a, b int) int {
				return a + b
			},
			"isNil": isNil,
		},
	}))
	s := session.New(sessionTimeout)
	t.Use(middlewares.Auth("/login", s))
	t.Use(flash.Flashes(flash.Options{
		CookiePath: "/",
	}))
	t.Use(s)

	t.Any("/", new(actions.Home))
	t.Any("/login", new(actions.Login))
	t.Any("/logout", new(actions.Logout))
	t.Any("/addb", new(actions.Addb))
	t.Any("/view", new(actions.View))
	t.Any("/del", new(actions.Del))
	t.Any("/delRecord", new(actions.DelRecord))
	return t
}
