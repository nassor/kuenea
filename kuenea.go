package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"labix.org/v2/mgo"
)

var (
	gfs    *mgo.GridFS
	config Config
)

func main() {
	config.ReadConfigFile("/etc/kuenea/kuenea.json")

	mdbSession, err := mgo.Dial(config.DialServers())
	if err != nil {
		panic(err)
	}
	defer mdbSession.Close()

	mdbSession.SetMode(mgo.Monotonic, true)
	session := mdbSession.DB(config.Database.DBName)
	gfs = session.GridFS("fs")

	http.HandleFunc("/"+config.Database.Path, gridHandler)
	s := &http.Server{
		Addr:         config.BindWithPort(),
		ReadTimeout:  time.Duration(config.Http.Timeout) * time.Millisecond,
		WriteTimeout: 0 * time.Second}

	s.ListenAndServe()
}

// Handle server requests, find file and response.
func gridHandler(w http.ResponseWriter, r *http.Request) {
	file, err := gfs.Open(strings.Replace(r.URL.Path[1:], config.Database.Path, "", 1))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", "Not Found")
		return
	}
	
	w.Header().Set("Content-Type", file.ContentType())
	_, err = io.Copy(w, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Can't reach file")
	}
	
	file.Close()
}
