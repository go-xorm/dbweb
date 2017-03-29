// +build !bindata

// Copyright 2017 The XORM Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package templates

import "net/http"

// FileSystem implements the macaron handler for serving the templates.
func FileSystem(templatesDir string) http.FileSystem {
	return http.Dir(templatesDir)
}
