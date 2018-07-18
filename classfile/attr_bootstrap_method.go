package classfile

/*
BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];
    } bootstrap_methods[num_bootstrap_methods];
}
*/
type BootstrapMethodsAttribute struct {
	bootstrapMethods []*BootstrapMethod
}

type BootstrapMethod struct {
	bootstrapMethodRef uint16
	bootstrapArguments []uint16
}

func (attr *BootstrapMethodsAttribute) readInfo(reader *ContentReader) error {
	var err error
	var numBootstrapMethods uint16
	numBootstrapMethods, err = reader.readUint16()

	if err != nil {
		return err
	}
	attr.bootstrapMethods = make([]*BootstrapMethod, numBootstrapMethods)
	for i := range attr.bootstrapMethods {
		attr.bootstrapMethods[i] = &BootstrapMethod{}
		attr.bootstrapMethods[i].bootstrapMethodRef, err = reader.readUint16()
		attr.bootstrapMethods[i].bootstrapArguments, err = reader.readUint16s()

		if err != nil {
			break
		}
	}
	return err
}
