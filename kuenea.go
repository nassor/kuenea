package main

import (
	"net/http"
	"time"

	. "kuenea/conf"
	"kuenea/handler"
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

	http.Handle("/"+config.Database.Path, handler.GridFSServer(gfs, &config))
	s := &http.Server{
		Addr:         config.BindWithPort(),
		ReadTimeout:  time.Duration(config.Http.Timeout) * time.Millisecond,
		WriteTimeout: 0 * time.Second}

	s.ListenAndServe()
}
