// SPDX-FileCopyrightText: 2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package niljson

import (
	"encoding/json"
	"testing"
)

func TestVariable_UnmarshalJSON(t *testing.T) {
	jsonBytes := []byte(`{"string":"test", "int":123, "int64": 12345678901234, "float":123.456, "nil":null, "bool":true}`)
	type JSONType struct {
		Bool       NilBoolean `json:"bool"`
		Float64    NilFloat64 `json:"float"`
		Int        NilInt     `json:"int"`
		Int64      NilInt64   `json:"int64"`
		NullString NilString  `json:"nil"`
		String     NilString  `json:"string"`
	}
	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf("failed to unmarshal json with nil types: %v", err)
	}
	if jt.Float64.IsNil() {
		t.Errorf("expected not nil float64")
	}
	if jt.Int.IsNil() {
		t.Errorf("expected not nil int")
	}
	if jt.Int64.IsNil() {
		t.Errorf("expected not nil int64")
	}
	if jt.String.IsNil() {
		t.Errorf("expected not nil string")
	}
	if jt.NullString.NotNil() {
		t.Errorf("expected nil string")
	}
	if !jt.Bool.Get() {
		t.Errorf("expected bool to be true, got %t", jt.Bool.Get())
	}
	if jt.Float64.Get() != 123.456 {
		t.Errorf("expected float64 to be 123.456, got %f", jt.Float64.Get())
	}
	if jt.Int.Get() != 123 {
		t.Errorf("expected int to be 123, got %d", jt.Int.Get())
	}
	if jt.Int64.Get() != 12345678901234 {
		t.Errorf("expected int to be 12345678901234, got %d", jt.Int64.Get())
	}
	if jt.String.Get() != "test" {
		t.Errorf("expected string to be 'test', got %s", jt.String.Get())
	}

	jt.Bool.Reset()
	if jt.Bool.NotNil() {
		t.Errorf("expected bool to be nil after reset")
	}
	jt.Float64.Reset()
	if jt.Float64.NotNil() {
		t.Errorf("expected float64 to be nil after reset")
	}
	jt.Int.Reset()
	if jt.Int.NotNil() {
		t.Errorf("expected int to be nil after reset")
	}
	jt.Int64.Reset()
	if jt.Int64.NotNil() {
		t.Errorf("expected int64 to be nil after reset")
	}
	jt.String.Reset()
	if jt.String.NotNil() {
		t.Errorf("expected string to be nil after reset")
	}
}
