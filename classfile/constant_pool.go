package classfile

import "fmt"

type ConstantPool struct {
	class *ClassFile
	info  []ConstantInfo
}

func (pool *ConstantPool) Infos() []ConstantInfo {
	return pool.info
}

func (pool *ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	info := pool.info[index]
	if info == nil {
		panic(fmt.Errorf("Bad constant pool index: %v!", index))
	}

	return info
}

func (pool *ConstantPool) getNameAndType(index uint16) (name, _type string) {
	ntInfo := pool.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name = pool.getUtf8(ntInfo.nameIndex)
	_type = pool.getUtf8(ntInfo.descriptorIndex)
	return
}

func (pool *ConstantPool) getClassName(index uint16) string {
	classInfo := pool.getConstantInfo(index).(*ConstantClassInfo)
	return pool.getUtf8(classInfo.nameIndex)
}

func (pool *ConstantPool) getUtf8(index uint16) string {
	utf8Info := pool.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

func (pool *ConstantPool) read(reader *ContentReader) error {
	count, err := reader.readUint16()

	if err != nil {
		return err
	}
	pool.info = make([]ConstantInfo, int(count))

	for i := 1; i < int(count); i++ {
		pool.info[i], err = readConstantInfo(reader, pool)

		if err != nil {
			return err
		}

		switch pool.info[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
}

func readConstantInfo(reader *ContentReader, pool *ConstantPool) (ConstantInfo, error) {
	tag, err := reader.readUint8()

	if err == nil {
		c := newConstantInfo(tag, pool)
		err = c.readInfo(reader)
		return c, err
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
