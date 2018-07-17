package classfile

import (
	"encoding/binary"
	"errors"
	"math"
)

type ContentReader struct {
	content []byte
	offset  int
	length  int
}

func NewContentReader(content []byte) *ContentReader {
	reader := ContentReader{content}
	reader.length = len(content)
	return &reader
}

func (reader *ContentReader) readBytes(size int) ([]byte, error) {
	if reader.length-reader.offset >= size {
		oldOffset := reader.offset
		reader.offset += size
		return reader.content[oldOffset:reader.offset], nil
	}
	return nil, errors.New("content not long enough")
}

func (reader *ContentReader) readUint8() (uint8, error) {
	if reader.length-reader.offset >= 1 {
		val := reader.content[reader.offset]
		reader.offset++
		return val, nil
	}
	return 0, errors.New("content not long enough")
}

func (reader *ContentReader) readUint16() (uint16, error) {
	bytes, err := reader.readBytes(2)
	if err == nil {
		return binary.BigEndian.Uint16(bytes), nil
	}
	return 0, err
}

func (reader *ContentReader) readUint32() (uint32, error) {
	bytes, err := reader.readBytes(4)
	if err == nil {
		return binary.BigEndian.Uint32(bytes), nil
	}
	return 0, err
}

func (reader *ContentReader) readUint64() (uint32, error) {
	bytes, err := reader.readBytes(8)
	if err == nil {
		return binary.BigEndian.Uint64(bytes), nil
	}
	return 0, err
}

func (reader *ContentReader) readFloat() (float32, error) {
	bytes, err := reader.readUint32()
	if err == nil {
		return math.Float32frombits(bytes), nil
	}
	return 0, err
}

func (reader *ContentReader) readDouble() (float64, error) {
	bytes, err := reader.readUint64()
	if err == nil {
		return math.Float64frombits(bytes), nil
	}
	return 0, err
}
