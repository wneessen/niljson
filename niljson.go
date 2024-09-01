// SPDX-FileCopyrightText: 2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package niljson

import (
	"encoding/json"
)

// Variable is a generic variable type that can be null.
type Variable[T any] struct {
	value  T
	notNil bool
}

// Value returns the value of the Variable
func (v *Variable[T]) Value() T {
	return v.value
}

// Get is an alias method for Value()
func (v *Variable[T]) Get() T {
	return v.Value()
}

// NotNil returns true when a Variable is not nil
func (v *Variable[T]) NotNil() bool {
	return v.notNil
}

// IsNil returns true when a Variable is nil
func (v *Variable[T]) IsNil() bool {
	return !v.notNil
}

// Reset resets the value to the Variable to a zero value and sets it to be nil
func (v *Variable[T]) Reset() {
	var newVal T
	v.value = newVal
	v.notNil = false
}

// NilBoolean is an boolean type that can be nil
type NilBoolean = Variable[bool]

// NilByteSlice is a []byte type that can be nil
type NilByteSlice = Variable[[]byte]

// NilInt is an int type that can be nil
type NilInt = Variable[int]

// NilInt64 is an int64 type that can be nil
type NilInt64 = Variable[int64]

// NilFloat32 is an float32 type that can be nil
type NilFloat32 = Variable[float32]

// NilFloat64 is an float64 type that can be nil
type NilFloat64 = Variable[float64]

// NilString is a string type that can be nil
type NilString = Variable[string]

// UnmarshalJSON interprets the generic Nil types and sets the value and notnil of the type
func (v *Variable[T]) UnmarshalJSON(data []byte) error {
	if string(data) != "null" {
		v.value = *new(T)
		v.notNil = true
		return json.Unmarshal(data, &v.value)
	}
	return nil
}
