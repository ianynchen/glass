package classfile

import (
	"errors"
	"fmt"

	"github.com/ianynchen/glass/util"
)

type ClassParser interface {
	Parse() (*ClassFile, error)
}

type ClassFileParser struct {
	content       []byte
	offset        uint64
	previousError error
}

type parserFunc func(class *ClassFile, content []byte) (uint64, error)

// NewParser creates a new class file reader
func NewParser(content []byte) *ClassParser {
	return &ClassFileParser(content, 0)
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
		offset, err := parseClass(class, parser.content[parser.offset:])
		parser.offset += offset
		parser.previousError = err
	}
}

func parseMagic(class *ClassFile, content []byte) (uint64, error) {
	magic, err := util.ParseUint32(content)

	if err != nil {
		return 0, err
	} else if magic != 0xCAFEBABE {
		return 0, errors.New("invalid magic")
	}
	return 4, nil
}

func parseVersion(class *ClassFile, content []byte) (uint64, error) {
	var err error
	class.minorVersion, err = util.ParseUint16(content)

	if err != nil {
		return 0, err
	}

	class.majorVersion, err = util.ParseUint16(content[2:])

	if err != nil {
		return 0, err
	}
	return 4, nil
}

func parseConstantPool(class *ClassFile, content []byte) (uint64, error) {
	var offset uint64
	offset = 2

	count, err := util.ParseUint16(content)
	class.constantPool = new(ConstantPool)
	class.constantPool.info = make([]ConstantInfo, int(count))

	if err != nil {
		return offset, err
	}

	for i := 1; i < int(count); i++ {
		class.constantPool.info[i] = readConstantInfo(content[offset:], class.constantPool)
		switch class.constantPool.info[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
}

func readConstantInfo(content []byte, pool *ConstantPool) ConstantInfo {
	tag := util.ParseUint8(content)
	c := newConstantInfo(tag, pool)
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
	default: // todo
		panic(fmt.Errorf("BAD constant pool tag: %v", tag))
		return nil
	}
}

func parseClassFlags(class *ClassFile, content []byte) (uint64, error) {
	var err error
	class.accessFlags, err = util.ParseUint16(content)
	return 2, err
}

func parseThisClass(class *ClassFile, content []byte) (uint64, error) {
	var err error
	class.thisClass, err = util.ParseUint16(content)
	return 2, err
}

func parseSuperClass(class *ClassFile, content []byte) (uint64, error) {
	var err error
	class.superClass, err = util.ParseUint16(content)
	return 2, err
}
