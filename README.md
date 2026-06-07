# tuple
Package tuple provides extremely simple generic tuple types and utility functions for creating tuples.

[![Go Reference](https://pkg.go.dev/badge/github.com/goaux/tuple.svg)](https://pkg.go.dev/github.com/goaux/tuple)
[![Go Report Card](https://goreportcard.com/badge/github.com/goaux/tuple)](https://goreportcard.com/report/github.com/goaux/tuple)

## Usage Example

Here's an example of how you might use the tuple package for managing multiple
return values from functions:

### Handling Multiple Return Values

Consider a scenario where you have a function that returns two integers and
another function that accepts a string and a pair of integers.

```go
package main

import (
	"fmt"

	"github.com/goaux/tuple"
)

// GetXY is an example function returning a pair of values.
func GetXY() (x, y int) {
	return 12, 34
}

// PrintXY is an example function accepting a string and a pair of ints.
func PrintXY(tag string, x, y int) {
	fmt.Printf("%s: X=%d Y=%d\n", tag, x, y)
}

// PrintXYPair is an example function wraps PrintXY with accepting a string and a Pair[int, int].
func PrintXYPair(tag string, xy tuple.Pair[int, int]) {
	PrintXY(tag, xy.First, xy.Second)
}

func main() {
	// Cannot write PrintXY("here", GetXY()) directly.
	x, y := GetXY()
	PrintXY("here", x, y)

	// Using tuple eliminates the need for intermediate variables.
	PrintXYPair("here", tuple.MakePair(GetXY()))
}
```

## Tuples Types

### Pair[T0, T1]

Represents a pair of values.

- `First` holds the first value.
- `Second` holds the second value.

### Triple[T0, T1, T2]

Represents a triple of values.

- `First`, `Second`, and `Third` hold the respective values.

### Quadruple[T0, T1, T2, T3]

Represents a quadruple of values.

- `First`, `Second`, `Third`, and `Fourth` hold the respective values.

### Quintuple[T0, T1, T2, T3, T4]

Represents a quintuple of values.

- `First`, `Second`, `Third`, `Fourth`, and `Fifth` hold the respective values.

## Utility Functions

Each tuple type includes utility functions:

- `MakePair`, `NewPair`: Create new pairs.
- `Len() int`: Returns the number of elements in the tuple.
- `Get(int) any`: Gets an element at index i. Returns nil if out-of-bounds.

Similar methods exist for `Triple`, `Quadruple`, and `Quintuple`.

## Tuple Interface

The `Tuple` interface defines two methods: `Len() int` and `Get(int) any`.
`Pair`, `Triple`, `Quadruple`, and `Quintuple` all implement the `Tuple` interface.

