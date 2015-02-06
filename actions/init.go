package actions

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

var (
	ormCache  = make(map[string]*xorm.Engine)
	cacheLock sync.RWMutex

	SupportDBs = []string{
		"mysql",
		"postgres",
	}
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
