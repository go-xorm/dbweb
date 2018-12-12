package actions

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/go-xorm/core"
	"github.com/go-xorm/dbweb/models"
	"github.com/tango-contrib/renders"
)

type View struct {
	AuthRenderBase
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

	// Sort tables by name.
	sort.Slice(tables, func(i, j int) bool {
		return tables[i].Name < tables[j].Name
	})

	var records = make([][]*string, 0)
	var columns = make([]*core.Column, 0)
	tb := c.Req().FormValue("tb")
	tb = strings.Replace(tb, `"`, "", -1)
	tb = strings.Replace(tb, `'`, "", -1)
	tb = strings.Replace(tb, "`", "", -1)

	var isTableView = len(tb) > 0

	sql := c.Req().FormValue("sql")
	var table *core.Table
	var pkIdx int
	var isExecute bool
	var hasRowNum bool
	var affected int64
	var total int
	var countSql string
	var execSql string
	var args = make([]interface{}, 0)

	start, _ := strconv.Atoi(c.Req().FormValue("start"))
	limit, _ := strconv.Atoi(c.Req().FormValue("limit"))
	if limit == 0 {
		limit = 20
	}
	if sql != "" || tb != "" {
		hasRowNum = false
		if sql != "" {
			isExecute = !strings.HasPrefix(strings.ToLower(sql), "select")
			execSql = sql
		} else if tb != "" {
			for _, tt := range tables {
				if tb == tt.Name {
					table = tt
					break
				}
			}
			countSql = "select count(*) from " + tb
			if engine.Driver == "mssql" {
				orderBy := ""
				pkCols := table.PKColumns()
				for _, pk := range pkCols {
					if len(orderBy) > 0 {
						orderBy += ", " + pk.Name
					} else {
						orderBy = pk.Name
					}
				}
				hasRowNum = true
				sql = "SELECT * FROM "+tb
				execSql = fmt.Sprintf("SELECT TOP %d * FROM (SELECT ROW_NUMBER() OVER(ORDER BY "+orderBy+") AS RowNumber"+
					", * FROM "+tb+") AS Res WHERE RowNumber > %d", limit, start)
			} else {
				execSql = fmt.Sprintf("select * from "+tb+" LIMIT %d OFFSET %d", limit, start)
			}
		} else {
			return errors.New("unknow operation")
		}

		if isExecute {
			res, err := o.Exec(execSql)
			if err != nil {
				return err
			}
			affected, _ = res.RowsAffected()
		} else {
			if len(countSql) > 0 {
				if err = o.DB().QueryRow(countSql).Scan(&total); err != nil {
					return err
				}
			}

			rows, err := o.DB().Query(execSql, args...)
			if err != nil {
				return err
			}
			defer rows.Close()

			cols, err := rows.Columns()
			if err != nil {
				return err
			}

			if table != nil {
				for i, col := range cols {
					c := table.GetColumn(col)
					if c != nil {
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
				datas := make([]*string, len(cols))
				err = rows.ScanSlice(&datas)
				if err != nil {
					return err
				}
				if hasRowNum {
					datas = datas[1:]
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
		"engines":     engines,
		"tables":      tables,
		"table":       table,
		"records":     records,
		"columns":     columns,
		"id":          id,
		"sql":         sql,
		"tb":          tb,
		"isExecute":   isExecute,
		"isTableView": isTableView,
		"limit":       limit,
		"curPage":     start / limit,
		"totalPage":   pager(total, limit),
		"affected":    affected,
		"pkIdx":       pkIdx,
		"curEngine":   engine.Name,
		"IsLogin":     c.IsLogin(),
	})
}

func pager(total, limit int) int {
	if total%limit == 0 {
		return total / limit
	}
	return total/limit + 1
}

func curPage(start, limit int) int {
	return start / limit
}
