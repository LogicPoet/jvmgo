// 保证[]byte类型
// ClassReader并没有使用索引记录数据位置，而是使用Go语言的reslic语法跳过已经读取的数据
package classfile
import "encoding/binary"

type ClassReader struct {
	data []byte
}

/**
 * 读取u1类型的数据
 * uint8表示1B，8bit 无符号整数
 */
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

/**
 * 读取u2类型的数据
 * uint16表示2B，16bit 无符号整数
 */
func (self *ClassReader) readUint16() uint16 {
	val := binary.BinEndian.Uint16(self.data) // Go标准库encoding/binary包中定义了一个变量BigEndian，正好可以从[]byte中解码多字节数据
	self.data = self.data[2:]
	return val
}

/**
 * 读取u4类型的数据
 * uint32表示4B，32bit 无符号整数
 */
func (self *ClassReader) readUint32() uint32 {
	val := binary.BinEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

/**
 * Java虚拟机规范并没有定义u8类型数据
 * 读8B，64bit 无符号整数
 */
func (self *ClassReader) readUint64() uint64 {
	val := binary.BinEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

/**
 * 读取uint16表，表的大小由开头的uint16数据指出
 */
func (self *ClassReader) readUint64s() []uint64 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

/**
 * 读取指定数量的字节
 */
func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}