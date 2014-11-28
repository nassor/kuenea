package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nassor/kuenea/conf"
	"gopkg.in/mgo.v2"
)

var (
	mongoServer string
	session     mgo.Session
)

func setup() {
	mongoServer = "localhost:27017/kuenea_test"
	session, _ := mgo.Dial(mongoServer)
	defer session.Close()

	file, _ := session.DB("").GridFS("fs").Create("test_file.txt")
	file.Write([]byte("Hello world!"))
	file.Close()

}

func tearDown() {
	session, _ := mgo.Dial(mongoServer)
	session.DB("").GridFS("fs").Remove("test_file.txt")
	defer session.Close()
}

func TestGridFSServer(t *testing.T) {
	setup()

	gfsConfig := conf.GridFSConfig{mongoServer, "test/", false}
	server := httptest.NewServer(GridFSServer(gfsConfig))
	defer server.Close()

	res, err := http.Get(server.URL + "/test/test_file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	if string(fileData) != "Hello world!" {
		t.Fail()
	}

	tearDown()
}
