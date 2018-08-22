// Code generated by thriftrw v1.13.0. DO NOT EDIT.
// @generated

package enum_conflict

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/thriftrw/gen/internal/tests/enums"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"math"
	"strconv"
	"strings"
)

type RecordType int32

const (
	RecordTypeName  RecordType = 0
	RecordTypeEmail RecordType = 1
)

// RecordType_Values returns all recognized values of RecordType.
func RecordType_Values() []RecordType {
	return []RecordType{
		RecordTypeName,
		RecordTypeEmail,
	}
}

// UnmarshalText tries to decode RecordType from a byte slice
// containing its name.
//
//   var v RecordType
//   err := v.UnmarshalText([]byte("Name"))
func (v *RecordType) UnmarshalText(value []byte) error {
	switch s := string(value); s {
	case "Name":
		*v = RecordTypeName
		return nil
	case "Email":
		*v = RecordTypeEmail
		return nil
	default:
		val, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return fmt.Errorf("unknown enum value %q for %q: %v", s, "RecordType", err)
		}
		*v = RecordType(val)
		return nil
	}
}

// MarshalText encodes RecordType to text.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements the TextMarshaler interface.
func (v RecordType) MarshalText() ([]byte, error) {
	switch int32(v) {
	case 0:
		return []byte("Name"), nil
	case 1:
		return []byte("Email"), nil
	}
	return []byte(strconv.FormatInt(int64(v), 10)), nil
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of RecordType.
// Enums are logged as objects, where the value is logged with key "value", and
// if this value's name is known, the name is logged with key "name".
func (v RecordType) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt32("value", int32(v))
	switch int32(v) {
	case 0:
		enc.AddString("name", "Name")
	case 1:
		enc.AddString("name", "Email")
	}
	return nil
}

// Ptr returns a pointer to this enum value.
func (v RecordType) Ptr() *RecordType {
	return &v
}

// ToWire translates RecordType into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// Enums are represented as 32-bit integers over the wire.
func (v RecordType) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

// FromWire deserializes RecordType from its Thrift-level
// representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TI32)
//   if err != nil {
//     return RecordType(0), err
//   }
//
//   var v RecordType
//   if err := v.FromWire(x); err != nil {
//     return RecordType(0), err
//   }
//   return v, nil
func (v *RecordType) FromWire(w wire.Value) error {
	*v = (RecordType)(w.GetI32())
	return nil
}

// String returns a readable string representation of RecordType.
func (v RecordType) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "Name"
	case 1:
		return "Email"
	}
	return fmt.Sprintf("RecordType(%d)", w)
}

// Equals returns true if this RecordType value matches the provided
// value.
func (v RecordType) Equals(rhs RecordType) bool {
	return v == rhs
}

// MarshalJSON serializes RecordType into JSON.
//
// If the enum value is recognized, its name is returned. Otherwise,
// its integer value is returned.
//
// This implements json.Marshaler.
func (v RecordType) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"Name\""), nil
	case 1:
		return ([]byte)("\"Email\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

// UnmarshalJSON attempts to decode RecordType from its JSON
// representation.
//
// This implementation supports both, numeric and string inputs. If a
// string is provided, it must be a known enum name.
//
// This implements json.Unmarshaler.
func (v *RecordType) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}

	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "RecordType")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "RecordType")
		}
		*v = (RecordType)(x)
		return nil
	case string:
		return v.UnmarshalText([]byte(w))
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "RecordType")
	}
}

type Records struct {
	RecordType      *RecordType       `json:"recordType,omitempty"`
	OtherRecordType *enums.RecordType `json:"otherRecordType,omitempty"`
}

func _RecordType_ptr(v RecordType) *RecordType {
	return &v
}

func _RecordType_1_ptr(v enums.RecordType) *enums.RecordType {
	return &v
}

// ToWire translates a Records struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Records) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.RecordType == nil {
		v.RecordType = _RecordType_ptr(DefaultRecordType)
	}
	{
		w, err = v.RecordType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.OtherRecordType == nil {
		v.OtherRecordType = _RecordType_1_ptr(DefaultOtherRecordType)
	}
	{
		w, err = v.OtherRecordType.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _RecordType_Read(w wire.Value) (RecordType, error) {
	var v RecordType
	err := v.FromWire(w)
	return v, err
}

func _RecordType_1_Read(w wire.Value) (enums.RecordType, error) {
	var v enums.RecordType
	err := v.FromWire(w)
	return v, err
}

// FromWire deserializes a Records struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Records struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Records
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Records) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x RecordType
				x, err = _RecordType_Read(field.Value)
				v.RecordType = &x
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TI32 {
				var x enums.RecordType
				x, err = _RecordType_1_Read(field.Value)
				v.OtherRecordType = &x
				if err != nil {
					return err
				}

			}
		}
	}

	if v.RecordType == nil {
		v.RecordType = _RecordType_ptr(DefaultRecordType)
	}

	if v.OtherRecordType == nil {
		v.OtherRecordType = _RecordType_1_ptr(DefaultOtherRecordType)
	}

	return nil
}

// String returns a readable string representation of a Records
// struct.
func (v *Records) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	if v.RecordType != nil {
		fields[i] = fmt.Sprintf("RecordType: %v", *(v.RecordType))
		i++
	}
	if v.OtherRecordType != nil {
		fields[i] = fmt.Sprintf("OtherRecordType: %v", *(v.OtherRecordType))
		i++
	}

	return fmt.Sprintf("Records{%v}", strings.Join(fields[:i], ", "))
}

func _RecordType_EqualsPtr(lhs, rhs *RecordType) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

func _RecordType_1_EqualsPtr(lhs, rhs *enums.RecordType) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Records match the
// provided Records.
//
// This function performs a deep comparison.
func (v *Records) Equals(rhs *Records) bool {
	if !_RecordType_EqualsPtr(v.RecordType, rhs.RecordType) {
		return false
	}
	if !_RecordType_1_EqualsPtr(v.OtherRecordType, rhs.OtherRecordType) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Records.
func (v *Records) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if v.RecordType != nil {
		if err := enc.AddObject("recordType", *v.RecordType); err != nil {
			return err
		}
	}
	if v.OtherRecordType != nil {
		if err := enc.AddObject("otherRecordType", *v.OtherRecordType); err != nil {
			return err
		}
	}
	return nil
}

// GetRecordType returns the value of RecordType if it is set or its
// default value if it is unset.
func (v *Records) GetRecordType() (o RecordType) {
	if v.RecordType != nil {
		return *v.RecordType
	}
	o = DefaultRecordType
	return
}

// GetOtherRecordType returns the value of OtherRecordType if it is set or its
// default value if it is unset.
func (v *Records) GetOtherRecordType() (o enums.RecordType) {
	if v.OtherRecordType != nil {
		return *v.OtherRecordType
	}
	o = DefaultOtherRecordType
	return
}
