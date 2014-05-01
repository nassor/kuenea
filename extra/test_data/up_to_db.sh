curl http://clips.vorwaerts-gmbh.de/big_buck_bunny.webm > bbb.webm
curl http://techslides.com/demos/sample-videos/small.mp4 > robot.mp4
mongofiles -d kuenea_test put bbb.webm
mongofiles -d kuenea_test put robot.mp4
rm bbb.webm
rm robot.mp4
mongofiles -d kuenea_test put gopher.png
