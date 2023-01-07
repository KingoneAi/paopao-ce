// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build docs
// +build docs

package routers

import (
	"github.com/alimy/cfg"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/docs/openapi"
)

// registerDocs register docs asset route
func registerDocs(e *gin.Engine) {
	cfg.Be("Docs:OpenAPI", func() {
		e.StaticFS("/docs/openapi", openapi.NewFileSystem())
	})
}
