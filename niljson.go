// SPDX-FileCopyrightText: 2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package niljson

import (
	"encoding/json"
)

// NilBoolean is an boolean type that can be nil
type NilBoolean = Variable[bool]

// NilByteSlice is a []byte type that can be nil
type NilByteSlice = Variable[[]byte]

// NilInt is an int type that can be nil
type NilInt = Variable[int]

// NilInt64 is an int64 type that can be nil
type NilInt64 = Variable[int64]

// NilUInt is an uint type that can be nil
type NilUInt = Variable[uint]

// NilUInt8 is an uint8 type that can be nil
type NilUInt8 = Variable[uint8]

// NilUInt16 is an uint16 type that can be nil
type NilUInt16 = Variable[uint16]

// NilUInt32 is an uint32 type that can be nil
type NilUInt32 = Variable[uint32]

// NilUInt64 is an uint64 type that can be nil
type NilUInt64 = Variable[uint64]

// NilFloat32 is an float32 type that can be nil
type NilFloat32 = Variable[float32]

// NilFloat64 is an float64 type that can be nil
type NilFloat64 = Variable[float64]

// NilString is a string type that can be nil
type NilString = Variable[string]

// Variable is a generic variable type that can be null.
type Variable[T any] struct {
	value   T
	notNil  bool
	present bool
}

// NewVariable returns a new Variable of generic type
func NewVariable[T any](value T) Variable[T] {
	return Variable[T]{
		notNil: true,
		value:  value,
	}
}

// IsNil returns true when a Variable is nil
func (v *Variable[T]) IsNil() bool {
	return !v.notNil
}

// NotNil returns true when a Variable is not nil
func (v *Variable[T]) NotNil() bool {
	return v.notNil
}

// Omitted returns true if a value was omitted in the JSON
func (v *Variable[T]) Omitted() bool {
	return !v.present
}

// Reset resets the value to the Variable to a zero value and sets it to be nil
func (v *Variable[T]) Reset() {
	var newVal T
	v.value = newVal
	v.notNil = false
}

// Value returns the value of the Variable
func (v *Variable[T]) Value() T {
	return v.value
}

// Set makes the Variable valid with the given value
func (v *Variable[T]) Set(val T) {
	v.value = val
	v.present = true
	v.notNil = true
}

// MarshalJSON satisfies the json.Marshaler interface for generic Variable types
func (v *Variable[T]) MarshalJSON() ([]byte, error) {
	if !v.notNil {
		return json.Marshal(nil)
	}
	return json.Marshal(v.value)
}

// UnmarshalJSON satisfies the json.Unmarshaler interface for generic Variable types
func (v *Variable[T]) UnmarshalJSON(data []byte) error {
	v.present = true
	if string(data) != "null" {
		v.value = *new(T)
		v.notNil = true
		return json.Unmarshal(data, &v.value)
	}
	return nil
}
