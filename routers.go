package main

import (
	"html/template"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/Unknwon/i18n"
	"github.com/go-xorm/xorm"
	"github.com/lunny/nodb"
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

func InitTango(isDebug bool) *tango.Tango {
	t := tango.New()
	if isDebug {
		t.Use(debug.Debug(debug.Options{
			HideResponseBody: true,
			IgnorePrefix:     "/static",
		}))
	}
	t.Use(tango.ClassicHandlers...)
	sess := session.New(session.Options{
		MaxAge: sessionTimeout,
	})
	t.Use(
		binding.Bind(),
		tango.Static(tango.StaticOptions{
			RootPath: filepath.Join(*homeDir, "static"),
			Prefix:   "static",
		}),
		renders.New(renders.Options{
			Reload:    true,
			Directory: filepath.Join(*homeDir, "templates"),
			Funcs: template.FuncMap{
				"isempty": func(s string) bool {
					return len(s) == 0
				},
				"add": func(a, b int) int {
					return a + b
				},
				"isNil": isNil,
				"i18n":  i18n.Tr,
			},
			Vars: renders.T{
				"GoVer":    strings.Trim(runtime.Version(), "go"),
				"TangoVer": tango.Version(),
				"XormVer":  xorm.Version,
				"NodbVer":  nodb.Version,
			},
		}),
		middlewares.Auth("/login", sess),
		flash.Flashes(sess),
		sess,
	)

	t.Any("/", new(actions.Home))
	t.Any("/login", new(actions.Login))
	t.Any("/logout", new(actions.Logout))
	t.Any("/addb", new(actions.Addb))
	t.Any("/view", new(actions.View))
	t.Any("/del", new(actions.Del))
	t.Any("/delRecord", new(actions.DelRecord))
	t.Any("/chgpass", new(actions.ChgPass))
	return t
}
