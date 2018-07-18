package classfile

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

func (attr *LocalVariableTableAttribute) readInfo(reader *ContentReader) error {
	localVariableTableLength, err := reader.readUint16()
	if err != nil {
		return err
	}
	attr.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range attr.localVariableTable {
		attr.localVariableTable[i] = &LocalVariableTableEntry{}
		attr.localVariableTable[i].startPc, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.localVariableTable[i].length, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.localVariableTable[i].nameIndex, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.localVariableTable[i].descriptorIndex, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.localVariableTable[i].index, err = reader.readUint16()
	}
	return err
}
