// +build !bindata

// Copyright 2017 The XORM Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/Unknwon/com"
	"github.com/go-xorm/dbweb/modules/setting"
)

func Locale(name string) ([]byte, error) {
	return fileFromDir(path.Join("langs", name))
}

// fileFromDir is a helper to read files from static or custom path.
func fileFromDir(name string) ([]byte, error) {
	staticPath := path.Join(setting.StaticRootPath, "options", name)

	if com.IsFile(staticPath) {
		return ioutil.ReadFile(staticPath)
	}

	return []byte{}, fmt.Errorf("Asset file does not exist: %s", name)
}
