# go-bloomfilter

[![Go Report Card](https://goreportcard.com/badge/github.com/kenju/go-bloomfilter)](https://goreportcard.com/report/github.com/kenju/go-bloomfilter)
[![GoDoc](https://godoc.org/github.com/kenju/go-bloomfilter?status.svg)](http://godoc.org/github.com/kenju/go-bloomfilter)

## Usage

```go
package main

import (
	"github.com/kenju/go-bloomfilter"
)

func main() {
    bf := bloomfilter.New(1024)
    
    bf.Add([]byte("foo"))
    bf.Add([]byte("bar"))
    bf.Add([]byte("buz"))
    
    bf.Size() //=> 3
    bf.Test([]byte("foo")) //=> (maybe) true
    bf.Test([]byte("bar")) //=> (maybe) true
    bf.Test([]byte("buz")) //=> (maybe) true
    bf.Test([]byte("hey")) //=> false
    bf.Test([]byte("!!!")) //=> false
}

```