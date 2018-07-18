package classfile

/*
LocalVariableTypeTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_type_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 signature_index;
        u2 index;
    } local_variable_type_table[local_variable_type_table_length];
}
*/
type LocalVariableTypeTableAttribute struct {
	localVariableTypeTable []*LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
	startPc        uint16
	length         uint16
	nameIndex      uint16
	signatureIndex uint16
	index          uint16
}

func (attr *LocalVariableTypeTableAttribute) readInfo(reader *ContentReader) error {
	localVariableTypeTableLength, err := reader.readUint16()
	if err != nil {
		return err
	}
	attr.localVariableTypeTable = make([]*LocalVariableTypeTableEntry, localVariableTypeTableLength)
	for i := range attr.localVariableTypeTable {
		attr.localVariableTypeTable[i] = &LocalVariableTypeTableEntry{}
		attr.localVariableTypeTable[i].startPc, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.localVariableTypeTable[i].length, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.localVariableTypeTable[i].nameIndex, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.localVariableTypeTable[i].signatureIndex, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.localVariableTypeTable[i].index, err = reader.readUint16()
	}
	return err
}
