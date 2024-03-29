package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:dist/*
var files embed.FS

// NewFileSystem get an embed static assets http.FileSystem instance.
func NewFileSystem() http.FileSystem {
	subfs, _ := fs.Sub(files, "dist")
	return http.FS(subfs)
}
