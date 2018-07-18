package classfile

type MarkerAttribute struct{}

func (attr *MarkerAttribute) readInfo(reader *ContentReader) error {
	// read nothing
	return nil
}

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}
