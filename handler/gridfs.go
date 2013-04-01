package handler

import (
	"io"
	"net/http"
	"strings"

	. "kuenea/conf"
	"labix.org/v2/mgo"
)

type gridFSHandler struct {
	GFS  *mgo.GridFS
	Path string
}

func (g *gridFSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, err := g.GFS.Open(strings.Replace(r.URL.Path[1:], g.Path, "", 1))
	if err != nil {
		w.WriteHeader(500)
		return
	}

	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", file.ContentType())
	file.Close()

}

// Handle server requests, find file and response.
func GridFSServer(gfs *mgo.GridFS, config *Config) http.Handler {
	return &gridFSHandler{gfs, config.Database.Path}
}
