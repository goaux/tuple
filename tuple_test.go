package tuple_test

import (
	"fmt"
	"slices"
	"testing"

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

func Example() {
	// Cannot write PrintXY("here", GetXY()) directly.
	x, y := GetXY()
	PrintXY("here", x, y)

	// Using tuple eliminates the need for intermediate variables.
	PrintXYPair("here", tuple.MakePair(GetXY()))
	// Output:
	// here: X=12 Y=34
	// here: X=12 Y=34
}

var _ tuple.Tuple = tuple.MakePair(0, 1)
var _ tuple.Tuple = tuple.MakeTriple(0, 1, 2)
var _ tuple.Tuple = tuple.MakeQuadruple(0, 1, 2, 3)
var _ tuple.Tuple = tuple.MakeQuintuple(0, 1, 2, 3, 4)

func TestPair(t *testing.T) {
	t.Run("MakePair", func(t *testing.T) {
		pair := tuple.MakePair(1, "test")
		if pair.First != 1 {
			t.Errorf("First element should be 1, got %v", pair.First)
		}
		if pair.Second != "test" {
			t.Errorf("Seconc element should be 'test', got %v", pair.Second)
		}
		if pair.Len() != 2 {
			t.Errorf("Expected Len to be 2, got %d", pair.Len())
		}
		if first, ok := pair.Get(0).(int); !ok || first != 1 {
			t.Errorf("First element should be 1, got %v", first)
		}
		if second, ok := pair.Get(1).(string); !ok || second != "test" {
			t.Errorf("Second element should be 'test', got %s", second)
		}
		if v, ok := pair.Get(2).(int); v != 0 || ok {
			t.Errorf("Out of bounds element should be a zero value, got %v", v)
		}
	})

	t.Run("NewPair", func(t *testing.T) {
		pair := tuple.NewPair(1, "test")
		if pair.First != 1 {
			t.Errorf("First element should be 1, got %v", pair.First)
		}
		if pair.Second != "test" {
			t.Errorf("Seconc element should be 'test', got %v", pair.Second)
		}
		if pair.Len() != 2 {
			t.Errorf("Expected Len to be 2, got %d", pair.Len())
		}
		if first, ok := pair.Get(0).(int); !ok || first != 1 {
			t.Errorf("First element should be 1, got %v", first)
		}
		if second, ok := pair.Get(1).(string); !ok || second != "test" {
			t.Errorf("Second element should be 'test', got %s", second)
		}
		if v, ok := pair.Get(2).(int); v != 0 || ok {
			t.Errorf("Out of bounds element should be a zero value, got %v", v)
		}
	})
}

func TestTriple(t *testing.T) {
	t.Run("MakeTriple", func(t *testing.T) {
		triple := tuple.MakeTriple(1, "test", true)
		if triple.First != 1 {
			t.Errorf("First element should be 1, got %v", triple.First)
		}
		if triple.Second != "test" {
			t.Errorf("Second element should be 'test', got %v", triple.Second)
		}
		if triple.Third != true {
			t.Errorf("Third element should be true, got %v", triple.Third)
		}
		if triple.Len() != 3 {
			t.Errorf("Expected Len to be 3, got %d", triple.Len())
		}
		if first, ok := triple.Get(0).(int); !ok || first != 1 {
			t.Errorf("First element should be 1, got %v", first)
		}
		if second, ok := triple.Get(1).(string); !ok || second != "test" {
			t.Errorf("Second element should be 'test', got %s", second)
		}
		if third, ok := triple.Get(2).(bool); !ok || third != true {
			t.Errorf("Third element should be true, got %v", third)
		}
		if v, ok := triple.Get(3).(int); v != 0 || ok {
			t.Errorf("Out of bounds element should be a zero value, got %v", v)
		}
	})

	t.Run("NewTriple", func(t *testing.T) {
		triple := tuple.NewTriple(1, "test", true)
		if triple.First != 1 {
			t.Errorf("First element should be 1, got %v", triple.First)
		}
		if triple.Second != "test" {
			t.Errorf("Second element should be 'test', got %v", triple.Second)
		}
		if triple.Third != true {
			t.Errorf("Third element should be true, got %v", triple.Third)
		}
		if triple.Len() != 3 {
			t.Errorf("Expected Len to be 3, got %d", triple.Len())
		}
		if first, ok := triple.Get(0).(int); !ok || first != 1 {
			t.Errorf("First element should be 1, got %v", first)
		}
		if second, ok := triple.Get(1).(string); !ok || second != "test" {
			t.Errorf("Second element should be 'test', got %s", second)
		}
		if third, ok := triple.Get(2).(bool); !ok || third != true {
			t.Errorf("Third element should be true, got %v", third)
		}
		if v, ok := triple.Get(3).(int); v != 0 || ok {
			t.Errorf("Out of bounds element should be a zero value, got %v", v)
		}
	})
}

func TestQuadruple(t *testing.T) {
	t.Run("MakeQuadruple", func(t *testing.T) {
		quadruple := tuple.MakeQuadruple(1, "test", true, 2.0)
		if quadruple.First != 1 {
			t.Errorf("First element should be 1, got %v", quadruple.First)
		}
		if quadruple.Second != "test" {
			t.Errorf("Second element should be 'test', got %v", quadruple.Second)
		}
		if quadruple.Third != true {
			t.Errorf("Third element should be true, got %v", quadruple.Third)
		}
		if quadruple.Fourth != 2.0 {
			t.Errorf("Fourth element should be 2.0, got %v", quadruple.Fourth)
		}
		if quadruple.Len() != 4 {
			t.Errorf("Expected Len to be 4, got %d", quadruple.Len())
		}
		if first, ok := quadruple.Get(0).(int); !ok || first != 1 {
			t.Errorf("First element should be 1, got %v", first)
		}
		if second, ok := quadruple.Get(1).(string); !ok || second != "test" {
			t.Errorf("Second element should be 'test', got %s", second)
		}
		if third, ok := quadruple.Get(2).(bool); !ok || third != true {
			t.Errorf("Third element should be true, got %v", third)
		}
		if fourth, ok := quadruple.Get(3).(float64); !ok || fourth != 2.0 {
			t.Errorf("Fourth element should be 2.0, got %v", fourth)
		}
		if v, ok := quadruple.Get(4).(int); v != 0 || ok {
			t.Errorf("Out of bounds element should be a zero value, got %v", v)
		}
	})

	t.Run("NewQuadruple", func(t *testing.T) {
		quadruple := tuple.NewQuadruple(1, "test", true, 2.0)
		if quadruple.First != 1 {
			t.Errorf("First element should be 1, got %v", quadruple.First)
		}
		if quadruple.Second != "test" {
			t.Errorf("Second element should be 'test', got %v", quadruple.Second)
		}
		if quadruple.Third != true {
			t.Errorf("Third element should be true, got %v", quadruple.Third)
		}
		if quadruple.Fourth != 2.0 {
			t.Errorf("Fourth element should be 2.0, got %v", quadruple.Fourth)
		}
		if quadruple.Len() != 4 {
			t.Errorf("Expected Len to be 4, got %d", quadruple.Len())
		}
		if first, ok := quadruple.Get(0).(int); !ok || first != 1 {
			t.Errorf("First element should be 1, got %v", first)
		}
		if second, ok := quadruple.Get(1).(string); !ok || second != "test" {
			t.Errorf("Second element should be 'test', got %s", second)
		}
		if third, ok := quadruple.Get(2).(bool); !ok || third != true {
			t.Errorf("Third element should be true, got %v", third)
		}
		if fourth, ok := quadruple.Get(3).(float64); !ok || fourth != 2.0 {
			t.Errorf("Fourth element should be 2.0, got %v", fourth)
		}
		if v, ok := quadruple.Get(4).(int); v != 0 || ok {
			t.Errorf("Out of bounds element should be a zero value, got %v", v)
		}
	})
}

func TestQuintuple(t *testing.T) {
	t.Run("MakeQuintuple", func(t *testing.T) {
		quintuple := tuple.MakeQuintuple(1, "test", true, 2.0, []byte{1, 2, 3})
		if quintuple.First != 1 {
			t.Errorf("First element should be 1, got %v", quintuple.First)
		}
		if quintuple.Second != "test" {
			t.Errorf("Second element should be 'test', got %v", quintuple.Second)
		}
		if quintuple.Third != true {
			t.Errorf("Third element should be true, got %v", quintuple.Third)
		}
		if quintuple.Fourth != 2.0 {
			t.Errorf("Fourth element should be 2.0, got %v", quintuple.Fourth)
		}
		if !slices.Equal(quintuple.Fifth, []byte{1, 2, 3}) {
			t.Errorf("Fifth element should be [1 2 3], got %v", quintuple.Fifth)
		}
		if quintuple.Len() != 5 {
			t.Errorf("Expected Len to be 5, got %d", quintuple.Len())
		}
		if first, ok := quintuple.Get(0).(int); !ok || first != 1 {
			t.Errorf("First element should be 1, got %v", first)
		}
		if second, ok := quintuple.Get(1).(string); !ok || second != "test" {
			t.Errorf("Second element should be 'test', got %s", second)
		}
		if third, ok := quintuple.Get(2).(bool); !ok || third != true {
			t.Errorf("Third element should be true, got %v", third)
		}
		if fourth, ok := quintuple.Get(3).(float64); !ok || fourth != 2.0 {
			t.Errorf("Fourth element should be 2.0, got %v", fourth)
		}
		if fifth, ok := quintuple.Get(4).([]byte); !ok || !slices.Equal(fifth, []byte{1, 2, 3}) {
			t.Errorf("Fifth element should be [1 2 3], got %v", fifth)
		}
		if v, ok := quintuple.Get(5).(int); v != 0 || ok {
			t.Errorf("Out of bounds element should be a zero value, got %v", v)
		}
	})

	t.Run("NewQuintuple", func(t *testing.T) {
		quintuple := tuple.NewQuintuple(1, "test", true, 2.0, []byte{1, 2, 3})
		if quintuple.First != 1 {
			t.Errorf("First element should be 1, got %v", quintuple.First)
		}
		if quintuple.Second != "test" {
			t.Errorf("Second element should be 'test', got %v", quintuple.Second)
		}
		if quintuple.Third != true {
			t.Errorf("Third element should be true, got %v", quintuple.Third)
		}
		if quintuple.Fourth != 2.0 {
			t.Errorf("Fourth element should be 2.0, got %v", quintuple.Fourth)
		}
		if !slices.Equal(quintuple.Fifth, []byte{1, 2, 3}) {
			t.Errorf("Fifth element should be [1 2 3], got %v", quintuple.Fifth)
		}
		if quintuple.Len() != 5 {
			t.Errorf("Expected Len to be 5, got %d", quintuple.Len())
		}
		if first, ok := quintuple.Get(0).(int); !ok || first != 1 {
			t.Errorf("First element should be 1, got %v", first)
		}
		if second, ok := quintuple.Get(1).(string); !ok || second != "test" {
			t.Errorf("Second element should be 'test', got %s", second)
		}
		if third, ok := quintuple.Get(2).(bool); !ok || third != true {
			t.Errorf("Third element should be true, got %v", third)
		}
		if fourth, ok := quintuple.Get(3).(float64); !ok || fourth != 2.0 {
			t.Errorf("Fourth element should be 2.0, got %v", fourth)
		}
		if fifth, ok := quintuple.Get(4).([]byte); !ok || !slices.Equal(fifth, []byte{1, 2, 3}) {
			t.Errorf("Fifth element should be [1 2 3], got %v", fifth)
		}
		if v, ok := quintuple.Get(5).(int); v != 0 || ok {
			t.Errorf("Out of bounds element should be a zero value, got %v", v)
		}
	})
}
