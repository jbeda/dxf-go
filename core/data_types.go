// Package core provides functions and data structures for basic DXF operations.
package core

import "strconv"

// DataType is an interface for a DXF type.
type DataType interface {

	// ToString returns a string representation of the value
	ToString() string

	// Value returns the encapsulated value as an interface{} the caller should cast it appropriately
	// or use the 'AsString', 'AsInt' and 'AsFloat' functions
	Value() interface{}
}

// String DataType implementation
type String struct {
	value string
}

// NewString creates a new String object with the provided value.
func NewString(value string) (DataType, error) {
	returnValue := new(String)
	returnValue.value = value
	return returnValue, nil
}

// ToString returns a string representation of the value
func (s String) ToString() string {
	return s.value
}

// Value returns the encapsulated value
func (s String) Value() interface{} {
	return s.value
}

// Integer DataType implementation
type Integer struct {
	value int
}

// NewInteger creates a new Integer object with the provided value as a string
func NewInteger(value string) (DataType, error) {
	returnValue := new(Integer)
	v, err := strconv.Atoi(value)
	returnValue.value = v
	return returnValue, err
}

// ToString returns a string representation of the value
func (i Integer) ToString() string {
	return strconv.Itoa(i.value)
}

// Value returns the encapsulated value
func (i Integer) Value() interface{} {
	return i.value
}

// Float DataType implementation
type Float struct {
	value float64
}

// NewFloat creates a new Float object with the provided value as a string
func NewFloat(value string) (DataType, error) {
	returnValue := new(Float)
	v, err := strconv.ParseFloat(value, 64)
	returnValue.value = v
	return returnValue, err
}

// ToString returns a string representation of the value
func (f Float) ToString() string {
	return strconv.FormatFloat(f.value, 'f', -1, 64)
}

// Value returns the encapsulated value
func (f Float) Value() interface{} {
	return f.value
}

// AsString is the acessor for a String DataType.
// If d is String, it will return the (value, true), otherwise ("", false)
func AsString(d DataType) (string, bool) {
	value, ok := d.Value().(string)
	return value, ok
}

// AsInt is the acessor for an Integer DataType.
// If d is Integer, it will return the (value, true), otherwise (0, false)
func AsInt(d DataType) (int, bool) {
	value, ok := d.Value().(int)
	return value, ok
}

// AsFloat is the acessor for a Float DataType.
// If d is Float, it will return the (value, true), otherwise (0.0, false)
func AsFloat(d DataType) (float64, bool) {
	value, ok := d.Value().(float64)
	return value, ok
}