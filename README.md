# rmforks

[![Go](https://github.com/royge/rmforks/actions/workflows/go.yml/badge.svg)](https://github.com/royge/rmforks/actions/workflows/go.yml)

Delete github forked repos using API

## Installation

### Using `go get`

```
$ go get github.com/royge/rmforks
```

After successful install the binary will be inside `$GOPATH/bin` directory.

### From source

**NOTE:** Requires `go 1.10`.

1. Clone repo

    ```
    $ git clone https://github.com/royge/rmforks.git
    ```

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

    Run

    ```
    $ ./rmforks
    ```

## TODO

Create repo package tests.
