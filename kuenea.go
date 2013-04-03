package main

import (
	"log"
	"net/http"
	"time"

	"kuenea/conf"
	"kuenea/handler"
	"labix.org/v2/mgo"
)

func main() {
	log.Println("Start Kuenea Assets Server")

	var config conf.Config
	config.ReadConfigFile("/etc/kuenea/kuenea.json")

	for server := range config.Databases {
		mdbSession, err := mgo.Dial(config.Databases[server].DialServers())
		if err != nil {
			panic(err)
		}
		defer mdbSession.Close()

		mdbSession.SetMode(mgo.Monotonic, true)
		session := mdbSession.DB(config.Databases[server].DBName)

		log.Println("MongoDB: " + config.Databases[server].DialServers() + ":" + config.Databases[server].DBName + " -> " + config.Databases[server].Path)
		http.Handle("/"+config.Databases[server].Path, handler.GridFSServer(session.GridFS("fs"), config.Databases[server].Path))
	}

	for folder := range config.Local {
		log.Println("LocalFS: " + config.Local[folder].Root + " -> " + config.Local[folder].Path)
		http.Handle("/"+config.Local[folder].Path, http.StripPrefix("/"+config.Local[folder].Path, http.FileServer(http.Dir(config.Local[folder].Root))))
	}

	s := &http.Server{
		Addr:         config.BindWithPort(),
		ReadTimeout:  time.Duration(config.Http.Timeout) * time.Millisecond,
		WriteTimeout: 0 * time.Second}

	log.Fatal(s.ListenAndServe())
}
