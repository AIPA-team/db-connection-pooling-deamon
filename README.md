[![GoDoc](https://godoc.org/github.com/AIPA-team/db-connection-pooling-deamon?status.svg)](https://godoc.org/github.com/AIPA-team/db-connection-pooling-deamon)
[![Build Status](https://drone.io/github.com/AIPA-team/db-connection-pooling-daemon/status.png)](https://drone.io/github.com/AIPA-team/db-connection-pooling-daemon/latest)
[![Coverage Status](https://coveralls.io/repos/AIPA-team/db-connection-pooling-daemon/badge.svg?branch=master&service=github)](https://coveralls.io/github/AIPA-team/db-connection-pooling-daemon?branch=master)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

# db-connection-pooling-deamon

Is a daemon that allows usage of pooling database connections via RESTful API for languages/frameworks that doesn't have such capabilities which leads to noticable increase of response time for each database query.
Initially it supports postgres. MySQL and other databases are work in progress. MySQL should be supported in few days.

This is for sure not finished and production-ready code yet!

## How to use:

- config config.yml (yaml syntax)  
example:
```
port: 7000
db:
  dbName   : "pool"       #name of database
  password : "test"       #password to database
  host     : "localhost"  #IP address
  user     : "postgres"   #username
  poolMin  : 2            #minimum number of always hold connections in pool
  poolMax  : 20           #maximum number of avalible connections in pool
```
- build and run the daemon
- call it with `POST` method to `localhost:7000/query` URL with json structured as follows:
```json
{  
  "query": "SELECT * from test where name = $1 and lastname = $2",  
  "params": ["jack", "black"]  
}
```
- it will return json with
```json
  [
    {"columnName1": "columnValue", "columnName2": "columnValue"}
    {"columnName1": "columnValue", "columnName2": "columnValue"}
  ]
```
- or in case of error
```json 
{"err": "error message"} 
```

## TODO:
- jwt/basic auth authentification
- mysql support
- add more tests
- refactor code to make it more idiomatic

# License: The MIT License (MIT)

Copyright (c) 2015 AIPA-team

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
