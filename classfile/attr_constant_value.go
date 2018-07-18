package classfile

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (attr *ConstantValueAttribute) readInfo(reader *ContentReader) error {
	var err error
	attr.constantValueIndex, err = reader.readUint16()
	return err
}

func (attr *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return attr.constantValueIndex
}
