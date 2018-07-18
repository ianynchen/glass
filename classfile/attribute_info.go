package classfile

var (
	_attrDeprecated = &DeprecatedAttribute{}
	_attrSynthetic  = &SyntheticAttribute{}
)

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *ContentReader) error
}

func readAttributes(reader *ContentReader, pool *ConstantPool) ([]AttributeInfo, error) {
	var err error
	attributesCount, err := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i], err = readAttribute(reader, pool)
		if err != nil {
			break
		}
	}
	return attributes, err
}

func readAttribute(reader *ContentReader, pool *ConstantPool) (AttributeInfo, error) {
	var err error
	var attrNameIndex uint16
	var attrLen uint32
	var attrName string
	attrNameIndex, err = reader.readUint16()
	attrLen, err = reader.readUint32()
	attrName, err = pool.getUtf8(attrNameIndex)
	attrInfo := newAttributeInfo(attrName, pool)
	if attrInfo == nil {
		attrInfo = &UnparsedAttribute{
			name:   attrName,
			length: attrLen,
		}
	}

	err = attrInfo.readInfo(reader)
	return attrInfo, err
}

func newAttributeInfo(attrName string, pool *ConstantPool) AttributeInfo {
	switch attrName {
	// case "AnnotationDefault":
	case "BootstrapMethods":
		return &BootstrapMethodsAttribute{}
	case "Code":
		return &CodeAttribute{pool: pool}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return _attrDeprecated
	case "EnclosingMethod":
		return &EnclosingMethodAttribute{pool: pool}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "LocalVariableTypeTable":
		return &LocalVariableTypeTableAttribute{}
	// case "MethodParameters":
	// case "RuntimeInvisibleAnnotations":
	// case "RuntimeInvisibleParameterAnnotations":
	// case "RuntimeInvisibleTypeAnnotations":
	// case "RuntimeVisibleAnnotations":
	// case "RuntimeVisibleParameterAnnotations":
	// case "RuntimeVisibleTypeAnnotations":
	case "Signature":
		return &SignatureAttribute{pool: pool}
	case "SourceFile":
		return &SourceFileAttribute{pool: pool}
	// case "SourceDebugExtension":
	// case "StackMapTable":
	case "Synthetic":
		return _attrSynthetic
	default:
		return nil // undefined attr
	}
}
