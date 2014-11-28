deps:
	go get -u gopkg.in/yaml.v2
	go get -u gopkg.in/mgo.v2
	go get -u github.com/golang/groupcache/lru
	
run:
	cd src/woh && ../../bin/fresh -c runner.conf
