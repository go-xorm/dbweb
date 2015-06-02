package actions

import (
	"errors"

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
	dataSource := t.Req().FormValue("data_source")
	o := GetOrm(&models.Engine{
		Name:name, 
		Driver:driver,
		DataSource:dataSource,
	})
	if o == nil {
		return errors.New("driver failed")
	}
	err := o.Ping()
	if err !=nil {
		return err
	}
	return map[string]interface{}{
		"status": "ok",
	}
}
