package handler

import (
	"io"
	"net/http"
	"strings"

	"labix.org/v2/mgo"
)

type gridFSHandler struct {
	GFS      *mgo.GridFS
	PathFile string
}

func (g *gridFSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, err := g.GFS.Open(strings.Replace(r.URL.Path[1:], g.PathFile, "", 1))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer file.Close()

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "gridfs read error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", file.ContentType())
}

// Handle server requests, find file and response.
func GridFSServer(gfs *mgo.GridFS, pathFile string) http.Handler {
	return &gridFSHandler{gfs, pathFile}
}
