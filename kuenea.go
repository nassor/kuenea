package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"kuenea/conf"
	"kuenea/handler"
	"labix.org/v2/mgo"
)

func main() {
	log.Println("Starting Kuenea file server...")

	var config conf.Config

	err := config.ReadConfigFile("/etc/kuenea/kuenea.json")
	if err != nil {
		log.Fatalf("Could not read config file: %v", err.Error())
	}

	for _, db := range config.Databases {
		mdbSession, err := mgo.Dial(db.DialServers())
		if err != nil {
			log.Fatalf("Could not contact server %v: %v", db.DialServers(), err.Error())
		}
		defer mdbSession.Close()

		mdbSession.SetMode(mgo.Monotonic, true)
		session := mdbSession.DB(db.DBName)

		log.Printf("MongoDB: %v:%v -> %v", db.DialServers(), db.DBName, db.Path)
		http.Handle(fmt.Sprintf("/%v", db.Path), handler.GridFSServer(session.GridFS("fs"), db.Path))
	}

	for _, local := range config.Local {
		log.Printf("LocalFS: %v -> %v", local.Root, local.Path)
		http.Handle("/"+local.Path, http.StripPrefix("/"+local.Path, http.FileServer(http.Dir(local.Root))))
	}

	s := &http.Server{
		Addr:         config.BindWithPort(),
		ReadTimeout:  time.Duration(config.Http.Timeout) * time.Millisecond,
		WriteTimeout: 0}

	log.Fatal(s.ListenAndServe())
}
