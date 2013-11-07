# Kuenea

File Server using [GridFS](http://docs.mongodb.org/manual/applications/gridfs/) or/and filesystem over HTTP

If you have a distributed application, is using mongodb, prefer keep all your assets in mongodb using gridfs and you think filesystem or CDN solutions are too painful to manage, probably Kuenea is for you.

_Tested with go 1.1.2 and labix/mgo r2013.05.19_

## Benchmark
* __Hardware__: Intel® Core™ i7-2720QM CPU @ 2.20GHz / 6GB DDR3-1333 / 7200RPM SATA Disk | Ubuntu 13.04
* __Set__: Reach 55.6kB image _only local requests_
* __Software__: `Apache Benchmark`

__Requests: 10000 / Concurrency: 1000__

|Server      |Req/s     |Time taken  |
|------------|----------|------------|
|Kuenea - GridFS | 5070.85   |1.972 s |
|Kuenea - Filesystem | 10806.92   |0.925 s  |

## TODO
* __Tests__ _(still learning how to do in Go)_
* __Dial via Mongodb connection string__
* __Improve Error Messages__
* HTTP Cache-Control
* GroupCache Support
* Improve Docs

## Thank You
* Go Community
* Gustavo Niemeyer and Contributors for awesome [MongoDB Driver (mgo)](http://labix.org/mgo)
