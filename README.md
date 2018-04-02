# rmforks

[![Build Status](https://travis-ci.org/royge/rmforks.svg?branch=master)](https://travis-ci.org/royge/rmforks)

Delete github forked repos using API

## Installation

### Using `go get`

	```$ go get github.com/royge/rmforks```

	After successful install the binary will be inside `$GOPATH/bin` directory.

### From source

1. Clone repo

	```$ git clone https://github.com/royge/rmforks.git```

1. Install `vgo`

	```
	$ go get -u golang.org/x/vgo
	```

1. Build

	```
	$ cd rmforks
	$ CC=gcc vgo build
	```

1. How to use

	Copy/rename ```config.json.dist``` to ```config.json``` and provide your github username and access token.

	Run ```$ ./rmforks```

## TODO

Create repo package tests.
