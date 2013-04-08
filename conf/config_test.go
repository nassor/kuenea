package conf

import (
	"testing"
	"reflect"
)

func readConfigFile() Config {
	var config Config
	config.ReadConfigFile("../deploy/kuenea.json.example")
	return config
}

func TestGridFSDatabases(t *testing.T) {
	var config = readConfigFile()

	if dbType := reflect.TypeOf(config.Databases); dbType.Kind() != reflect.Slice {
		t.Error("databases config is a Slice (one or more databases) or 'nil', test detected a: " + dbType.Kind().String())
	}

	if (config.Databases != nil && config.Databases[0].DBName != "test") {
		t.Error("can't find 'test' database on [test] config file")
	}
}

func TestLocalFS(t *testing.T) {
	var config = readConfigFile() 

	if localType := reflect.TypeOf(config.Local); localType.Kind() != reflect.Slice {
		t.Error("local filesystem config is a Slice (one or more filesystem path) or 'nil', test detected a: " + localType.Kind().String())
	}
}
