package conf

import (
	"io/ioutil"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Database connection config
type GridFSConfig struct {
	ConnectURI  string `yaml:"connect_uri"` // MongoDB Connection URI
	Path        string `yaml:"path`         // One Path for each Database
	ReadSeeker  bool   `yaml:"read_seeker"`
	CachedItems int    `yaml:"cached_items"`
}

// Filesystem folder config
type LocalFSConfig struct {
	Root       string `yaml:"root`
	Path       string `yaml:"path`
	ReadSeeker bool   `yaml:"read_seeker`
}

// HTTP Server config
type HttpServerConfig struct {
	Bind    string // IP Bind
	Port    int    // Port to use
	Timeout int64  // Conn timeout
}

// Configuration structure of asset server
type Config struct {
	GridFS []GridFSConfig   `yaml:"gridfs"`
	Local  []LocalFSConfig  `yaml:"local"`
	HTTP   HttpServerConfig `yaml:"http"`
}

// Read file json config file and setup asset server
func (config *Config) ReadConfigFile(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	// fmt.Println(string(file))

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return err
	}

	return nil
}

// Return string <bind>:<port> as tcp connect setting
func (config Config) BindWithPort() string {
	return config.HTTP.Bind + ":" + strconv.Itoa(config.HTTP.Port)
}
