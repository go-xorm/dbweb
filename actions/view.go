package actions

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-xorm/core"
	"github.com/go-xorm/dbweb/middlewares"
	"github.com/go-xorm/dbweb/models"
	"github.com/tango-contrib/renders"
)

type View struct {
	RenderBase

	middlewares.Auther
}

func (c *View) Get() error {
	id, err := strconv.ParseInt(c.Req().FormValue("id"), 10, 64)
	if err != nil {
		return err
	}

	engine, err := models.GetEngineById(id)
	if err != nil {
		return err
	}

	o := GetOrm(engine)
	if o == nil {
		return fmt.Errorf("get engine %s failed", engine.Name)
	}

	tables, err := o.DBMetas()
	if err != nil {
		return err
	}

	var records = make([][]*string, 0)
	var columns = make([]*core.Column, 0)
	tb := c.Req().FormValue("tb")
	sql := c.Req().FormValue("sql")
	var table *core.Table
	var pkIdx int
	var isExecute bool
	var affected int64
	if sql != "" || tb != "" {
		if sql != "" {
			isExecute = !strings.HasPrefix(strings.ToLower(sql), "select")
		} else if tb != "" {
			sql = "select * from " + tb
		} else {
			return errors.New("unknow operation")
		}

		if isExecute {
			res, err := o.Exec(sql)
			if err != nil {
				return err
			}
			affected, _ = res.RowsAffected()
		} else {
			rows, err := o.DB().Query(sql)
			if err != nil {
				return err
			}
			defer rows.Close()

			cols, err := rows.Columns()
			if err != nil {
				return err
			}

			if len(tb) > 0 {
				for _, tt := range tables {
					if tb == tt.Name {
						table = tt
						break
					}
				}
				if table != nil {
					for i, col := range cols {
						c := table.GetColumn(col)
						if len(table.PKColumns()) == 1 && c.IsPrimaryKey {
							pkIdx = i
						}
						columns = append(columns, c)
					}
				}
			} else {
				for _, col := range cols {
					columns = append(columns, &core.Column{
						Name: col,
					})
				}
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
	}

	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("root.html", renders.T{
		"engines":   engines,
		"tables":    tables,
		"table":     table,
		"records":   records,
		"columns":   columns,
		"id":        id,
		"sql":       sql,
		"tb":        tb,
		"isExecute": isExecute,
		"affected":  affected,
		"pkIdx":     pkIdx,
		"curEngine": engine.Name,
		"IsLogin":   c.IsLogin(),
	})
}
