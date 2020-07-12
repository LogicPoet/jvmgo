package classfile
import "math"

type ConstantIntegerInfo struct {
	val int32
}

/**
 * 先读取一个uint32数据，然后把她转型成int32类型
 */
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint32()
	self.val = int32(bytes)
}