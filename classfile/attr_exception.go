package classfile

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (attr *ExceptionsAttribute) readInfo(reader *ContentReader) error {
	var err error
	attr.exceptionIndexTable, err = reader.readUint16s()
	return err
}

func (attr *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return attr.exceptionIndexTable
}
