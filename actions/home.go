package actions

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/dbweb/middlewares"
	"github.com/go-xorm/dbweb/models"
	"github.com/tango-contrib/renders"
)

type Home struct {
	RenderBase

	middlewares.Auther
}

func (c *Home) Get() error {
	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("root.html", renders.T{
		"engines": engines,
		"tables":  []core.Table{},
		"records": [][]string{},
		"columns": []string{},
		"id":      0,
		"ishome":  true,
		"IsLogin": c.IsLogin(),
	})
}
