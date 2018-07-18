package classfile

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (attr *LineNumberTableAttribute) readInfo(reader *ContentReader) error {
	lineNumberTableLength, err := reader.readUint16()

	if err != nil {
		return err
	}
	attr.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range attr.lineNumberTable {
		attr.lineNumberTable[i] = &LineNumberTableEntry{}
		attr.lineNumberTable[i].startPc, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.lineNumberTable[i].lineNumber, err = reader.readUint16()
	}
	return err
}

func (attr *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for i := len(attr.lineNumberTable) - 1; i >= 0; i-- {
		entry := attr.lineNumberTable[i]
		if pc >= int(entry.startPc) {
			return int(entry.lineNumber)
		}
	}
	return -1
}
