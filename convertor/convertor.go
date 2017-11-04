package convertor

import (
	e "gowork/error"
	"strconv"
)

func ToInt(v string) (int, *e.WError) {
	i64, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return int(i64), nil
}

func ToInt8(v string) (int8, *e.WError) {
	i64, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return int8(i64), nil
}

func ToInt16(v string) (int16, *e.WError) {
	i64, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return int16(i64), nil
}

func ToInt32(v string) (int32, *e.WError) {
	i64, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return int32(i64), nil
}

func ToInt64(v string) (int64, *e.WError) {
	i64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return i64, nil
}

func ToRune(v string) (rune, *e.WError) {
	i64, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return rune(i64), nil
}

func ToUint(v string) (uint, *e.WError) {
	ui64, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return uint(ui64), nil
}

func ToUint8(v string) (uint8, *e.WError) {
	ui64, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return uint8(ui64), nil
}

func ToByte(v string) (byte, *e.WError) {
	ui64, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return uint8(ui64), nil
}

func ToUint16(v string) (uint16, *e.WError) {
	ui64, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return uint16(ui64), nil
}

func ToUint32(v string) (uint32, *e.WError) {
	ui64, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return uint32(ui64), nil
}

func ToUint64(v string) (uint64, *e.WError) {
	ui64, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return ui64, nil
}

func ToFloat32(v string) (float32, *e.WError) {
	f64, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return float32(f64), nil
}

func ToFloat64(v string) (float64, *e.WError) {
	f64, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "type mismatch, value: \"%s\"", v)
	}

	return f64, nil
}

func ToBool(v string) (bool, *e.WError) {
	ok, err := strconv.ParseBool(v)
	if err != nil {
		return false, e.NewWError(e.ERR_CODE_CONVERT_TYPE, "unknown mismatch, value: \"%s\"", v)
	}

	return ok, nil
}