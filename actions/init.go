package actions

import (
	"sync"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/dbweb/models"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

type DB struct {
	Name        string
	DefaultPort int
}

var (
	ormCache  = make(map[string]*xorm.Engine)
	cacheLock sync.RWMutex

	SupportDBs = []DB{
		{"mysql", 3306},
		{"postgres", 5432},
		{"mssql", 1433},
	}
)

func GetOrm(engine *models.Engine) *xorm.Engine {
	cacheLock.Lock()
	defer cacheLock.Unlock()

	o := getOrm(engine.Name)
	if o == nil {
		var err error
		o, err = xorm.NewEngine(engine.Driver, engine.DataSource)
		if err != nil {
			return nil
		}

		setOrm(engine.Name, o)
	}
	return o
}

func getOrm(name string) *xorm.Engine {
	if o, ok := ormCache[name]; ok {
		return o
	}
	return nil
}

func setOrm(name string, o *xorm.Engine) {
	ormCache[name] = o
}
