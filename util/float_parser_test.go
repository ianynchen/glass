package util

import (
	"math"
	"testing"

	"github.com/issue9/assert"
)

func TestParseFloat32Normal(t *testing.T) {

	numerical, err := ParseFloat32([]byte{0xC0, 0xD9, 0x99, 0x9A})
	assert.IsNil(err)
	assert.True(t, numerical.FieldRange == IsNumber)
	assert.True(t, math.Abs(float64(numerical.FieldValue.(float32)+float32(6.800000190734863))) < 1e-5)
}

func BenchmarkParseFloat32Normal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFloat32([]byte{0xC0, 0xD9, 0x99, 0x9A})
	}
}

func TestParseFloat64Normal(t *testing.T) {

	numerical, err := ParseFloat64([]byte{0x3f, 0xeb, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33})
	assert.IsNil(err)
	assert.True(t, numerical.FieldRange == IsNumber)
	assert.True(t, math.Abs(numerical.FieldValue.(float64)-0.85) < 1e-5)
}

func BenchmarkParseFloat64Normal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseFloat64([]byte{0x3f, 0xeb, 0x33, 0x33, 0x33, 0x33, 0x33, 0x33})
	}
}
