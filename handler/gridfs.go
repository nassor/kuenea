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
	dbConf   *conf.DatabaseConfig
	PathFile string
	Session  *mgo.Session
}

func (g *gridFSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var file *mgo.GridFile
	var err error

	path := strings.Replace(r.URL.Path[1:], g.PathFile, "", 1)

	file, err = g.Session.DB(g.dbConf.DBName).GridFS("fs").Open(path)
	if err != nil {
		g.Session.Refresh()
		time.Sleep(1 * time.Second)
		file, err = g.Session.DB(g.dbConf.DBName).GridFS("fs").Open(path)
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
func GridFSServer(dbConf *conf.DatabaseConfig, pathFile string) http.Handler {
	session, err := mgo.Dial(dbConf.DialServers())
	if err != nil {
		log.Fatalf("Could not conected to %s - %s database: %v", dbConf.DialServers(), dbConf.DBName, err.Error())
	}
	session.SetMode(mgo.Monotonic, true)

	return &gridFSHandler{dbConf, pathFile, session}
}
