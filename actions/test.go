package actions

import (
	"errors"
	"fmt"

	"github.com/lunny/tango"

	"github.com/go-xorm/dbweb/middlewares"
	"github.com/go-xorm/dbweb/models"
)

type Test struct {
	RenderBase
	middlewares.Auther
	tango.Json
}

func (t *Test) Get() interface{} {
	name := t.Req().FormValue("name")
	driver := t.Req().FormValue("driver")

	host := t.Form("host")
	port := t.Form("port")
	dbname := t.Form("dbname")
	username := t.Form("username")
	passwd := t.Form("passwd")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, passwd, host, port, dbname)

	//dataSource := t.Req().FormValue("data_source")
	o := GetOrm(&models.Engine{
		Name:       name,
		Driver:     driver,
		DataSource: dataSource,
	})
	if o == nil {
		return errors.New("driver failed")
	}
	err := o.Ping()
	if err != nil {
		return err
	}
	return map[string]interface{}{
		"status": "ok",
	}
}
