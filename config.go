package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Configuration structure of asset server
type Config struct {
	Bind    string   // IP Bind
	Port    int      // Port to use
	Path    string   // Path for server execution
	Servers []string // MongoDB Server for mgo.Dial
	DBName  string   // MongoDB Database
}

// Read file json config file and setup asset server
func (config *Config) ReadConfigFile(fileName string) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		log.Fatal(err)
	}
}

// Return a string with all mongodb servers.
// Used by mgo.Dial()
func (config Config) DialServers() string {
	return strings.Join(config.Servers, ",")
}

// Return string <bind>:<port> as tcp connect setting
func (config Config) BindWithPort() string {
	return config.Bind + ":" + strconv.Itoa(config.Port)
}
