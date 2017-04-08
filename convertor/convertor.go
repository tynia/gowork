package convertor

import (
	e "gowork/error"
)

func ToInt(v interface{}) (int, *e.WError) {
	if v == nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if i, ok := v.(int); ok {
		return i, nil
	}

	return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToInt8(v interface{}) (int8, *e.WError) {
	if v == nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if i8, ok := v.(int8); ok {
		return i8, nil
	}

	return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToInt16(v interface{}) (int16, *e.WError) {
	if v == nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if i16, ok := v.(int16); ok {
		return i16, nil
	}

	return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToInt32(v interface{}) (int32, *e.WError) {
	if v == nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if i32, ok := v.(int32); ok {
		return i32, nil
	}

	return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToInt64(v interface{}) (int64, *e.WError) {
	if v == nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if i64, ok := v.(int64); ok {
		return i64, nil
	}

	return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToRune(v interface{}) (rune, *e.WError) {
	i, err := ToInt(v)
	if err != nil {
		return rune(i), err
	}

	return rune(i), nil
}

func ToUint(v interface{}) (uint, *e.WError) {
	if v == nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if ui, ok := v.(uint); ok {
		return ui, nil
	}

	return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToUint8(v interface{}) (uint8, *e.WError) {
	if v == nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if ui8, ok := v.(uint8); ok {
		return ui8, nil
	}

	return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToByte(v interface{}) (byte, *e.WError) {
	ui8, err := ToUint8(v)
	if err != nil {
		return byte(ui8), err
	}

	return byte(ui8), nil
}

func ToUint16(v interface{}) (uint16, *e.WError) {
	if v == nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if ui16, ok := v.(uint16); ok {
		return ui16, nil
	}

	return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToUint32(v interface{}) (uint32, *e.WError) {
	if v == nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if ui32, ok := v.(uint32); ok {
		return ui32, nil
	}

	return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToUint64(v interface{}) (uint64, *e.WError) {
	if v == nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if ui64, ok := v.(uint64); ok {
		return ui64, nil
	}

	return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToFloat32(v interface{}) (float32, *e.WError) {
	var f32 float32
	if v == nil {
		return f32, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	var ok bool
	if f32, ok = v.(float32); ok {
		return f32, nil
	}

	return f32, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToFloat64(v interface{}) (float64, *e.WError) {
	var f64 float64
	if v == nil {
		return f64, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	var ok bool
	if f64, ok = v.(float64); ok {
		return f64, nil
	}

	return f64, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToBool(v interface{}) (bool, *e.WError) {
	if v == nil {
		return false, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if b, ok := v.(bool); ok {
		return b, nil
	}

	return false, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToString(v interface{}) (string, *e.WError) {
	if v == nil {
		return "", e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	if str, ok := v.(string); ok {
		return str, nil
	}

	return "", e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToComplex64(v interface{}) (complex64, *e.WError) {
	var comp64 complex64
	if v == nil {
		return comp64, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	var ok bool
	if comp64, ok = v.(complex64); ok {
		return comp64, nil
	}

	return comp64, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}

func ToComplex128(v interface{}) (complex128, *e.WError) {
	var comp128 complex128

	if v == nil {
		return comp128, e.NewWError(e.ERR_CODE_CONVERT_NIL, "Invalid in value[v: %#+v]", v)
	}

	var ok bool
	if comp128, ok = v.(complex128); ok {
		return comp128, nil
	}

	return comp128, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, v: %#+v", v)
}
