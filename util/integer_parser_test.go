package util

import (
	"testing"

	"github.com/issue9/assert"
)

func TestParseUint8(t *testing.T) {

	content := []byte{0xCA}
	val, err := ParseUint8(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint8(0xCA))
}

func BenchmarkParseUint8(b *testing.B) {
	content := []byte{0xCA}
	for i := 0; i < b.N; i++ {
		ParseUint8(content)
	}
}

func TestParseUint16(t *testing.T) {

	content := []byte{0xCA, 0xFE}
	val, err := ParseUint16(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint16(0xCAFE))
}

func BenchmarkParseUint16(b *testing.B) {
	content := []byte{0xCA, 0xFE}
	for i := 0; i < b.N; i++ {
		ParseUint16(content)
	}
}

func TestParseUint32(t *testing.T) {

	content := []byte{0xCA, 0xFE, 0xBA, 0xBE}
	val, err := ParseUint32(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint32(0xCAFEBABE))
}

func BenchmarkParseUint32(b *testing.B) {
	content := []byte{0xCA, 0xFE, 0xBA, 0xBE}
	for i := 0; i < b.N; i++ {
		ParseUint32(content)
	}
}

func TestParseUint64(t *testing.T) {

	content := []byte{0xCA, 0xFE, 0xBA, 0xBE, 0x00, 0x00, 0x00, 0x00}
	val, err := ParseUint64(content)

	assert.True(t, err == nil)
	assert.Equal(t, val, uint64(0xCAFEBABE00000000))
}

func BenchmarkParseUint64(b *testing.B) {
	content := []byte{0xCA, 0xFE, 0xBA, 0xBE, 0x00, 0x00, 0x00, 0x00}
	for i := 0; i < b.N; i++ {
		ParseUint64(content)
	}
}
