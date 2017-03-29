// Copyright 2017 The XORM Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package public

//go:generate go-bindata -tags "bindata" -ignore "\\.go|\\.less" -pkg "public" -o "bindata.go" ../../public/...
//go:generate go fmt bindata.go
//go:generate rm -f bindata.go.bak
