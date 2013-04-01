package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// Database connection config
type DatabaseConfig struct {
	Servers []string // MongoDB Server for mgo.Dial
	DBName  string   // MongoDB Database
	DBUser  string
	DBPass  string
	Path    string // One Path for each Database
}

// HTTP Server config
type HttpServerConfig struct {
	Bind    string // IP Bind
	Port    int    // Port to use
	Timeout int64  // Conn timeout
}

// Configuration structure of asset server
type Config struct {
	Database DatabaseConfig
	Http     HttpServerConfig
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
	return strings.Join(config.Database.Servers, ",")
}

// Return string <bind>:<port> as tcp connect setting
func (config Config) BindWithPort() string {
	return config.Http.Bind + ":" + strconv.Itoa(config.Http.Port)
}
