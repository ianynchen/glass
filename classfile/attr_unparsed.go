package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (attr *UnparsedAttribute) readInfo(reader *ContentReader) error {
	var err error
	attr.info, err = reader.readBytes(int(attr.length))
	return err
}

func (attr *UnparsedAttribute) Info() []byte {
	return attr.info
}
