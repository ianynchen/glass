package classfile

import (
	"fmt"
	"unicode/utf16"
)

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

/*
cp_info {
    u1 tag;
    u1 info[];
}
*/
type ConstantInfo interface {
	readInfo(reader *ContentReader) error
}

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
	pool      *ConstantPool
	nameIndex uint16
}

func (info *ConstantClassInfo) readInfo(reader *ContentReader) error {
	var err error
	info.nameIndex, err = reader.readUint16()
	return err
}

func (info *ConstantClassInfo) Name() string {
	return info.pool.getUtf8(info.nameIndex)
}

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantIntegerInfo struct {
	val int32
}

func (info *ConstantIntegerInfo) readInfo(reader *ContentReader) error {
	var err error
	info.val, err = reader.readInt32()
	return err
}

func (info *ConstantIntegerInfo) Value() int32 {
	return info.val
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
type ConstantFloatInfo struct {
	val float32
}

func (info *ConstantFloatInfo) readInfo(reader *ContentReader) error {
	var err error
	info.val, err = reader.readFloat()
	return err
}

func (info *ConstantFloatInfo) Value() float32 {
	return info.val
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantLongInfo struct {
	val int64
}

func (info *ConstantLongInfo) readInfo(reader *ContentReader) error {
	var err error
	info.val, err = reader.readInt64()
	return err
}

func (info *ConstantLongInfo) Value() int64 {
	return info.val
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantDoubleInfo struct {
	val float64
}

func (info *ConstantDoubleInfo) readInfo(reader *ContentReader) error {
	var err error
	info.val, err = reader.readDouble()
	return err
}

func (info *ConstantDoubleInfo) Value() float64 {
	return info.val
}

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (info *ConstantNameAndTypeInfo) readInfo(reader *ContentReader) error {
	var err error
	info.nameIndex, err = reader.readUint16()
	info.descriptorIndex, err = reader.readUint16()
	return err
}

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	pool        *ConstantPool
	stringIndex uint16
}

func (info *ConstantStringInfo) readInfo(reader *ContentReader) error {
	var err error
	info.stringIndex, err = reader.readUint16()
	return err
}

func (info *ConstantStringInfo) String() string {
	return info.pool.getUtf8(info.stringIndex)
}

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ContentReader) error {
	length, err := reader.readUint16()
	bytes, err := reader.readBytes(int(length))
	self.str = decodeMUTF8(bytes)
	return err
}

func (info *ConstantUtf8Info) Str() string {
	return info.str
}

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}

/*
CONSTANT_Fieldref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

/*
CONSTANT_Methodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

/*
CONSTANT_InterfaceMethodref_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMemberrefInfo struct {
	pool             *ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (info *ConstantMemberrefInfo) readInfo(reader *ContentReader) error {
	var err error
	info.classIndex, err = reader.readUint16()
	info.nameAndTypeIndex, err = reader.readUint16()
	return err
}

func (info *ConstantMemberrefInfo) ClassName() string {
	return info.pool.getClassName(info.classIndex)
}

func (info *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return info.pool.getNameAndType(info.nameAndTypeIndex)
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	pool                     *ConstantPool
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (info *ConstantInvokeDynamicInfo) readInfo(reader *ContentReader) error {
	var err error
	info.bootstrapMethodAttrIndex, err = reader.readUint16()
	info.nameAndTypeIndex, err = reader.readUint16()
	return err
}

func (info *ConstantInvokeDynamicInfo) NameAndType() (string, string) {
	return info.pool.getNameAndType(info.nameAndTypeIndex)
}

func (info *ConstantInvokeDynamicInfo) BootstrapMethodInfo() (uint16, []uint16) {
	bmAttr := info.pool.class.BootstrapMethodsAttribute()
	bm := bmAttr.bootstrapMethods[info.bootstrapMethodAttrIndex]

	return bm.bootstrapMethodRef, bm.bootstrapArguments
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (info *ConstantMethodHandleInfo) readInfo(reader *ContentReader) error {
	var err error
	info.referenceKind, err = reader.readUint8()
	info.referenceIndex, err = reader.readUint16()
	return err
}

func (info *ConstantMethodHandleInfo) ReferenceKind() uint8 {
	return info.referenceKind
}

func (info *ConstantMethodHandleInfo) ReferenceIndex() uint16 {
	return info.referenceIndex
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (info *ConstantMethodTypeInfo) readInfo(reader *ContentReader) error {
	var err error
	info.descriptorIndex, err = reader.readUint16()
	return err
}
