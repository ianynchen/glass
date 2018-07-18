package classfile

type AttributeTable struct {
	attributes []AttributeInfo
}

/* group 1 */
func (table *AttributeTable) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range table.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (table *AttributeTable) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range table.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (table *AttributeTable) ExceptionsAttribute() *ExceptionsAttribute {
	for _, attrInfo := range table.attributes {
		switch attrInfo.(type) {
		case *ExceptionsAttribute:
			return attrInfo.(*ExceptionsAttribute)
		}
	}
	return nil
}

func (table *AttributeTable) BootstrapMethodsAttribute() *BootstrapMethodsAttribute {
	for _, attrInfo := range table.attributes {
		switch attrInfo.(type) {
		case *BootstrapMethodsAttribute:
			return attrInfo.(*BootstrapMethodsAttribute)
		}
	}
	return nil
}

/* group 2 */

func (table *AttributeTable) EnclosingMethodAttribute() *EnclosingMethodAttribute {
	for _, attrInfo := range table.attributes {
		switch attrInfo.(type) {
		case *EnclosingMethodAttribute:
			return attrInfo.(*EnclosingMethodAttribute)
		}
	}
	return nil
}

func (table *AttributeTable) SignatureAttribute() *SignatureAttribute {
	for _, attrInfo := range table.attributes {
		switch attrInfo.(type) {
		case *SignatureAttribute:
			return attrInfo.(*SignatureAttribute)
		}
	}
	return nil
}

/* group 3 */

func (table *AttributeTable) SourceFileAttribute() *SourceFileAttribute {
	for _, attrInfo := range table.attributes {
		switch attrInfo.(type) {
		case *SourceFileAttribute:
			return attrInfo.(*SourceFileAttribute)
		}
	}
	return nil
}

func (table *AttributeTable) LineNumberTableAttribute() *LineNumberTableAttribute {
	for _, attrInfo := range table.attributes {
		switch attrInfo.(type) {
		case *LineNumberTableAttribute:
			return attrInfo.(*LineNumberTableAttribute)
		}
	}
	return nil
}

/* unparsed */

func (table *AttributeTable) RuntimeVisibleAnnotationsAttributeData() []byte {
	return table.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (table *AttributeTable) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return table.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (table *AttributeTable) AnnotationDefaultAttributeData() []byte {
	return table.getUnparsedAttributeData("AnnotationDefault")
}

func (table *AttributeTable) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range table.attributes {
		switch attrInfo.(type) {
		case *UnparsedAttribute:
			unparsedAttr := attrInfo.(*UnparsedAttribute)
			if unparsedAttr.name == name {
				return unparsedAttr.info
			}
		}
	}
	return nil
}
