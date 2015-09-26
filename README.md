[![GoDoc](https://godoc.org/github.com/AIPA-team/db-connection-pooling-deamon?status.svg)](https://godoc.org/github.com/AIPA-team/db-connection-pooling-deamon)
[![Build Status](https://travis-ci.org/AIPA-team/db-connection-pooling-daemon.svg)](https://travis-ci.org/AIPA-team/db-connection-pooling-daemon)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

# db-connection-pooling-deamon

Is a daemon that allows usage of pooling database connections via RESTful API for languages/frameworks that dont have such capabilities.  
Initially it supports postgres. MySQL and other databases are work in progress.

This is for sure not finished and production-ready code!

## How to use:

- config config.yml
- build and run the daemon
- call it's `/query` url with json structured as follows:
```json
{  
  "query": "SELECT * from test where name = $1 and lastname = $2",  
  "params": ["jack", "black"]  
}
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
