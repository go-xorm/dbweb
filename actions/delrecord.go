package actions

import (
	"fmt"
	"strconv"

	"github.com/go-xorm/dbweb/models"
	"github.com/go-xorm/xorm"
)

type DelRecord struct {
	AuthRenderBase
}

func (d *DelRecord) Get() error {
	id, err := strconv.ParseInt(d.Req().FormValue("id"), 10, 64)
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

	tb := d.Req().FormValue("tb")
	colname := d.Req().FormValue("colname")
	colval := d.Req().FormValue("colval")
	isnumeric, _ := strconv.ParseBool(d.Req().FormValue("isnumeric"))

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
