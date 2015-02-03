package actions

import (
	"fmt"
	"strconv"

	"github.com/go-xorm/dbweb/middlewares"
	"github.com/go-xorm/dbweb/models"
	"github.com/go-xorm/xorm"
	"github.com/lunny/tango"
)

type DelRecord struct {
	tango.Req
	tango.Ctx
	middlewares.Auther
}

func (d *DelRecord) Get() error {
	id, err := strconv.ParseInt(d.FormValue("id"), 10, 64)
	if err != nil {
		return err
	}

	engine, err := models.GetEngineById(id)
	if err != nil {
		return err
	}

	o := getOrm(engine.Name)
	if o == nil {
		o, err = xorm.NewEngine(engine.Driver, engine.DataSource)
		if err != nil {
			return err
		}

		setOrm(engine.Name, o)
	}

	tb := d.FormValue("tb")
	colname := d.FormValue("colname")
	colval := d.FormValue("colval")
	isnumeric, _ := strconv.ParseBool(d.FormValue("isnumeric"))

	var val string = colval
	if !isnumeric {
		val = "'" + val + "'"
	}

	_, err = o.Exec(fmt.Sprintf("delete from %s where %s = %s", tb, colname, val))
	if err != nil {
		return err
	}

	d.Redirect(fmt.Sprintf("/view?id=%d&tb=%s", id, tb))

	return nil
}
