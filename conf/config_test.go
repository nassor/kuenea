package conf

import (
	"reflect"
	"testing"
)

func readConfigFile() Config {
	var config Config
	config.ReadConfigFile("../kuenea-config.json.example")
	return config
}

func TestGridFSDatabases(t *testing.T) {
	var config = readConfigFile()

	if dbType := reflect.TypeOf(config.GridFS); dbType.Kind() != reflect.Slice {
		t.Error("databases config is a Slice (one or more databases) or 'nil', test detected a: " + dbType.Kind().String())
	}

	if config.GridFS[0].ConnectURI != "mongodb://localhost/kuenea_test" {
		t.Error("can't find 'kuenea_test' for database on example config file")
	}

	if config.GridFS[0].Path != "fast/" {
		t.Error("can't find 'img/' path for database on example config file")
	}

	if config.GridFS[1].Path != "seeker/" {
		t.Error("can't find 'video/' path for database on example config file")
	}
}

func TestLocalFS(t *testing.T) {
	var config = readConfigFile()

	if localType := reflect.TypeOf(config.Local); localType.Kind() != reflect.Slice {
		t.Error("local filesystem config is a Slice (one or more filesystem path) or 'nil', test detected a: " + localType.Kind().String())
	}

	if config.Local != nil && config.Local[0].Root != "/home/user/go/src/kuenea/extra/test_data" {
		t.Error("can't find Root '/home/user/go/src/kuenea/extra/test_data' for Local on example config file")
	}

	if config.Local != nil && config.Local[0].Path != "fs_test/" {
		t.Error("can't find 'fs_test/' path for database on example config file")
	}
}
