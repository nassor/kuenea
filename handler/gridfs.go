package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"kuenea/conf"
	"labix.org/v2/mgo"
)

type gridFSHandler struct {
	mdbConf  *conf.DatabaseConfig
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
	w.Header().Set("Content-Type", file.ContentType())
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "gridfs read error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	file.Close()
}

// Handle server requests, find file and response.
func GridFSServer(mdbConf *conf.DatabaseConfig, pathFile string) http.Handler {
	session, err := mgo.Dial(mdbConf.ConnectURI)
	if err != nil {
		log.Fatalf("Could not conected to database: %v", err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	log.Printf("MongoDB: %v -> %v", session.LiveServers(), mdbConf.Path)

	return &gridFSHandler{mdbConf, pathFile, session}
}
