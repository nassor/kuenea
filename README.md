# Kuenea

Simple Asset Server for [GridFS](http://docs.mongodb.org/manual/applications/gridfs/).

## Benchmark
__Hardware__: Intel® Core™ i7-2720QM CPU @ 2.20GHz / 6GB DDR3-1333 / 7200RPM SATA Disk | Ubuntu 12.04
__Set__: Reach 72.2kb image / requests = 1k / concurrency = 100 | _only local_
__Command__: `ab -n 1000 -c 100 http://localhost:8080/file.png`

### Ruby Metal:
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
    Document Length:        73800 bytes

    Concurrency Level:      100
    Time taken for tests:   0.309 seconds
    Complete requests:      1000
    Failed requests:        0
    Write errors:           0
    Total transferred:      73901000 bytes
    HTML transferred:       73800000 bytes
    Requests per second:    3237.52 [#/sec] (mean) <<<<<<<
    Time per request:       30.888 [ms] (mean)
    Time per request:       0.309 [ms] (mean, across all concurrent requests)
    Transfer rate:          233648.71 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0    0   0.6      0       3
    Processing:    12   29   3.4     30      38
    Waiting:       12   29   3.4     30      38
    Total:         15   30   3.4     30      40

    Percentage of the requests served within a certain time (ms)
      50%     30
      66%     31
      75%     32
      80%     33
      90%     34
      95%     35
      98%     36
      99%     38
     100%     40 (longest request)


## TODO
* __Tests__
* __Multi database/path support__
* __Improve Error Messages__
* Config by cmd flags
* Improve Docs
* SysV init config file

### Thank You
Go Community
Gustavo Niemeyer and Contributors for awesome [MongoDB Driver (mgo)](http://labix.org/mgo)
