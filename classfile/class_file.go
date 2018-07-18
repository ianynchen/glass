package classfile

// ClassFile represents a java class file structure
type ClassFile struct {
	//magic           uint32
	minorVersion uint16
	majorVersion uint16
	constantPool *ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	AttributeTable
}
