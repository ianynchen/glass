package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
	pool            *ConstantPool
	sourceFileIndex uint16
}

func (attr *SourceFileAttribute) readInfo(reader *ContentReader) error {
	var err error
	attr.sourceFileIndex, err = reader.readUint16()
	return err
}

func (attr *SourceFileAttribute) FileName() string {
	return attr.pool.getUtf8(attr.sourceFileIndex)
}
