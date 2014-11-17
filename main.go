package main

import (
	"fmt"
	"sync"

	"github.com/go-xorm/xorm"
	"github.com/go-xweb/xweb"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-xorm/ql"
	_ "github.com/lib/pq"
	_ "github.com/lunny/ql/driver"
)

var (
	orm *xorm.Engine
)

var (
	ormCache  = make(map[string]*xorm.Engine)
	cacheLock sync.RWMutex
)

func getOrm(name string) *xorm.Engine {
	cacheLock.RLock()
	defer cacheLock.RUnlock()
	if o, ok := ormCache[name]; ok {
		return o
	}
	return nil
}

func setOrm(name string, o *xorm.Engine) {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	ormCache[name] = o
}

func main() {
	var err error
	orm, err = xorm.NewEngine("ql", "./xorm.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = orm.Sync(&Engine{})
	if err != nil {
		fmt.Println(err)
		return
	}

	xweb.AddAction(&MainAction{})
	xweb.Run(":8989")
}
