package classfile

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/
type SignatureAttribute struct {
	pool           *ConstantPool
	signatureIndex uint16
}

func (attr *SignatureAttribute) readInfo(reader *ContentReader) error {
	var err error
	attr.signatureIndex, err = reader.readUint16()
	return err
}

func (attr *SignatureAttribute) Signature() string {
	return attr.pool.getUtf8(attr.signatureIndex)
}
