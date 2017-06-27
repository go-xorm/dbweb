// +build sqlite3

package actions

import (
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	SupportDBs = append(SupportDBs, DB{"sqlite3", 0})
}
