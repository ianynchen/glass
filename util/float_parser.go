package util

import "math"

const (
	// PositiveInfinity indicates value is positively infinite
	PositiveInfinity int = iota
	// NegativeInfinity indicates value is negatively infinite.
	NegativeInfinity
	// NaN not a number
	NaN
	// IsNumber is numerical value
	IsNumber
)

// Numerical can be either an integer or a floating point number
// The values could be expressed in numerical values, or special values
// such as +/- infinity, NaN
type Numerical struct {
	FieldRange int
	FieldValue interface{}
}

// ParseFloat32 parse a []byte into numerical value as a 32 bit float number
func ParseFloat32(content []byte) (Numerical, error) {

	numerical := Numerical{}
	bits, err := ParseUint32(content)

	if err != nil {
		return numerical, err
	}

	if bits == 0x7f800000 {
		numerical.FieldRange = PositiveInfinity
	} else if bits == 0xff800000 {
		numerical.FieldRange = NegativeInfinity
	} else if (bits >= 0x7f800001 && bits <= 0x7fffffff) ||
		(bits >= 0xff800001 && bits <= 0xffffffff) {
		numerical.FieldRange = NaN
	} else {
		var s int
		if (bits >> 31) == 0 {
			s = 1
		} else {
			s = -1
		}
		e := int((bits >> 23) & 0xff)

		var m int
		if e == 0 {
			m = int((bits & 0x7fffff) << 1)
		} else {
			m = int((bits & 0x7fffff) | 0x800000)
		}
		numerical.FieldRange = IsNumber
		numerical.FieldValue = float32(float64(s) * float64(m) * math.Pow(2, float64(e-150)))
	}
	return numerical, nil
}

// ParseFloat64 converts []byte to Numerical, where FieldValue is of type float64
func ParseFloat64(content []byte) (Numerical, error) {

	numerical := Numerical{}
	bits, err := ParseUint64(content)

	if err != nil {
		return numerical, err
	}
	if bits == 0x7ff0000000000000 {
		numerical.FieldRange = PositiveInfinity
	} else if bits == 0xfff0000000000000 {
		numerical.FieldRange = NegativeInfinity
	} else if (bits >= 0x7ff0000000000001 && bits <= 0x7fffffffffffffff) ||
		(bits >= 0xfff0000000000001 && bits <= 0xffffffffffffffff) {
		numerical.FieldRange = NaN
	} else {
		var s int
		if (bits >> 63) == 0 {
			s = 1
		} else {
			s = -1
		}
		e := int((bits >> 52) & 0x7ff)

		var m int64
		if e == 0 {
			m = int64((bits & 0xfffffffffffff) << 1)
		} else {
			m = int64((bits & 0xfffffffffffff) | 0x10000000000000)
		}
		numerical.FieldRange = IsNumber
		numerical.FieldValue = float64(s) * float64(m) * math.Pow(2, float64(e-1075))
	}
	return numerical, nil
}
