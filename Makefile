all:	
	go get -u gopkg.in/mgo.v2
	
run:
	cd src/woh && ../../bin/fresh -c runner.conf
