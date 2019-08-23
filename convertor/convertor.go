package convertor

import (
	"github.com/tynia/gowork/xerr"
	"strconv"
	"strings"
)

func ToInt(v string) (int, error) {
	i64, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return int(i64), nil
}

func ToInt8(v string) (int8, error) {
	i64, err := strconv.ParseInt(v, 10, 8)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return int8(i64), nil
}

func ToInt16(v string) (int16, error) {
	i64, err := strconv.ParseInt(v, 10, 16)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return int16(i64), nil
}

func ToInt32(v string) (int32, error) {
	i64, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return int32(i64), nil
}

func ToInt64(v string) (int64, error) {
	i64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return i64, nil
}

func ToRune(v string) (rune, error) {
	i64, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return rune(i64), nil
}

func ToUint(v string) (uint, error) {
	ui64, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return uint(ui64), nil
}

func ToUint8(v string) (uint8, error) {
	ui64, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return uint8(ui64), nil
}

func ToByte(v string) (byte, error) {
	ui64, err := strconv.ParseUint(v, 10, 8)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return uint8(ui64), nil
}

func ToUint16(v string) (uint16, error) {
	ui64, err := strconv.ParseUint(v, 10, 16)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return uint16(ui64), nil
}

func ToUint32(v string) (uint32, error) {
	ui64, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return uint32(ui64), nil
}

func ToUint64(v string) (uint64, error) {
	ui64, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return ui64, nil
}

func ToFloat32(v string) (float32, error) {
	f64, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return float32(f64), nil
}

func ToFloat64(v string) (float64, error) {
	f64, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, xerr.New(2000, "type mismatch, value: \"%s\"", v)
	}

	return f64, nil
}

func ToBool(v string) (bool, error) {
	ok, err := strconv.ParseBool(v)
	if err != nil {
		return false, xerr.New(2000, "unknown mismatch, value: \"%s\"", v)
	}

	return ok, nil
}

func ToArray(v string, sep string) []string {
	return strings.Split(v, sep)
}