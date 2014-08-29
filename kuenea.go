package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rossan/kuenea/conf"
	"github.com/rossan/kuenea/handler"
)

func main() {

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

	log.Println("Starting Kuenea file server at " + config.BindWithPort())

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
	for _, gridConf := range config.GridFS {
		http.Handle(fmt.Sprintf("/%v", gridConf.Path), handler.GridFSServer(gridConf))
	}
	return nil
}

func loadPaths(config conf.Config) error {
	for _, localConf := range config.Local {
		localConf.Path = strings.Trim(localConf.Path, "/")
		localConf.Path = "/" + localConf.Path + "/"
		http.Handle(localConf.Path, handler.LocalFSServer(localConf))
	}
	return nil
}
