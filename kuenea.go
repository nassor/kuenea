package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"kuenea/conf"
	"kuenea/handler"
)

func main() {
	log.Println("Starting Kuenea file server...")

	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err.Error())
	}

	err = loadGridsFS(config)
	if err != nil {
		log.Fatalf("Could not load databases: %v", err.Error())
	}

	err = loadPaths(config)
	if err != nil {
		log.Fatalf("Could not load paths: %v", err.Error())
	}

	s := &http.Server{
		Addr:         config.BindWithPort(),
		ReadTimeout:  time.Duration(config.Http.Timeout) * time.Millisecond,
		WriteTimeout: 0}

	log.Fatal(s.ListenAndServe())
}

func loadConfig() (conf.Config, error) {
	var config conf.Config

	dir, err := os.Getwd()
	if err != nil {
		return config, err
	}

	var configFile = flag.String("c", "", "location of the configuration file")
	flag.Parse()
	if *configFile == "" {
		*configFile = dir + "/kuenea-config.json"
	}

	err = config.ReadConfigFile(*configFile)
	if err != nil {
		return config, err
	}

	return config, nil
}

func loadGridsFS(config conf.Config) error {
	for _, db := range config.Databases {
		log.Printf("MongoDB: %v -> %v", db.ConnectURI, db.Path)
		http.Handle(fmt.Sprintf("/%v", db.Path), handler.GridFSServer(&db, db.Path))
	}
	return nil
}

func loadPaths(config conf.Config) error {
	for _, local := range config.Local {
		local.Root = strings.TrimSuffix(local.Root, "/")
		local.Path = strings.TrimSuffix(local.Path, "/")
		local.Path = strings.TrimPrefix(local.Path, "/")
		local.Path = "/" + local.Path + "/"
		log.Printf("LocalFS: %v -> %v", local.Root, local.Path)
		http.Handle(local.Path, handler.LocalFSServer(local.Root, local.Path))
	}
	return nil
}
