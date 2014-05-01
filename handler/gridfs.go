package handler

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"kuenea/conf"
	"labix.org/v2/mgo"
)

type gridFSHandler struct {
	gridFS   *conf.GridFSConfig
	PathFile string
	Session  *mgo.Session
}

func (g *gridFSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var file *mgo.GridFile
	var err error

	path := strings.Replace(r.URL.Path[1:], g.PathFile, "", 1)

	file, err = g.Session.DB("").GridFS("fs").Open(path)
	if err != nil {
		g.Session.Refresh()
		time.Sleep(1 * time.Second)
		file, err = g.Session.DB("").GridFS("fs").Open(path)
		if err != nil {
			fmt.Printf(err.Error())
			http.NotFound(w, r)
			return
		}
	}

	if g.gridFS.ReadSeeker {
		http.ServeContent(w, r, file.ContentType(), file.UploadDate(), file)
	} else {

		if t, err := time.Parse(http.TimeFormat, r.Header.Get("If-Modified-Since")); err == nil && file.UploadDate().Before(t.Add(2*time.Second)) {
			delete(w.Header(), "Content-Type")
			delete(w.Header(), "Content-Length")
			w.WriteHeader(http.StatusNotModified)
			return
		}

		if file.ContentType() == "" {
			contentType := mime.TypeByExtension(filepath.Ext(file.Name()))
			w.Header().Set("Content-Type", contentType)
		} else {
			w.Header().Set("Content-Type", file.ContentType())
		}

		w.Header().Set("Last-Modified", file.UploadDate().UTC().Format(http.TimeFormat))

		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "gridfs read error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		file.Close()
	}
}

// Handle server requests, find file and response.
func GridFSServer(gridFS conf.GridFSConfig) http.Handler {

	session, err := mgo.Dial(gridFS.ConnectURI)
	if err != nil {
		log.Fatalf("Could not conected to database: %v", err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	log.Printf("MongoDB: %v -> %v", session.LiveServers(), gridFS.Path)

	return &gridFSHandler{&gridFS, gridFS.Path, session}
}
