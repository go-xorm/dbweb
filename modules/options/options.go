// Copyright 2016 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package options

//go:generate go-bindata -tags "bindata" -ignore "TRANSLATORS" -pkg "options" -o "bindata.go" ../../options/...
//go:generate go fmt bindata.go
//go:generate sed -i.bak s/..\/..\/options\/// bindata.go
//go:generate rm -f bindata.go.bak
