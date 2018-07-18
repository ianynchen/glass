package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	pool           *ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	AttributeTable
}

func (attr *CodeAttribute) readInfo(reader *ContentReader) error {
	var err error
	attr.maxStack, err = reader.readUint16()
	if err != nil {
		return err
	}

	attr.maxLocals, err = reader.readUint16()
	if err != nil {
		return err
	}

	var codeLength uint32
	codeLength, err = reader.readUint32()
	if err != nil {
		return err
	}

	attr.code, err = reader.readBytes(int(codeLength))
	if err != nil {
		return err
	}

	attr.exceptionTable, err = readExceptionTable(reader)
	if err != nil {
		return err
	}

	attr.attributes, err = readAttributes(reader, attr.pool)
	return err
}

func (attr *CodeAttribute) MaxStack() uint {
	return uint(attr.maxStack)
}
func (attr *CodeAttribute) MaxLocals() uint {
	return uint(attr.maxLocals)
}
func (attr *CodeAttribute) Code() []byte {
	return attr.code
}
func (attr *CodeAttribute) ExceptionTable() []*ExceptionTableEntry {
	return attr.exceptionTable
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func readExceptionTable(reader *ContentReader) ([]*ExceptionTableEntry, error) {
	var err error
	var exceptionTableLength uint16
	exceptionTableLength, err = reader.readUint16()

	if err != nil {
		return nil, err
	}
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{}
		exceptionTable[i].startPc, err = reader.readUint16()
		exceptionTable[i].endPc, err = reader.readUint16()
		exceptionTable[i].handlerPc, err = reader.readUint16()
		exceptionTable[i].catchType, err = reader.readUint16()
	}
	return exceptionTable, err
}

func (attr *ExceptionTableEntry) StartPc() uint16 {
	return attr.startPc
}
func (attr *ExceptionTableEntry) EndPc() uint16 {
	return attr.endPc
}
func (attr *ExceptionTableEntry) HandlerPc() uint16 {
	return attr.handlerPc
}
func (attr *ExceptionTableEntry) CatchType() uint16 {
	return attr.catchType
}
