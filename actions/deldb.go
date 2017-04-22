package actions

import (
	"strconv"

	"github.com/go-xorm/dbweb/models"
	"github.com/tango-contrib/renders"
)

type Del struct {
	AuthRenderBase
}

func (c *Del) Get() error {
	id, err := strconv.ParseInt(c.Req().FormValue("id"), 10, 64)
	if err != nil {
		return err
	}

	if err := models.DelEngineById(id); err != nil {
		return err
	}

	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("delsuccess.html", renders.T{
		"engines": engines,
		"IsLogin": c.IsLogin(),
	})
}
