# Mira [![Go Report Card](https://goreportcard.com/badge/github.com/stevenferrer/mira)](https://goreportcard.com/report/github.com/stevenferrer/mira)

This package wraps the `reflect` package and adds some useful methods for introspecting Go values. 

## Installation

```console
$ go get github.com/sf9v/mira
```

## Usage

```go

import (
    "fmt"

    "github.com/sf9v/mira"
)

func main() {
    v := int64(1)
    typeInfo := mira.NewTypeInfo(v)

    // check if value is numeric
    if typeInfo.IsNumeric() {
        fmt.Println("value is a numeric type")
    }

    // check if value is nillable
    if typeInfo.IsNillable() {
        fmt.Println("value is nillable")
    }

    // print type's name
    fmt.Println(typeInfo.Name())
    // print type's pkg path
    fmt.Println(typeInfo.PkgPath())
    // print the interface value
    fmt.Printf("%v\n", typeInfo.V())
    // print the reflect.Type's string representation
    fmt.Printf("%v", typeInfo.T().String())
    ...
}
```

## License

MIT