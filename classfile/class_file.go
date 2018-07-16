package classfile

import (
	"errors"

	"github.com/ianynchen/glass/util"
)

// ClassFile represents a java class file structure
type ClassFile struct {
	//magic           uint32
	minorVersion uint16
	majorVersion uint16
	constantPool *ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	AttributeTable
}

type parseFunc func(class *ClassFile)

// ClassReader reads content of class file and produces *ClassFile
type ClassReader interface {
	Read(content []byte) (*ClassFile, error)
}

func NewClassReader(major uint16, minor uint16) ClassReader {
	return &ClassFileReader{major, minor, nil, nil}
}

type ClassFileReader struct {
	major         uint16
	minor         uint16
	index         uint64
	content       []byte
	previousError error
}

func (reader *ClassFileReader) parse(parser parseFunc) {
	if reader.previousError == nil {
		parser()
	}
}

func (reader *ClassFileReader) Read(content []byte) (*ClassFile, error) {
	reader.content = content
	reader.previousError = nil

	class := &ClassFile{}

	reader.parse(reader.readAndCheckMagic)
	reader.parse(reader.readAndCheckVersion)

	return nil, self.previousError
}

func (reader *ClassFileReader) readAndCheckMagic(class *ClassFile) {
	magic, err := util.ParseUint32(self.content)
	reader.index = 4
	if err == nil {
		if magic != 0xCAFEBABE {
			reader.previousError = errors.New("Invalid magic number")
		}
		return
	}
	reader.previousError = err
}

func (reader *ClassFileReader) readAndCheckVersion(class *ClassFile) {
	major, err := util.ParseUint16()
	reader.index += 2

	if err != nil {
		reader.previousError = err
		return
	}

	minor, err := util.ParseUint16()
	reader.index += 2

	if err != nil {
		reader.previousError = err
		return
	}

	class.majorVersion = major
	class.minorVersion = minor
}
