package handlers

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

//go:embed dist
var distFS embed.FS

func RegisterSPA(e *echo.Echo) error {
	root, _ := fs.Sub(distFS, "dist")
	fileServer := http.FileServer(http.FS(root))

	// files are served if exist, otherwise serve index.html (SPA)
	e.GET("/*", func(c echo.Context) error {
		path := c.Param("*")
		if path == "" || !existsInFS(root, path) {
			data, _ := fs.ReadFile(root, "index.html")
			return c.Blob(http.StatusOK, "text/html; charset=utf-8", data)
		}
		// let http.FileServer serve the requested file
		echo.WrapHandler(http.StripPrefix("/", fileServer))(c)
		return nil
	})
	return nil
}

func existsInFS(root fs.FS, p string) bool {
	p = strings.TrimPrefix(p, "/")
	_, err := fs.Stat(root, p)
	return err == nil
}
