> [!INFO]
> [README for v1](README.v1.md)

# goaux/tuple v2

[![Go Reference](https://pkg.go.dev/badge/github.com/goaux/tuple/v2.svg)](https://pkg.go.dev/github.com/goaux/tuple/v2)
[![Go Report Card](https://goreportcard.com/badge/github.com/goaux/tuple/v2)](https://goreportcard.com/report/github.com/goaux/tuple/v2)

A tiny, compile‑time‑safe library that gives you a **generic tuple** type for every size from 2 up to 32 values.
The package is split into a separate sub‑package per tuple length so you only import what you need – no unnecessary dependencies or large monolithic packages.

All source files are produced by `gen-tuple`.

## Features

| Length | Package name | What it provides |
|--------|--------------|------------------|
| 2      | `tuple2`     | `Tuple[A, B]` – two values |
| 3      | `tuple3`     | `Tuple[A, B, C]` – three values |
| 4      | `tuple4`     | … |
| …      | …            | … |
| 31     | `tuple31`    | `Tuple[...AE]` – thirty-one values |
| 32     | `tuple32`    | `Tuple[...AE, AF]` – thirty-two values |

Each package contains:

```go
// Pack creates a new tuple with the supplied values.
func Pack[A, B, …](a A, b B, …) Tuple[A, B, …]

// Unpack returns all elements of the tuple.
func (t Tuple[A, B, …]) Unpack() (A, B, …)
```

The types are fully generic – no type assertions or `interface{}` needed.
They work with any Go value, including structs, interfaces, pointers, etc.

## Installation

```bash
go get github.com/goaux/tuple/v2@latest
```

Because every tuple size lives in its own package you can import only the one(s) you need:

```go
import (
    "github.com/goaux/tuple/v2/tuple3"
    "github.com/goaux/tuple/v2/tuple5"
)
```

## Example

```go
package main

import (
    "fmt"

    "github.com/goaux/tuple/v2/tuple3"
)

func main() {
    t := tuple3.Pack(42, "answer", true)

    a, b, c := t.Unpack()
    fmt.Println(a, b, c) // 42 answer true
}
```

## Build Constraints

This module uses Go build constraints to allow you to selectively include or exclude certain features (Comparison and Getter methods) during compilation. This is useful for reducing binary size or simplifying the API surface in specific use cases.

### Available Build Tags

#### 1. Comparison Methods (`Compare`)
By default, `Compare` methods are included. If you want to exclude them globally, use the `notuplecompare` tag. If you need to explicitly re-enable them for a specific package after excluding them, use the corresponding package-specific tag.

- **To exclude all comparison methods:** `-tags notuplecompare`
- **To specifically enable comparison for a package (even if excluded globally):**
  - `tuple2compare`
  - `tuple3compare`
  - ...
  - `tuple32compare`

#### 2. Getter Methods (`GetA`, `GetB`, etc.)
By default, getter methods are included. If you want to exclude them globally, use the `notuplegetter` tag. If you need to explicitly re-enable them for a specific package after excluding them, use the corresponding package-specific tag.

- **To exclude all getter methods:** `-tags notuplegetter`
- **To specifically enable getters for a package (even if excluded globally):**
  - `tuple2getter`
  - `tuple3getter`
  - ...
  - `tuple32getter`

### Example Usage

To compile a project while excluding all comparison and getter methods to minimize the API surface:

```bash
go build -tags "notuplecompare,notuplegetter" .
```

To exclude all comparison methods and all getter but keep the getters for `tuple2`:

```bash
go build -tags "notuplecompare,notuplegetter,tuple2getter" .
```
