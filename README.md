# Kuenea

Simple Asset Server for [GridFS](http://docs.mongodb.org/manual/applications/gridfs/).

_Only tested with go tip, you may have problems with go 1.0_

## Benchmark
* __Hardware__: Intel® Core™ i7-2720QM CPU @ 2.20GHz / 6GB DDR3-1333 / 7200RPM SATA Disk | Ubuntu 12.04
* __Set__: Reach 72.2kb image / requests = 1k / concurrency = 100 | _only local_
* __Command__: `ab -n 1000 -c 100 http://localhost:8080/file.png`


### Rails Metal:
    Document Length:        73800 bytes

    Concurrency Level:      100
    Time taken for tests:   24.722 seconds
    Complete requests:      1000
    Failed requests:        0
    Write errors:           0
    Total transferred:      74149000 bytes
    HTML transferred:       73800000 bytes
    Requests per second:    40.45 [#/sec] (mean) <<<<<<<
    Time per request:       2472.246 [ms] (mean)
    Time per request:       24.722 [ms] (mean, across all concurrent requests)
    Transfer rate:          2928.96 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    1   1.5      0       6
    Processing:    33 2353 425.6   2461    2642
    Waiting:       16 2344 425.8   2452    2633
    Total:         39 2354 424.3   2461    2642

    Percentage of the requests served within a certain time (ms)
      50%   2461
      66%   2497
      75%   2518
      80%   2531
      90%   2551
      95%   2569
      98%   2583
      99%   2608
     100%   2642 (longest request)

### Kuenea Asset Server
    Concurrency Level:      100
    Time taken for tests:   0.237 seconds
    Complete requests:      1000
    Failed requests:        0
    Write errors:           0
    Total transferred:      73901000 bytes
    HTML transferred:       73800000 bytes
    Requests per second:    4227.40 [#/sec] (mean)
    Time per request:       23.655 [ms] (mean)
    Time per request:       0.237 [ms] (mean, across all concurrent requests)
    Transfer rate:          305087.02 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    0   0.6      0       3
    Processing:    15   23   4.0     22      36
    Waiting:       15   22   3.9     22      35
    Total:         15   23   4.3     22      37

    Percentage of the requests served within a certain time (ms)
      50%     22
      66%     24
      75%     25
      80%     26
      90%     28
      95%     33
      98%     36
      99%     37
     100%     37 (longest request)


## TODO
* __Tests__
* __DB Auth__
* __Multi database/path support__
* __Improve Error Messages__
* Cache-Control
* Config by cmd flags
* Improve Docs
* SysV init config file

## Thank You
* Go Community
* Gustavo Niemeyer and Contributors for awesome [MongoDB Driver (mgo)](http://labix.org/mgo)
