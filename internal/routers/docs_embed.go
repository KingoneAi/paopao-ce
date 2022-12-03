//go:build docs
// +build docs

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/docs/openapi"
	"github.com/rocboss/paopao-ce/pkg/cfg"
)

// registerDocs register docs asset route
func registerDocs(e *gin.Engine) {
	cfg.Be("Docs:OpenAPI", func() {
		e.StaticFS("/docs/openapi", openapi.NewFileSystem())
	})
}
