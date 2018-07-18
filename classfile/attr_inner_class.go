package classfile

/*
InnerClasses_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_classes;
    {   u2 inner_class_info_index;
        u2 outer_class_info_index;
        u2 inner_name_index;
        u2 inner_class_access_flags;
    } classes[number_of_classes];
}
*/
type InnerClassesAttribute struct {
	classes []*InnerClassInfo
}

type InnerClassInfo struct {
	innerClassInfoIndex   uint16
	outerClassInfoIndex   uint16
	innerNameIndex        uint16
	innerClassAccessFlags uint16
}

func (attr *InnerClassesAttribute) readInfo(reader *ContentReader) error {
	numberOfClasses, err := reader.readUint16()
	if err != nil {
		return err
	}
	attr.classes = make([]*InnerClassInfo, numberOfClasses)
	for i := range attr.classes {
		attr.classes[i] = &InnerClassInfo{}
		attr.classes[i].innerClassInfoIndex, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.classes[i].outerClassInfoIndex, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.classes[i].innerNameIndex, err = reader.readUint16()
		if err != nil {
			return err
		}
		attr.classes[i].innerClassAccessFlags, err = reader.readUint16()
	}
	return err
}
