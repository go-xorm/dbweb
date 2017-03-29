// +build bindata

// Copyright 2017 The XORM Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

import (
	"path"
)

// Locale reads the content of a specific locale from bindata or custom path.
func Locale(name string) ([]byte, error) {
	return fileFromDir(path.Join("langs", name))
}

// fileFromDir is a helper to read files from bindata or custom path.
func fileFromDir(name string) ([]byte, error) {
	return Asset(name)
}
