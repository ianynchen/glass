package classfile

/*
EnclosingMethod_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 class_index;
    u2 method_index;
}
*/
type EnclosingMethodAttribute struct {
	pool        *ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (attr *EnclosingMethodAttribute) readInfo(reader *ContentReader) error {
	var err error
	attr.classIndex, err = reader.readUint16()
	if err != nil {
		return err
	}
	attr.methodIndex, err = reader.readUint16()
	return err
}

func (attr *EnclosingMethodAttribute) ClassName() string {
	return attr.pool.getClassName(attr.classIndex)
}

func (attr *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if attr.methodIndex > 0 {
		return attr.pool.getNameAndType(attr.methodIndex)
	}
	return "", ""
}
