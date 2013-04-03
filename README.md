# Kuenea

Simple Asset Server using [GridFS](http://docs.mongodb.org/manual/applications/gridfs/) or filesystem.

_Only tested with go tip, you may have problems with go 1.0_

## Why Kuenea?
After trying use gridfs-fuse and nginx-gridfs without success i decided develop gridfs asset server for web projects.

## Benchmark
* __Hardware__: Intel® Core™ i7-2720QM CPU @ 2.20GHz / 6GB DDR3-1333 / 7200RPM SATA Disk | Ubuntu 12.04
* __Set__: Reach 55.6kB image _only local requests_
* __Software__: `Apache Benchmark`

__Requests: 1000 / Concurrency: 100__

|Server      |Req/s     |Time taken  |Time per Req       |
|------------|----------|------------|-------------------|
|Kuenea(Go)  |4944.67   |0.202 s     |20.224 [ms] (mean) |
|Node.js     |2060.11   |0.485 s     |48.541 [ms] (mean) |
|Rack(Ruby)  |408.02    |2.451 s     |245.084 [ms] (mean)|


## TODO
* __Tests__
* __DB Auth__
* __Improve Error Messages__
* Cache-Control
* Config by cmd flags
* Improve Docs
* SysV init config file

## Thank You
* Go Community
* Gustavo Niemeyer and Contributors for awesome [MongoDB Driver (mgo)](http://labix.org/mgo)
