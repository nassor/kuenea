package handler

import (
	"log"
	"net/http"
	"strings"

	"kuenea/conf"
)

// Handle local files (without filesystem directory)
func LocalFSServer(localConf conf.LocalFSConfig) http.Handler {
	log.Printf("LocalFS: %v -> %v", localConf.Root, localConf.Path)
	return http.StripPrefix(localConf.Path, removeDirListing(http.FileServer(http.Dir(localConf.Root))))
}

// 404 for directory listing
func removeDirListing(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || r.URL.Path == "" {
			http.NotFound(w, r)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
