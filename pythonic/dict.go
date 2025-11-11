package pythonic

import (
	"errors"
	"reflect"
)

// Dict is a Python-like generic dictionary implementation.
type Dict[K comparable, V any] struct {
	m map[K]V
}

// Pair represents a key/value pair returned by Items().
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// NewDictFromMap constructs a Dict by copying entries from provided map.
func NewDictFromMap[K comparable, V any](src map[K]V) *Dict[K, V] {
	d := &Dict[K, V]{m: make(map[K]V, len(src))}
	for k, v := range src {
		d.m[k] = v
	}
	return d
}

// NewDict creates an empty Dict with optional capacity.
func NewDict[K comparable, V any](cap int) *Dict[K, V] {
	if cap <= 0 {
		return &Dict[K, V]{m: make(map[K]V)}
	}
	return &Dict[K, V]{m: make(map[K]V, cap)}
}

// Len returns number of key/value pairs.
func (d *Dict[K, V]) Len() int {
	if d == nil || d.m == nil {
		return 0
	}
	return len(d.m)
}

// Contains returns true if the key exists.
func (d *Dict[K, V]) Contains(key K) bool {
	if d == nil || d.m == nil {
		return false
	}
	_, ok := d.m[key]
	return ok
}

// Get returns value for key or error if missing.
func (d *Dict[K, V]) Get(key K) (V, error) {
	var zero V
	if d == nil || d.m == nil {
		return zero, errors.New("key not found")
	}
	v, ok := d.m[key]
	if !ok {
		return zero, errors.New("key not found")
	}
	return v, nil
}

// Set assigns value for key.
func (d *Dict[K, V]) Set(key K, value V) {
	if d.m == nil {
		d.m = make(map[K]V)
	}
	d.m[key] = value
}

// Remove deletes key and returns error if key not present.
func (d *Dict[K, V]) Remove(key K) error {
	if d == nil || d.m == nil {
		return errors.New("key not found")
	}
	if _, ok := d.m[key]; !ok {
		return errors.New("key not found")
	}
	delete(d.m, key)
	return nil
}

// Pop removes key and returns its value. Error if key missing.
func (d *Dict[K, V]) Pop(key K) (V, error) {
	var zero V
	if d == nil || d.m == nil {
		return zero, errors.New("key not found")
	}
	v, ok := d.m[key]
	if !ok {
		return zero, errors.New("key not found")
	}
	delete(d.m, key)
	return v, nil
}

// PopItem removes and returns an arbitrary key/value pair. Error if empty.
func (d *Dict[K, V]) PopItem() (K, V, error) {
	var zk K
	var zv V
	if d == nil || len(d.m) == 0 {
		return zk, zv, errors.New("popitem from empty dict")
	}
	for k, v := range d.m {
		delete(d.m, k)
		return k, v, nil
	}
	return zk, zv, errors.New("popitem from empty dict")
}

// Keys returns a slice with all keys (iteration order unspecified).
func (d *Dict[K, V]) Keys() []K {
	out := make([]K, 0, d.Len())
	if d == nil || d.m == nil {
		return out
	}
	for k := range d.m {
		out = append(out, k)
	}
	return out
}

// Values returns a slice with all values (iteration order unspecified).
func (d *Dict[K, V]) Values() []V {
	out := make([]V, 0, d.Len())
	if d == nil || d.m == nil {
		return out
	}
	for _, v := range d.m {
		out = append(out, v)
	}
	return out
}

// Items returns a slice with all key/value pairs.
func (d *Dict[K, V]) Items() []Pair[K, V] {
	out := make([]Pair[K, V], 0, d.Len())
	if d == nil || d.m == nil {
		return out
	}
	for k, v := range d.m {
		out = append(out, Pair[K, V]{Key: k, Value: v})
	}
	return out
}

// Clear removes all entries.
func (d *Dict[K, V]) Clear() { d.m = make(map[K]V) }

// Update merges entries from other into this dict (overwrites existing keys).
func (d *Dict[K, V]) Update(other *Dict[K, V]) {
	if other == nil || other.m == nil {
		return
	}
	if d.m == nil {
		d.m = make(map[K]V, len(other.m))
	}
	for k, v := range other.m {
		d.m[k] = v
	}
}

// Copy returns a shallow copy of the dict.
func (d *Dict[K, V]) Copy() *Dict[K, V] {
	if d == nil || d.m == nil {
		return NewDict[K, V](0)
	}
	return NewDictFromMap(d.m)
}

// SetDefault returns the value for key if present; otherwise sets key to defaultValue and returns it.
func (d *Dict[K, V]) SetDefault(key K, defaultValue V) V {
	if d == nil {
		d = NewDict[K, V](0)
	}
	if d.m == nil {
		d.m = make(map[K]V)
	}
	if v, ok := d.m[key]; ok {
		return v
	}
	d.m[key] = defaultValue
	return defaultValue
}

// Equal returns true if two dicts have identical key/value mappings.
func (d *Dict[K, V]) Equal(other *Dict[K, V]) bool {
	if d == nil && other == nil {
		return true
	}
	if d == nil || other == nil {
		return false
	}
	if d.Len() != other.Len() {
		return false
	}
	for k, v := range d.m {
		ov, ok := other.m[k]
		if !ok {
			return false
		}
		if !reflect.DeepEqual(v, ov) {
			return false
		}
	}
	return true
}
