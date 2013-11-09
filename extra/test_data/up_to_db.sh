curl http://clips.vorwaerts-gmbh.de/big_buck_bunny.webm > bbb.webm
mongofiles -d kuenea_test put bbb.webm
rm bbb.webm
mongofiles -d kuenea_test put gopher.png
