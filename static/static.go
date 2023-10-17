package static

import (
	"carot-admin/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterWebStatic register web static assets route
func RegisterWebStatic(e *gin.Engine) {
	routeWebStatic(e, "/", "/index.html", "/favicon.ico", "/logo.png", "/sw.js", "/manifest.json", "/assets/*filepath")
}

func routeWebStatic(e *gin.Engine, paths ...string) {
	staticHandler := http.FileServer(web.NewFileSystem())
	handler := func(c *gin.Context) {
		staticHandler.ServeHTTP(c.Writer, c.Request)
	}
	for _, path := range paths {
		e.GET(path, handler)
		e.HEAD(path, handler)
	}
}
