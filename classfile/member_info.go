package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

// MemberInfo represents either a field or method in a class
type MemberInfo struct {
	pool            *ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	AttributeTable
}

func readMembers(reader *ContentReader, pool *ConstantPool) ([]*MemberInfo, error) {
	var err error
	var count uint16
	count, err = reader.readUint16()

	if err != nil {
		return nil, err
	}

	members := make([]*MemberInfo, count)
	for i := range members {
		members[i] = &MemberInfo{pool: pool}
		err = members[i].read(reader)

		if err != nil {
			break
		}
	}
	return members, err
}

func (member *MemberInfo) read(reader *ContentReader) error {
	var err error
	member.accessFlags, err = reader.readUint16()
	if err != nil {
		return err
	}
	member.nameIndex, err = reader.readUint16()
	if err != nil {
		return err
	}
	member.descriptorIndex, err = reader.readUint16()
	if err != nil {
		return err
	}
	member.attributes, err = readAttributes(reader, member.pool)
	return err
}

func (member *MemberInfo) AccessFlags() uint16 {
	return member.accessFlags
}

func (member *MemberInfo) Name() string {
	return member.pool.getUtf8(member.nameIndex)
}

func (member *MemberInfo) Descriptor() string {
	return member.pool.getUtf8(member.descriptorIndex)
}

func (member *MemberInfo) Signature() string {
	signatureAttr := member.SignatureAttribute()
	if signatureAttr != nil {
		return signatureAttr.Signature()
	}
	return ""
}
