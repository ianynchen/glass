package util

import "errors"

// ByteArrayNotLongEnough is the error message to generate when the input byte array
// is not long enough for integer types to generate.
const ByteArrayNotLongEnough string = "byte array not long enough"

// ParseUint8 is used to generate a uint8 from a byte array.
func ParseUint8(content []byte) (uint8, error) {

	if len(content) < 1 {
		return uint8(0), errors.New(ByteArrayNotLongEnough)
	}

	return uint8(content[0]), nil
}

// ParseUint16 is used to generate a uint16 from a byte array.
// Bigendian encoding is assumed.
func ParseUint16(content []byte) (uint16, error) {

	if len(content) < 2 {
		return uint16(0), errors.New(ByteArrayNotLongEnough)
	}

	return (uint16(content[0]) << 8) | (uint16(content[1])), nil
}

// ParseUint32 is used to generate a uint32 from a byte array.
// Bigendian encoding is assumed.
func ParseUint32(content []byte) (uint32, error) {

	if len(content) < 4 {
		return uint32(0), errors.New(ByteArrayNotLongEnough)
	}

	higher, err1 := ParseUint16(content)
	lower, err2 := ParseUint16(content[2:])

	if err1 != nil {
		return uint32(0), err1
	} else if err2 != nil {
		return uint32(0), err2
	}

	return (uint32(higher) << 16) | uint32(lower), nil
}

// ParseUint64 is used to generate a uint64 from a byte array.
// Bigendian encoding is assumed.
func ParseUint64(content []byte) (uint64, error) {

	if len(content) < 8 {
		return uint64(0), errors.New(ByteArrayNotLongEnough)
	}

	higher, err1 := ParseUint32(content)
	lower, err2 := ParseUint32(content[4:])

	if err1 != nil {
		return uint64(0), err1
	} else if err2 != nil {
		return uint64(0), err2
	}

	return (uint64(higher) << 32) | uint64(lower), nil
}
