# StructTag
[![Build Status](https://travis-ci.org/mehrdadrad/stags.svg?branch=master)](https://travis-ci.org/mehrdadrad/stags)
[![Go Report Card](https://goreportcard.com/badge/github.com/mehrdadrad/stags)](https://goreportcard.com/report/github.com/mehrdadrad/stags)
[![GoCover](https://gocover.io/_badge/github.com/mehrdadrad/stags)](https://gocover.io/github.com/mehrdadrad/stags)
[![Godoc](https://godoc.org/github.com/mehrdadrad/stags?status.svg)](https://godoc.org/github.com/mehrdadrad/stags)

Go library to get and convert struct tags

## Features
- supports nested structure
- convert the tag's value to given data type


## Installation

Standard `go get`:

```bash
$ go get github.com/jinwong001/structTag
```

## Usage & Example


```go
type Person struct {
    Address *Address `json:"address" bson:"address"`
    Age     int
    Name    string `json:"name"`
}

type Address struct {
    Nation string `json:"nation" bson:"nation"`
    City   string `json:"city" bson:"city"`
    Street string `json:"street" bson:"-"`
}
```

```go
func TestTagPoint(t *testing.T) {
    cfg := Person{}
    s := New(cfg)
    t.Log("tag:", s.GetPoint("bson", "Address.Nation"))
    t.Log("tag:", s.GetPoint("bson", "Address.City"))
    t.Log("tag:", s.GetPoint("bson", "Address.Street"))
    t.Log("tag:", s.GetPoint("bson", "Age"))
    t.Log("tag:", s.GetPoint("bson", "Name"))
}

```

## License
This project is licensed under MIT license. Please read the LICENSE file.