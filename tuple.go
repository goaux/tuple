// Package tuple provides extremely simple generic tuple types
// and utility functions for creating tuples.
package tuple

// Tuple is a generic interface that represents a tuple of elements.
type Tuple interface {
	// Len returns the number of elements in the tuple.
	Len() int

	// Get returns an element at index i from the tuple.
	// If the index is out of bounds, it returns nil.
	Get(int) any
}

// Pair is a generic type for a pair of values.
type Pair[T0, T1 any] struct {
	First  T0
	Second T1
}

// MakePair returns a new [Pair] with the given values.
func MakePair[T0, T1 any](first T0, second T1) Pair[T0, T1] {
	return Pair[T0, T1]{First: first, Second: second}
}

// NewPair returns a pointer to a new [Pair] with the given values.
func NewPair[T0, T1 any](first T0, second T1) *Pair[T0, T1] {
	return &Pair[T0, T1]{First: first, Second: second}
}

// Len returns 2.
func (Pair[_, _]) Len() int { return 2 }

// Get returns the value at index i.
// If the index is out of bounds, it returns nil.
func (v Pair[_, _]) Get(i int) any {
	switch i {
	case 0:
		return v.First
	case 1:
		return v.Second
	}
	return nil
}

// Unpack returns the elements of the pair as a tuple.
func (v Pair[T0, T1]) Unpack() (T0, T1) {
	return v.First, v.Second
}

// Triple is a generic type for a triple of values.
type Triple[T0, T1, T2 any] struct {
	First  T0
	Second T1
	Third  T2
}

// MakeTriple returns a new [Triple] with the given values.
func MakeTriple[T0, T1, T2 any](first T0, second T1, third T2) Triple[T0, T1, T2] {
	return Triple[T0, T1, T2]{First: first, Second: second, Third: third}
}

// NewTriple returns a pointer to a new [Triple] with the given values.
func NewTriple[T0, T1, T2 any](first T0, second T1, third T2) *Triple[T0, T1, T2] {
	return &Triple[T0, T1, T2]{First: first, Second: second, Third: third}
}

// Len returns 3.
func (t Triple[_, _, _]) Len() int { return 3 }

// Get returns the value at index i.
// If the index is out of bounds, it returns nil.
func (t Triple[_, _, _]) Get(i int) any {
	switch i {
	case 0:
		return t.First
	case 1:
		return t.Second
	case 2:
		return t.Third
	}
	return nil
}

// Unpack returns the elements of the triple as a tuple.
func (v Triple[T0, T1, T2]) Unpack() (T0, T1, T2) {
	return v.First, v.Second, v.Third
}

// Quadruple is a generic type for a quadruple of values.
type Quadruple[T0, T1, T2, T3 any] struct {
	First  T0
	Second T1
	Third  T2
	Fourth T3
}

// MakeQuadruple returns a new [Quadruple] with the given values.
func MakeQuadruple[T0, T1, T2, T3 any](first T0, second T1, third T2, fourth T3) Quadruple[T0, T1, T2, T3] {
	return Quadruple[T0, T1, T2, T3]{First: first, Second: second, Third: third, Fourth: fourth}
}

// NewQuadruple returns a pointer to a new [Quadruple] with the given values.
func NewQuadruple[T0, T1, T2, T3 any](first T0, second T1, third T2, fourth T3) *Quadruple[T0, T1, T2, T3] {
	return &Quadruple[T0, T1, T2, T3]{First: first, Second: second, Third: third, Fourth: fourth}
}

// Len returns 4.
func (q Quadruple[_, _, _, _]) Len() int { return 4 }

// Get returns the value at index i.
// If the index is out of bounds, it returns nil.
func (q Quadruple[_, _, _, _]) Get(i int) any {
	switch i {
	case 0:
		return q.First
	case 1:
		return q.Second
	case 2:
		return q.Third
	case 3:
		return q.Fourth
	}
	return nil
}

// Unpack returns the elements of the quadruple as a tuple.
func (v Quadruple[T0, T1, T2, T3]) Unpack() (T0, T1, T2, T3) {
	return v.First, v.Second, v.Third, v.Fourth
}

// Quintuple is a generic type for a quintuple of values.
type Quintuple[T0, T1, T2, T3, T4 any] struct {
	First  T0
	Second T1
	Third  T2
	Fourth T3
	Fifth  T4
}

// MakeQuintuple returns a new [Quintuple] with the given values.
func MakeQuintuple[T0, T1, T2, T3, T4 any](first T0, second T1, third T2, fourth T3, fifth T4) Quintuple[T0, T1, T2, T3, T4] {
	return Quintuple[T0, T1, T2, T3, T4]{First: first, Second: second, Third: third, Fourth: fourth, Fifth: fifth}
}

// NewQuintuple returns a pointer to a new [Quintuple] with the given values.
func NewQuintuple[T0, T1, T2, T3, T4 any](first T0, second T1, third T2, fourth T3, fifth T4) *Quintuple[T0, T1, T2, T3, T4] {
	return &Quintuple[T0, T1, T2, T3, T4]{First: first, Second: second, Third: third, Fourth: fourth, Fifth: fifth}
}

// Len returns 5.
func (q Quintuple[_, _, _, _, _]) Len() int { return 5 }

// Get returns the value at index i.
// If the index is out of bounds, it returns nil.
func (q Quintuple[_, _, _, _, _]) Get(i int) any {
	switch i {
	case 0:
		return q.First
	case 1:
		return q.Second
	case 2:
		return q.Third
	case 3:
		return q.Fourth
	case 4:
		return q.Fifth
	}
	return nil
}

// Unpack returns the elements of the quintuple as a tuple.
func (v Quintuple[T0, T1, T2, T3, T4]) Unpack() (T0, T1, T2, T3, T4) {
	return v.First, v.Second, v.Third, v.Fourth, v.Fifth
}
