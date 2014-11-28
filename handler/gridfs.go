package handler

import (
	"io"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang/groupcache/lru"
	"github.com/nassor/kuenea/conf"
	"gopkg.in/mgo.v2"
)

// Handle server requests, find file and response.
func GridFSServer(gridFS conf.GridFSConfig) http.Handler {
	session, err := mgo.Dial(gridFS.ConnectURI)
	if err != nil {
		log.Fatalf("Could not conected to database: %v", err.Error())
	}
	session.SetMode(mgo.Monotonic, true)
	log.Printf("MongoDB: %v -> %v", session.LiveServers(), gridFS.Path)

	assetsCache := lru.New(gridFS.CachedItems)
	assetsCache.OnEvicted = func(key lru.Key, value interface{}) {
		file := value.(*mgo.GridFile)
		defer file.Close()
	}

	return &gridFSHandler{&gridFS, gridFS.Path, session, assetsCache}
}

type gridFSHandler struct {
	gridFS      *conf.GridFSConfig
	PathFile    string
	Session     *mgo.Session
	assetsCache *lru.Cache
}

func (g *gridFSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var file *mgo.GridFile
	var err error

	path := strings.Replace(r.URL.Path[1:], g.PathFile, "", 1)

	// Find and return file from cache
	if cachedFile, ok := g.assetsCache.Get(path); ok {
		file, ok = cachedFile.(*mgo.GridFile)
		file.Seek(0, 0)
		if !ok {
			file, err = g.Session.DB("").GridFS("fs").Open(path)
		}
	} else {
		file, err = g.Session.DB("").GridFS("fs").Open(path)
		g.assetsCache.Add(path, file)
	}

	// If stuff dont feel good
	if err != nil {
		// Try reconnect database to reach file again
		g.Session.Refresh()
		time.Sleep(1 * time.Second)
		file, err = g.Session.DB("").GridFS("fs").Open(path)
		// Nothing to do
		if err != nil {
			log.Printf(err.Error())
			http.NotFound(w, r)
			return
		}
	}

	if g.gridFS.ReadSeeker {
		http.ServeContent(w, r, file.ContentType(), file.UploadDate(), file)
	} else {

		if t, err := time.Parse(http.TimeFormat, r.Header.Get("If-Modified-Since")); err == nil && file.UploadDate().Before(t.Add(2*time.Second)) {
			delete(w.Header(), "Content-Type")
			delete(w.Header(), "Content-Length")
			w.WriteHeader(http.StatusNotModified)
			return
		}

		if file.ContentType() == "" {
			contentType := mime.TypeByExtension(filepath.Ext(file.Name()))
			w.Header().Set("Content-Type", contentType)
		} else {
			w.Header().Set("Content-Type", file.ContentType())
		}

		w.Header().Set("Last-Modified", file.UploadDate().UTC().Format(http.TimeFormat))

		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, "gridfs read error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
