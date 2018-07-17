package classfile

import (
	"errors"
	"fmt"
)

type ClassParser interface {
	Parse() (*ClassFile, error)
}

type ClassFileParser struct {
	reader        *ContentReader
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
	class := &ClassFile{}

	parser.parse(parseMagic, class)
	parser.parse(parseVersion, class)

	return class, parser.previousError
}

func (parser *ClassFileParser) parse(parseClass parserFunc, class *ClassFile) {
	if parser.previousError == nil {
		err := parseClass(parser.reader, class)
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
	var offset uint64
	offset = 2

	count, err := reader.readUint16()
	class.constantPool = new(ConstantPool)
	class.constantPool.info = make([]ConstantInfo, int(count))

	if err != nil {
		return err
	}

	for i := 1; i < int(count); i++ {
		class.constantPool.info[i], err = readConstantInfo(reader, class.constantPool)

		if err != nil {
			return err
		}

		switch class.constantPool.info[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return nil
}

func readConstantInfo(reader *ContentReader, pool *ConstantPool) (ConstantInfo, error) {
	tag, err := reader.readUint8()

	if err == nil {
		c := newConstantInfo(tag, pool)
		c.readInfo(reader)
		return c, nil
	}
	return nil, err
}

func newConstantInfo(tag uint8, cp *ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{cp: cp}
	default:
		return fmt.Errorf("Bad constant pool tag: %v", tag)
	}
}

func parseClassFlags(reader *ContentReader, class *ClassFile) (uint64, error) {
	var err error
	class.accessFlags, err = reader.readUint16()
	return 0, err
}

func parseThisClass(reader *ContentReader, class *ClassFile, content []byte) (uint64, error) {
	var err error
	class.thisClass, err = reader.readUint16()
	return 0, err
}

func parseSuperClass(reader *ContentReader, class *ClassFile, content []byte) (uint64, error) {
	var err error
	class.superClass, err = reader.readUint16()
	return 0, err
}
