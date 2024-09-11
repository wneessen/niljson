<!--
SPDX-FileCopyrightText: 2024 Winni Neessen <wn@neessen.dev>

SPDX-License-Identifier: CC0-1.0
-->

# niljson - A simple Go package for (un-)marshalling null-able JSON types

[![GoDoc](https://godoc.org/github.com/wneessen/niljson?status.svg)](https://pkg.go.dev/github.com/wneessen/niljson)
[![codecov](https://codecov.io/gh/wneessen/niljson/branch/main/graph/badge.svg?token=W4QI1RMR4L)](https://codecov.io/gh/wneessen/niljson)
[![Go Report Card](https://goreportcard.com/badge/github.com/wneessen/niljson)](https://goreportcard.com/report/github.com/wneessen/niljson)
[![REUSE status](https://api.reuse.software/badge/github.com/wneessen/niljson)](https://api.reuse.software/info/github.com/wneessen/niljson)
<a href="https://ko-fi.com/D1D24V9IX"><img src="https://uploads-ssl.webflow.com/5c14e387dab576fe667689cf/5cbed8a4ae2b88347c06c923_BuyMeACoffee_blue.png" height="20" alt="buy ma a coffee"></a>

niljson provides a simple and efficient way to handle nullable JSON fields during the (un-)marshalling process. 
In JSON, it's common to encounter fields that can be `null`, but handling these fields in Go can be cumbersome, 
especially when dealing with primitive types like `int`, `float64`, and `bool`. These types can all be either `0` 
(as a value) or `null`. In Go you can always work with pointers, but these can lead to unhandled nil 
pointer dereferences.

**niljson** addresses this challenge by offering a set of types that can seamlessly handle `null` values during 
unmarshalling, allowing your Go applications to work with JSON data more naturally and with fewer boilerplate 
checks for `nil` values.

### Key Features

- **Nullable Types**: Provides a range of nullable types (`NilString`, `NilInt`, `NilFloat`, `NilBool`, etc.) that 
  are easy to use and integrate into your existing Go structs.
- **JSON Unmarshalling Support**: Automatically handles the (un-)marshalling of JSON fields, converting `null` JSON 
  values to Go's `nil` or zero values, depending on the context.
- **Minimalistic and Lightweight**: Designed to be lightweight and unobtrusive, so it won't bloat your application 
  or introduce unnecessary dependencies (only relies on the Go standard library)

### Example Usage

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
    
    "github.com/wneessen/niljson"
)

type JSONType struct {
    Bool       niljson.NilBoolean `json:"bool"`
    Float32    niljson.NilFloat32 `json:"float32,omitempty"`
    Float64    niljson.NilFloat64 `json:"float64"`
    Int        niljson.NilInt     `json:"int"`
    Int64      niljson.NilInt64   `json:"int64"`
    NullString niljson.NilString  `json:"nil"`
    String     niljson.NilString  `json:"string"`
}

func main() {
  data := []byte(`{
 		"bytes": "Ynl0ZXM=",
		"bool": true,
		"float32": null,
		"float64":0,
		"int": 123,
		"int64": 12345678901234,
		"nilvalue": null,
		"string":"test"
	}`)

  var example JSONType
  var output string
  if err := json.Unmarshal(data, &example); err != nil {
    fmt.Println("failed to unmarshal JSON:", err)
    os.Exit(1)
  }

  if example.Bool.NotNil() {
    output += fmt.Sprintf("Bool is: %t, ", example.Bool.Value())
  }
  if example.Float32.IsNil() {
    output += "Float 32 is nil, "
  }
  if example.Float64.NotNil() {
    output += fmt.Sprintf("Float 64 is: %f, ", example.Float64.Value())
  }
  if example.String.NotNil() {
    output += fmt.Sprintf("String is: %s", example.String.Value())
  }
  fmt.Println(output)
  
  data, err := json.Marshal(&example)
  if err != nil {
    fmt.Printf("failed to marshal JSON: %s", err)
    os.Exit(1)
  }
  fmt.Println(data)
}
```
