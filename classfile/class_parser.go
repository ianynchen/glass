package classfile

import (
	"errors"
)

type ClassParser interface {
	Parse() (*ClassFile, error)
}

type ClassFileParser struct {
	reader        *ContentReader
	class         *ClassFile
	previousError error
}

type parserFunc func(reader *ContentReader, class *ClassFile) error

// NewParser creates a new class file reader
func NewParser(content []byte) ClassParser {
	parser := new(ClassFileParser)
	parser.reader = NewContentReader(content)
	return parser
}

// Parse parses content of a class file and tries to return a
// parsed class file or error if failure
func (parser *ClassFileParser) Parse() (*ClassFile, error) {
	parser.class = &ClassFile{}

	parser.parse(parseMagic)
	parser.parse(parseVersion)
	parser.parse(parseConstantPool)
	parser.parse(parseAccessFlags)
	parser.parse(parseThisClass)
	parser.parse(parseSuperClass)
	parser.parse(parseInterfaces)
	parser.parse(parseMembers)
	parser.parse(parseAttributes)

	return parser.class, parser.previousError
}

func (parser *ClassFileParser) parse(parseFunction parserFunc) {
	if parser.previousError == nil {
		err := parseFunction(parser.reader, parser.class)
		parser.previousError = err
	}
}

func parseMagic(reader *ContentReader, class *ClassFile) error {
	magic, err := reader.readUint32()

	if err != nil {
		return err
	} else if magic != 0xCAFEBABE {
		return errors.New("invalid magic")
	}
	return nil
}

func parseVersion(reader *ContentReader, class *ClassFile) error {
	var err error
	class.minorVersion, err = reader.readUint16()

	if err != nil {
		return err
	}

	class.majorVersion, err = reader.readUint16()

	if err != nil {
		return err
	}
	return nil
}

func parseConstantPool(reader *ContentReader, class *ClassFile) error {
	class.constantPool = new(ConstantPool)
	err := class.constantPool.read(reader)
	return err
}

func parseAccessFlags(reader *ContentReader, class *ClassFile) error {
	var err error
	class.accessFlags, err = reader.readUint16()
	return err
}

func parseClassFlags(reader *ContentReader, class *ClassFile) error {
	var err error
	class.accessFlags, err = reader.readUint16()
	return err
}

func parseThisClass(reader *ContentReader, class *ClassFile) error {
	var err error
	class.thisClass, err = reader.readUint16()
	return err
}

func parseSuperClass(reader *ContentReader, class *ClassFile) error {
	var err error
	class.superClass, err = reader.readUint16()
	return err
}

func parseInterfaces(reader *ContentReader, class *ClassFile) error {
	var err error
	class.interfaces, err = reader.readUint16s()
	return err
}

func parseMembers(reader *ContentReader, class *ClassFile) error {
	var err error
	class.fields, err = readMembers(reader, class.constantPool)
	class.methods, err = readMembers(reader, class.constantPool)
	return err
}

func parseAttributes(reader *ContentReader, class *ClassFile) error {
	var err error
	class.attributes, err = readAttributes(reader, class.constantPool)
	return err
}
