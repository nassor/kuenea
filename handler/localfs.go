package handler

import (
	"net/http"
	"strings"
)

// Handle local files (without filesystem directory)
func LocalFSServer(rootPath string, localPath string) http.Handler {
	return http.StripPrefix(localPath, removeDirListing(http.FileServer(http.Dir(rootPath))))
}

// 404 for directory listing
func removeDirListing(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
