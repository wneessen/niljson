// SPDX-FileCopyrightText: 2024 Winni Neessen <wn@neessen.dev>
//
// SPDX-License-Identifier: MIT

package niljson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

// List of common errors
const (
	ErrUnmarshalFailed    = "failed to unmarshal json with nil types: %s"
	ErrMarshalFailed      = "failed to marshal json with nil types: %s"
	ErrExpectedJSONString = "expected json to be %q, got %q"
	ErrExpectedJSONInt    = "expected json to be %d, got %d"
	ErrExpectedValue      = "expected value, but got nil"
	ErrExpectedNil        = "expected nil, but got value"
	ErrExpectedNilReset   = "expected value to be nil after reset"
)

var jsonBytes = []byte(
	`{
		"bool": true,
		"bytes": "Ynl0ZXM=",
		"float32": 1.6,
		"float64": 123.456,
		"int": 123,
		"int64": 12345678901234,
		"nilvalue": null,
		"string": "test",
		"uint": 1,
		"uint8": 2,
		"uint16": 3,
		"uint32": 4,
		"uint64": 5
	}`)

func TestVariable_UnmarshalJSON_Boolean(t *testing.T) {
	type JSONType struct {
		Value    NilBoolean `json:"bool"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if !jt.Value.Value() {
		t.Errorf("expected value to be true, got %t", jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_Boolean(t *testing.T) {
	type JSONType struct {
		Value    NilBoolean `json:"bool"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"bool":false,"nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable(false),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_ByteSlice(t *testing.T) {
	type JSONType struct {
		Value    NilByteSlice `json:"bytes"`
		NilValue NilBoolean   `json:"nilvalue,omitempty"`
	}

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if !bytes.Equal(jt.Value.Value(), []byte("bytes")) {
		t.Errorf("expected value to be %q, got %q", "bytes", jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_ByteSlice(t *testing.T) {
	type JSONType struct {
		Value    NilByteSlice `json:"bytes"`
		NilValue NilBoolean   `json:"nilvalue,omitempty"`
	}

	expected := `{"bytes":"Ynl0ZXM=","nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable([]byte("bytes")),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_Float32(t *testing.T) {
	type JSONType struct {
		Value    NilFloat32 `json:"float32"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}
	expected := float32(1.6)

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if jt.Value.Value() != expected {
		t.Errorf("expected value to be %f, got %f", expected, jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_Float32(t *testing.T) {
	type JSONType struct {
		Value    NilFloat32 `json:"float32"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"float32":1.234,"nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable(float32(1.234)),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_Float64(t *testing.T) {
	type JSONType struct {
		Value    NilFloat64 `json:"float64"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}
	expected := 123.456

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if jt.Value.Value() != expected {
		t.Errorf("expected value to be %f, got %f", expected, jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_Float64(t *testing.T) {
	type JSONType struct {
		Value    NilFloat64 `json:"float64"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"float64":123.456,"nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable(123.456),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_Int(t *testing.T) {
	type JSONType struct {
		Value    NilInt     `json:"int"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}
	expected := 123

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if jt.Value.Value() != expected {
		t.Errorf(ErrExpectedJSONInt, expected, jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_Int(t *testing.T) {
	type JSONType struct {
		Value    NilInt     `json:"int"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"int":123,"nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable(123),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_Int64(t *testing.T) {
	type JSONType struct {
		Value    NilInt64   `json:"int64"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}
	expected := int64(12345678901234)

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if jt.Value.Value() != expected {
		t.Errorf(ErrExpectedJSONInt, expected, jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_Int64(t *testing.T) {
	type JSONType struct {
		Value    NilInt64   `json:"int64"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"int64":12345678901234,"nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable(int64(12345678901234)),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_String(t *testing.T) {
	type JSONType struct {
		Value    NilString  `json:"string"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}
	expected := "test"

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if jt.Value.Value() != expected {
		t.Errorf("expected value to be %s, got %s", expected, jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_String(t *testing.T) {
	type JSONType struct {
		Value    NilString  `json:"string"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"string":"test123","nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable("test123"),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_UInt(t *testing.T) {
	type JSONType struct {
		Value    NilUInt    `json:"uint"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}
	expected := uint(1)

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if jt.Value.Value() != expected {
		t.Errorf(ErrExpectedJSONInt, expected, jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_UInt(t *testing.T) {
	type JSONType struct {
		Value    NilUInt    `json:"uint"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"uint":123,"nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable(uint(123)),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_UInt8(t *testing.T) {
	type JSONType struct {
		Value    NilUInt8   `json:"uint8"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}
	expected := uint8(2)

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if jt.Value.Value() != expected {
		t.Errorf(ErrExpectedJSONInt, expected, jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_UInt8(t *testing.T) {
	type JSONType struct {
		Value    NilUInt8   `json:"uint8"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"uint8":1,"nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable(uint8(1)),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_UInt16(t *testing.T) {
	type JSONType struct {
		Value    NilUInt16  `json:"uint16"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}
	expected := uint16(3)

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if jt.Value.Value() != expected {
		t.Errorf(ErrExpectedJSONInt, expected, jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_UInt16(t *testing.T) {
	type JSONType struct {
		Value    NilUInt16  `json:"uint16"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"uint16":2,"nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable(uint16(2)),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_UInt32(t *testing.T) {
	type JSONType struct {
		Value    NilUInt32  `json:"uint32"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}
	expected := uint32(4)

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if jt.Value.Value() != expected {
		t.Errorf(ErrExpectedJSONInt, expected, jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_UInt32(t *testing.T) {
	type JSONType struct {
		Value    NilUInt32  `json:"uint32"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"uint32":3,"nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable(uint32(3)),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func TestVariable_UnmarshalJSON_UInt64(t *testing.T) {
	type JSONType struct {
		Value    NilUInt64  `json:"uint64"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}
	expected := uint64(5)

	var jt JSONType
	if err := json.Unmarshal(jsonBytes, &jt); err != nil {
		t.Errorf(ErrUnmarshalFailed, err)
	}

	if jt.Value.IsNil() {
		t.Errorf(ErrExpectedValue)
	}
	if jt.NilValue.NotNil() {
		t.Errorf(ErrExpectedNil)
	}
	if jt.Value.Value() != expected {
		t.Errorf(ErrExpectedJSONInt, expected, jt.Value.Value())
	}

	jt.Value.Reset()
	if jt.Value.NotNil() {
		t.Errorf(ErrExpectedNilReset)
	}
}

func TestVariable_MarshalJSON_UInt64(t *testing.T) {
	type JSONType struct {
		Value    NilUInt64  `json:"uint64"`
		NilValue NilBoolean `json:"nilvalue,omitempty"`
	}

	expected := `{"uint64":4,"nilvalue":null}`
	jt := &JSONType{
		Value: NewVariable(uint64(4)),
	}
	data, err := json.Marshal(&jt)
	if err != nil {
		t.Errorf(ErrMarshalFailed, err)
	}
	if !bytes.Equal(data, []byte(expected)) {
		t.Errorf(ErrExpectedJSONString, expected, string(data))
	}
}

func ExampleVariable_UnmarshalJSON() {
	type JSONType struct {
		Bool       NilBoolean   `json:"bool"`
		ByteSlice  NilByteSlice `json:"bytes"`
		Float32    NilFloat32   `json:"float32,omitempty"`
		Float64    NilFloat64   `json:"float64"`
		Int        NilInt       `json:"int"`
		Int64      NilInt64     `json:"int64"`
		NullString NilString    `json:"nilvalue,omitempty"`
		String     NilString    `json:"string"`
	}
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
	// Output: Bool is: true, Float 32 is nil, Float 64 is: 0.000000, String is: test
}

func ExampleVariable_MarshalJSON() {
	type JSONType struct {
		Bool       NilBoolean   `json:"bool"`
		ByteSlice  NilByteSlice `json:"bytes"`
		Float32    NilFloat32   `json:"float32,omitempty"`
		Float64    NilFloat64   `json:"float64"`
		Int        NilInt       `json:"int"`
		Int64      NilInt64     `json:"int64"`
		NullString NilString    `json:"nilvalue,omitempty"`
		String     NilString    `json:"string"`
		UInt       NilUInt      `json:"uint"`
		UInt8      NilUInt8     `json:"uint8"`
	}

	example := &JSONType{
		Bool:      NewVariable(false),
		ByteSlice: NewVariable([]byte("bytes")),
		Float64:   NewVariable(123.456),
		Int:       NewVariable(123),
		Int64:     NewVariable(int64(12345678901234)),
		String:    NewVariable("test"),
		UInt:      NewVariable(uint(123)),
	}
	data, err := json.Marshal(example)
	if err != nil {
		fmt.Printf("failed to marshal JSON: %s", err)
		os.Exit(1)
	}
	fmt.Println(string(data))
	// Output: {"bool":false,"bytes":"Ynl0ZXM=","float32":null,"float64":123.456,"int":123,"int64":12345678901234,"nilvalue":null,"string":"test","uint":123,"uint8":null}
}
