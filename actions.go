package main

import (
	"errors"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/go-xweb/xweb"
)

type Engine struct {
	Id         int64
	Name       string `xorm:"unique"`
	Driver     string
	DataSource string
}

type MainAction struct {
	*xweb.Action

	home xweb.Mapper `xweb:"/"`
	addb xweb.Mapper
	del  xweb.Mapper
	view xweb.Mapper
}

func (c *MainAction) Home() error {
	engines := make([]Engine, 0)
	err := orm.Find(&engines)
	if err != nil {
		return err
	}
	return c.Render("root.html", &xweb.T{
		"engines": engines,
		"tables":  []core.Table{},
		"records": [][]string{},
		"columns": []string{},
		"id":      0,
	})
}

func (c *MainAction) Addb() error {
	if c.Method() == "GET" {
		return c.Render("add.html")
	}

	var engine Engine
	err := c.MapForm(&engine)
	if err != nil {
		return err
	}
	_, err = orm.Insert(&engine)
	if err != nil {
		return err
	}
	return c.Render("addsuccess.html")
}

func (c *MainAction) Del() error {
	id, err := c.GetInt("id")
	if err != nil {
		return err
	}

	_, err = orm.Id(id).Delete(new(Engine))
	if err != nil {
		return err
	}

	return c.Render("delsuccess.html")
}

func (c *MainAction) View() error {
	id, err := c.GetInt("id")
	if err != nil {
		return err
	}

	engine := new(Engine)
	has, err := orm.Id(id).Get(engine)
	if err != nil {
		return err
	}

	if !has {
		return errors.New("db is not exist")
	}

	o := getOrm(engine.Name)

	if o == nil {
		o, err = xorm.NewEngine(engine.Driver, engine.DataSource)
		if err != nil {
			return err
		}

		setOrm(engine.Name, o)
	}

	tables, err := o.DBMetas()
	if err != nil {
		return err
	}

	engines := make([]Engine, 0)
	err = orm.Find(&engines)
	if err != nil {
		return err
	}

	var records = make([][]*string, 0)
	var columns = make([]string, 0)
	tb := c.GetString("tb")
	sql := c.GetString("sql")
	if sql != "" || tb != "" {
		if sql != "" {
		} else if tb != "" {
			sql = "select * from " + tb
		}

		rows, err := o.DB().Query(sql)
		if err != nil {
			return err
		}
		defer rows.Close()

		columns, err = rows.Columns()
		if err != nil {
			return err
		}

		for rows.Next() {
			datas := make([]*string, len(columns))
			err = rows.ScanSlice(&datas)
			if err != nil {
				return err
			}
			records = append(records, datas)
		}
	}

	return c.Render("root.html", &xweb.T{
		"engines": engines,
		"tables":  tables,
		"records": records,
		"columns": columns,
		"id":      id,
	})
}
