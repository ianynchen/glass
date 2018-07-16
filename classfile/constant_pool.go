package classfile

import "github.com/ianynchen/glass/util"

type ConstantPool struct {
	info []ConstantInfo
}

func (pool *ConstantPool) read(content []byte, offset int) (int, error) {
	count, err := util.ParseUint16(content[offset:])
	pool.info = make([]ConstantInfo, int(count))

	if err != nil {
		return 2, err
	}

	for i := 1; i < int(count); i++ {
		pool.info[i] = readConstantInfo()
		switch pool.info[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
}
