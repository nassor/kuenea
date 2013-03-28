package main

import (
	"fmt"
	"net/http"
	"time"
	"strings"

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
	session := mdbSession.DB(config.DBName)
	gfs = session.GridFS("fs")

	http.HandleFunc("/" + config.Path, gridHandler)
	s := &http.Server{
		Addr:         config.BindWithPort(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 0 * time.Second}

	s.ListenAndServe()
}

// Handle server requests, find file and response.
func gridHandler(w http.ResponseWriter, r *http.Request) {
	file, err := gfs.Open(strings.Replace(r.URL.Path[1:], config.Path, "", 1))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", "Not Found")
		return
	}

	b := make([]byte, file.Size())
	file.Read(b)
	w.Header().Set("Content-Type", file.ContentType())
	fmt.Fprintf(w, "%s", b)
	file.Close()
}
